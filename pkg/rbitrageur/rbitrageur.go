package rbitrageur

import (
	"fmt"

	"github.com/cmdallas/rbitrage/pkg/config"
	"github.com/cmdallas/rbitrage/pkg/provider"
)

// Arbitrate Reach a settlement on the instances to run
func Arbitrate(c *config.Config) {
	// todo: move config validation
	if len(c.Applications) > 1 {
		panic("This version of rbitrage only supports 1 application")
	}

	app := c.Applications[0]
	fmt.Printf("Using config: %v\n", app)

	for _, p := range app.Properties.Providers {
		if p.Name != "aws" {
			panic("This version of rbitrage only supports provider aws")
		}
	}

	provider.PriceDiscovery(app)
}
