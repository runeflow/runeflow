package api

import (
	"net/http"
	"time"

	"github.com/runeflow/runeflow/config"
)

// API communicates with the Runeflow API server
type API struct {
	client *http.Client
	conf   *config.Config
}

// NewAPI creates a new API with the given configuration settings
func NewAPI(c *config.Config) *API {
	return &API{
		conf: c,
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}
