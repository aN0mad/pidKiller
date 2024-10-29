package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config struct is the config file
type Config struct {
	Processes []Process `yaml:"processes"`
	Terminate Terminate `yaml:"terminate"`
}

// Terminate is a top level attribute within the config file
type Terminate struct {
	Signal  string `yaml:"signal"` // Signal to send to process - TODO: Not used right now
	Hours   int    `yaml:"hours"`
	Minutes int    `yaml:"minutes"`
	Seconds int    `yaml:"seconds"`
}

// Process is the single process to handle
type Process struct {
	Name string `yaml:"name"`
	PID  int    `yaml:"pid"`
}

// ReadConfig reads the config file and marshals to structs
func ReadConfig(conf string) (Config, error) {
	// Read config file in
	b_conf, err := os.ReadFile(conf)
	if err != nil {
		panic(err)
	}

	// Marshal to struct
	var c Config
	err = yaml.Unmarshal([]byte(b_conf), &c)
	if err != nil {
		return c, err
	}
	return c, nil
}
