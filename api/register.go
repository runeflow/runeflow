package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// RegisterJSON is the message sent for registration
type RegisterJSON struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

// Register sends a message to the register endpoint
func (a *API) Register(email, firstName, lastName string) error {
	data, err := json.Marshal(&RegisterJSON{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	})
	if err != nil {
		return err
	}
	resp, err := a.client.Post(a.conf.GetRegisterURL(), "application/json", bytes.NewReader(data))
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
