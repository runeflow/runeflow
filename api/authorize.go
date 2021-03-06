package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// AuthorizeJSON holds agent authorization details
type AuthorizeJSON struct {
	Email    string `json:"email"`
	AgentID  string `json:"agentID"`
	Hostname string `json:"hostname"`
	OSName   string `json:"osName"`
}

// Authorize sends a request to authorize the configured API key for the
// supplied email address.
func (a *API) Authorize(email, hostname, osName string) error {
	agentID, err := a.conf.GetOrInitAgentID()
	if err != nil {
		return fmt.Errorf("error reading agent ID: %v", err)
	}
	data, err := json.Marshal(&AuthorizeJSON{
		Email:    email,
		AgentID:  agentID,
		Hostname: hostname,
		OSName:   osName,
	})
	if err != nil {
		return err
	}
	resp, err := a.client.Post(a.conf.GetAuthURL(), "application/json", bytes.NewReader(data))
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
