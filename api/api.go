package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

// Register sends a message to the register endpoint
func (a *API) Register(msg interface{}) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	resp, err := a.client.Post(a.conf.GetAPIRegister(), "application/json", bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("error making HTTP request: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return fmt.Errorf("error reading from server: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("registration error %d: %v", resp.StatusCode, string(body))
	}
	return nil
}

// Authorize sends a request to authorize the configured API key for the
// supplied email address.
func (a *API) Authorize(email string) error {
	data, err := json.Marshal(map[string]string{
		"email":   email,
		"agentID": a.conf.GetAPIKey(),
	})
	if err != nil {
		return err
	}
	resp, err := a.client.Post(a.conf.GetAPIAuth(), "application/json", bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("error making HTTP request: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return fmt.Errorf("error reading from server: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("authorization error %d: %v", resp.StatusCode, string(body))
	}
	return nil
}
