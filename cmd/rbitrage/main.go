package main

import (
	"github.com/alecthomas/kong"
	"github.com/cmdallas/rbitrage/cmd/rbitrage/core"
)

func main() {
	cli := core.CLI{}

	ctx := kong.Parse(
		&cli,
		kong.Name("rbitrage"),
		kong.Description("An open source multicloud arbitrage worker."),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{Compact: true}),
		kong.Vars{"version": "0.0.1"},
	)

	err := ctx.Run(&cli.Globals)
	ctx.FatalIfErrorf(err)
}
