package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"pidKiller/pkg/config"
	"pidKiller/pkg/log"
	"pidKiller/pkg/processes"
	"syscall"
	"time"

	"github.com/gosuri/uilive"
)

var (
	debug    *bool   // debug flag - long
	f_config *string // Storage variable for config file (short)
)

func init() {
	debug = flag.Bool("debug", false, "Enable debug logging")
	f_config = flag.String("c", "/workspace/example/example.yaml", "Config file to read")
}

func cleanup() {
	log.Warn("SIGINT caught... cleaning up")
	log.Info("pidKiller stopped")
}

func main() {

	// Signal handler
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	// Parse the command line arguments
	flag.Parse()

	// Setup logger
	if *debug {
		// Enable debug logging
		log.EnableDebug()
	} else {
		// Enable info logging
		log.Info("Starting pidKiller")
	}

	// Verify config file exists
	if _, err := os.Stat(*f_config); os.IsNotExist(err) {
		log.Fatal(fmt.Sprintf("Config file does not exist: %v", err))
		os.Exit(1)
	}

	// Read config file
	conf, err := config.ReadConfig(*f_config)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error reading config file: %v", err))
	}

	// Check time
	start := time.Now()
	log.Info(fmt.Sprintf("Start time: %v", start))

	// get an instance of writer and start listening for updates and render
	writer := uilive.New()
	writer.Start()
	breakLoop := false
	totalDuration := time.Duration(conf.Terminate.Hours)*time.Hour +
		time.Duration(conf.Terminate.Minutes)*time.Minute +
		time.Duration(conf.Terminate.Seconds)*time.Second

	// Loop time
	for !breakLoop {
		// Check time every second
		time.Sleep(1 * time.Second)
		elapsed := time.Since(start)
		remaining := totalDuration - elapsed

		// Format remaining time to HH:MM:SS
		hours := int(remaining.Hours())
		minutes := int(remaining.Minutes()) % 60
		seconds := int(remaining.Seconds()) % 60
		formattedTime := fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)

		//fmt.Printf("Time remaining: %s\n", formattedTime)
		fmt.Fprintf(writer, "Time remaining: %s\n", formattedTime)
		if time.Since(start) >= totalDuration {
			breakLoop = true
		}
	}
	writer.Stop()

	// Kill processes - Single thread for now
	ctx := context.Background()
	for _, p := range conf.Processes {
		err = processes.KillProcessCtx(ctx, int32(p.PID))
		if err != nil {
			log.Fatal(fmt.Sprintf("Error process: %v, PID: %v: %v", p.Name, p.PID, err))
		}
		log.Info(fmt.Sprintf("Killed process: %v, PID: %v", p.Name, p.PID))
	}
}
