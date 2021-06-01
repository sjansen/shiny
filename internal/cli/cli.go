package cli

import (
	"github.com/alecthomas/kong"
)

var cli struct {
	Fmt      fmtCmd      `kong:"cmd,help='Reformat files in the standard style'"`
	Help     helpCmd     `kong:"cmd,help='Show this help text'"`
	Generate generateCmd `kong:"cmd,help='Generate Go code for the schema directory'"`
	Version  versionCmd  `kong:"cmd,help='Show the current Shiny version'"`
}

// ParseAndRun parses command line arguments then runs the matching command.
func ParseAndRun(version string) {
	ctx := kong.Parse(&cli,
		kong.UsageOnError(),
	)
	cli.Help.ctx = ctx
	cli.Version.version = version

	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
