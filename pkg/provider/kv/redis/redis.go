package redis

import (
	"github.com/kvtools/valkeyrie/store"
	"github.com/traefik/traefik/v2/pkg/provider"
	"github.com/traefik/traefik/v2/pkg/provider/kv"
)

var _ provider.Provider = (*Provider)(nil)

// Provider holds configurations of the provider.
type Provider struct {
	kv.Provider `export:"true"`
}

// SetDefaults sets the default values.
func (p *Provider) SetDefaults() {
	p.Provider.SetDefaults()
	p.Endpoints = []string{"127.0.0.1:6379"}
}

// Init the provider.
func (p *Provider) Init() error {
	return p.Provider.Init(store.REDIS, "redis", "")
}
