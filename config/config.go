package config

import (
	"io/ioutil"
	"strings"

	"github.com/runeflow/runeflow/util"
	"github.com/spf13/viper"
)

const (
	endpoint    = "endpoint"
	registerURL = "register_url"
	authURL     = "auth_url"
)

const agentIDFile = "/var/lib/runeflow/agent_id"

// Config holds configuration info
type Config struct {
	v *viper.Viper
}

// NewConfig creates a new Config
func NewConfig() *Config {
	v := viper.New()
	v.SetConfigName("runeflow")
	v.SetConfigType("yaml")
	v.AddConfigPath("/etc/runeflow")
	v.SetDefault(endpoint, "wss://api.runeflow.com/agent")
	v.SetDefault(registerURL, "https://api.runeflow.com/register")
	v.SetDefault(authURL, "https://api.runeflow.com/agent")
	v.SetEnvPrefix("RUNEFLOW")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	v.ReadInConfig()
	return &Config{v: v}
}

// GetAgentID gets the configured API key
func (c *Config) GetAgentID() (string, error) {
	data, err := ioutil.ReadFile(agentIDFile)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

// GetOrInitAgentID gets the configured Agent ID, or if one has not been
// configured, generates a new one
func (c *Config) GetOrInitAgentID() (string, error) {
	id, err := c.GetAgentID()
	if err == nil {
		return id, nil
	}
	k, err := util.RandomString()
	if err != nil {
		return "", err
	}
	if err := ioutil.WriteFile(agentIDFile, []byte(k+"\n"), 0644); err != nil {
		return "", err
	}
	return k, nil
}

// GetEndpoint gets the configured websocket endpoint
func (c *Config) GetEndpoint() string {
	return c.v.GetString(endpoint)
}

// GetRegisterURL gets the registration URL
func (c *Config) GetRegisterURL() string {
	return c.v.GetString(registerURL)
}

// GetAuthURL gets the new agent authorization URL
func (c *Config) GetAuthURL() string {
	return c.v.GetString(authURL)
}
