package processes

import (
	"context"

	"github.com/shirou/gopsutil/v3/process"
)

func KillProcessCtx(ctx context.Context, pid int32) error {
	processes, err := process.ProcessesWithContext(ctx)
	if err != nil {
		return err
	}
	for _, p := range processes {
		//n, err := p.WithContext(ctx)
		//nn, err := p.CmdlineSliceWithContext(ctx)
		//n, err := p.CmdlineWithContext(ctx)
		//fmt.Printf("Process: %s, CMD Slice: %s", n, nnn)
		if err != nil {
			return err
		}
		if p.Pid == pid {
			//fmt.Println("Killing process with commandline: ", n)
			err = p.KillWithContext(ctx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
