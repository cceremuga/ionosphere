// Package config is a service for handling configuration items.
package config

import (
	"io/ioutil"
	"time"

	"github.com/cceremuga/ionosphere/services/log"
	"gopkg.in/yaml.v2"
)

// Rtl represents the RTL-SDR config.
type Rtl struct {
	Path            string
	Frequency       string
	Gain            string
	PpmError        string `yaml:"ppm-error"`
	SquelchLevel    string `yaml:"squelch-level"`
	SampleRate      string `yaml:"sample-rate"`
	AdditionalFlags string `yaml:"additional-flags"`
}

// Multimon represents the multimon-ng config.
type Multimon struct {
	Path            string
	AdditionalFlags string `yaml:"additional-flags"`
}

// Handler represents an individual handler's config.
type Handler struct {
	ID      string
	Name    string
	Options map[string]string
}

// Beacon represents the periodic beacon packet config.
type Beacon struct {
	Enabled     bool
	Latitude    float32
	Longitude   float32
	Interval    time.Duration
	Comment     string
	SymbolTable string `yaml:"symbol-table"`
	Symbol      string
}

// Config represents the full, unmarshalled YAML config.
type Config struct {
	Rtl      Rtl
	Multimon Multimon
	Beacon   Beacon
	Handlers []Handler
}

var c Config

// Load unmarshals and caches the YAML config.
func Load() *Config {
	if c.Rtl.Path != "" {
		return &c
	}

	f := yamlFile()

	err := yaml.Unmarshal(f, &c)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return &c
}

// HandlerOptions retrieves the map of options for a given Handler Id.
func HandlerOptions(id string) map[string]string {
	handlers := Load().Handlers

	for i := 0; i < len(handlers); i++ {
		handler := handlers[i]

		if handler.ID == id {
			return handler.Options
		}
	}

	return nil
}

func yamlFile() []byte {
	f, err := ioutil.ReadFile("config/config.yml")

	if err != nil {
		log.Fatal("Could not load config file.")
	}

	return f
}
