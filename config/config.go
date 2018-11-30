package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/runeflow/runeflow/util"
	"github.com/spf13/viper"
)

const (
	apikey      = "apikey"
	endpoint    = "endpoint"
	apiRegister = "api.register"
	apiAuth     = "api.auth"
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
	v.AddConfigPath("/etc/runeflow/")
	v.SetDefault("endpoint", "wss://api.runeflow.com/agent")
	v.SetDefault("api.register", "https://api.runeflow.com/register")
	v.SetDefault("api.auth", "https://api.runeflow.com/agent")
	v.SetEnvPrefix("RUNEFLOW")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	v.ReadInConfig()
	initAPIKey(v)
	return &Config{v: v}
}

// GetAPIKey gets the configured API key
func (c *Config) GetAPIKey() string {
	return c.v.GetString(apikey)
}

// GetEndpoint gets the configured websocket endpoint
func (c *Config) GetEndpoint() string {
	return c.v.GetString(endpoint)
}

// GetAPIRegister gets the registration URL
func (c *Config) GetAPIRegister() string {
	return c.v.GetString(apiRegister)
}

// GetAPIAuth gets the registration URL
func (c *Config) GetAPIAuth() string {
	return c.v.GetString(apiAuth)
}

func initAPIKey(v *viper.Viper) {
	if v.IsSet(apikey) {
		return
	}
	fmt.Println("no api key")
	k, err := util.RandomString()
	if err != nil {
		fmt.Printf("error creating key: %v\n", err)
	}
	fmt.Printf("created key: %s\n", k)
	v.Set(apikey, k)
	f, err := os.OpenFile(v.ConfigFileUsed(), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("error opening config file for writing: %v\n", err)
		return
	}
	defer f.Close()
	if _, err = f.WriteString(fmt.Sprintf("\n%s: %s\n", apikey, k)); err != nil {
		fmt.Printf("error appending api key to file: %v\n", err)
	}
}
