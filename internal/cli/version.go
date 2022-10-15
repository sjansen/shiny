package cli

import (
	"fmt"
	"os"
	"runtime/debug"
)

type versionCmd struct {
	version string
	Verbose bool `kong:"short='v',help='Show build details.'"`
}

func (cmd *versionCmd) Run() error {
	if cmd.Verbose {
		fmt.Println("Shiny", cmd.version)
		fmt.Println()
		fmt.Println("Build Info:")
		fmt.Println("===========")
		info, err := debug.ReadBuildInfo()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			fmt.Println(info)
		}
	} else {
		fmt.Println(cmd.version)
	}
	return nil
}
