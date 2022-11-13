// Package handler is a service for interacting with handlers in bulk.
package handler

import (
	"github.com/cceremuga/ionosphere/handlers/aprsis"
	"github.com/cceremuga/ionosphere/handlers/stdout"
	"github.com/cceremuga/ionosphere/interfaces"
	"github.com/cceremuga/ionosphere/services/config"
	"github.com/cceremuga/ionosphere/services/log"
)

var cache []interfaces.Handler

// Start initializes all handlers and caches them.
func Start() []interfaces.Handler {
	hConfigs := config.Load().Handlers

	for i := 0; i < len(hConfigs); i++ {
		h := instantiate(hConfigs[i].Name)

		if h == nil || !h.Enabled() {
			continue
		}

		h.Start()
		log.Debug("%s handler initialized.", h.Name())
		cache = append(cache, h)
	}

	return cache
}

// All returns all cached handlers. Double checks that cache is populated.
// If it is not, lazy loads it.
func All() []interfaces.Handler {
	if len(cache) > 0 {
		return cache
	}

	Start()

	return cache
}

func instantiate(name string) interfaces.Handler {
	// TODO: Someday, reflect this at runtime from package name in config maybe?
	switch name {
	case "stdout":
		return new(stdout.Stdout)
	case "aprsis":
		return new(aprsis.APRSIS)
	}

	return nil
}
