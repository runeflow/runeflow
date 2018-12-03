package config

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/runeflow/runeflow/util"
	"github.com/spf13/viper"
)

const configPath = "/etc/runeflow"

const (
	endpoint    = "endpoint"
	registerURL = "register_url"
	authURL     = "auth_url"
)

// Config holds configuration info
type Config struct {
	v *viper.Viper
}

// NewConfig creates a new Config
func NewConfig() *Config {
	v := viper.New()
	v.SetConfigName("runeflow")
	v.SetConfigType("yaml")
	v.AddConfigPath(configPath)
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
	agentIDFile := path.Join(configPath, "agent_id")
	data, err := ioutil.ReadFile(agentIDFile)
	if err == nil {
		return strings.TrimSpace(string(data)), nil
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
