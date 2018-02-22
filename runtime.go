package app

import (
	"github.com/sul-dlss-labs/identifier-service/config"
	"github.com/sul-dlss-labs/identifier-service/db"
	"github.com/sul-dlss-labs/identifier-service/druid"
)

// NewRuntime creates a new application level runtime that
// encapsulates the shared services for this application
func NewRuntime(config *config.Config) (*Runtime, error) {
	return &Runtime{config: config}, nil
}

// NewDefaultRuntime creates a new application level runtime with all
// necessary services pre-configured
func NewDefaultRuntime(config *config.Config) (*Runtime, error) {
	conn := db.NewConnection(config)
	repo := db.NewDynamoIdentifiers(config, conn)
	rt, _ := NewRuntime(config)
	rt.WithMinter(druid.NewMinter(repo))
	return rt, nil
}

// Runtime encapsulates the shared services for this application
type Runtime struct {
	config *config.Config
	minter druid.Minter
}

// Config returns the config for this application
func (r *Runtime) Config() *config.Config {
	return r.config
}

// Minter returns the druid minter for this application
func (r *Runtime) Minter() druid.Minter {
	return r.minter
}

// WithMinter sets the passed in minter on the runtime.
func (r *Runtime) WithMinter(minter druid.Minter) *Runtime {
	r.minter = minter
	return r
}
