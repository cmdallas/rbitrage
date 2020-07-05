package main

import (
	"github.com/alecthomas/kong"
	"github.com/cmdallas/rbitrage/cmd/rbitrage/core"
	"github.com/cmdallas/rbitrage/pkg/config"
	"github.com/cmdallas/rbitrage/pkg/rbitrageur"
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

	cfg, err := config.NewConfig(cli.Globals.Config)
	ctx.FatalIfErrorf(err)

	rbitrageur.Arbitrate(cfg)
}
