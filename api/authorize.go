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
	Email   string `json:"email"`
	AgentID string `json:"agentID"`
}

// Authorize sends a request to authorize the configured API key for the
// supplied email address.
func (a *API) Authorize(email string) error {
	data, err := json.Marshal(&AuthorizeJSON{
		Email:   email,
		AgentID: a.conf.GetAPIKey(),
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
