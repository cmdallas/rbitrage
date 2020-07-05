package provider

import (
	"github.com/cmdallas/rbitrage/pkg/config"
	provideraws "github.com/cmdallas/rbitrage/pkg/provider/aws"
)

// PriceDiscovery get pricing for specified instance types across all defined providers
func PriceDiscovery(app config.Application) {
	instanceTypes := app.Properties.Providers[0].Nodes.TypeOverrides
	region := app.Properties.Providers[0].Nodes.Region

	provideraws.DiscoverPrice(instanceTypes, region)
}
