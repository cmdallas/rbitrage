package core

import (
	"fmt"

	"github.com/alecthomas/kong"
)

// Globals global options
type Globals struct {
	Config  string      `help:"Location of config file" default:"~/go/src/github.com/cmdallas/rbitrage/examples/config/.rbitrage.yaml" type:"path"`
	Debug   bool        `short:"d" help:"Enable debug mode"`
	Version versionFlag `name:"version" help:"Print version and quit"`
}

// VersionFlag behavior
type versionFlag string

func (v versionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v versionFlag) IsBool() bool                         { return true }
func (v versionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(vars["version"])
	app.Exit(0)
	return nil
}

// RunCmd run arbitrage struct
type RunCmd struct{}

// Run The main rbitrage entrypoint
func (cmd *RunCmd) Run(globals *Globals) error {
	fmt.Println("To unite all peoples within our nation! To denounce the evils of truth and love!")
	fmt.Println("Team Rocket blasts off at the speed of light! Surrender now, or prepare to fight!")

	return nil
}

// CLI All rbitrage CLI options
type CLI struct {
	Globals
	Run RunCmd `cmd help:"start rbitrage"`
}
