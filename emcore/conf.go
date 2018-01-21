package main

import (
	"github.com/BurntSushi/toml"
)

// ec is the emersyxConfig global instance which holds all values from the config file.
var ec emersyxConfig

type processorConfig struct {
	Identifier string
	Config     string
	PluginPath string `toml:"plugin_path"`
}

type ircGatewayConfig struct {
	Identifier    *string
	Nick          *string
	Ident         *string
	Name          *string
	Version       *string
	ServerAddress *string `toml:"server_address"`
	ServerPort    *uint   `toml:"server_port"`
	ServerUseSSL  *bool   `toml:"server_use_ssl"`
	QuitMessage   *string `toml:"quit_message"`
	PluginPath    *string `toml:"plugin_path"`
}

type telegramGatewayConfig struct {
	Identifier     *string
	APIToken       *string   `toml:"api_token"`
	UpdatesLimit   *uint     `toml:"updates_limit"`
	UpdatesTimeout *uint     `toml:"updates_timeout"`
	UpdatesAllowed *[]string `toml:"updates_allowed"`
	PluginPath     *string   `toml:"plugin_path"`
}

type routerConfig struct {
	PluginPath string `toml:"plugin_path"`
}

type routeConfig struct {
	Source      string
	Destination []string
}

type emersyxConfig struct {
	Processors       []processorConfig
	IRCGateways      []ircGatewayConfig
	TelegramGateways []telegramGatewayConfig
	Router           routerConfig
	Routes           []routeConfig
}

// loadConfig opens, reads and parses the toml configuration file specified as command line argument. The parseFlags
// function needs to be called before this one.
func loadConfig() error {
	// read the parameters from the specified configuration file
	_, err := toml.DecodeFile(*flConfFile, &ec)
	if err != nil {
		return err
	}
	return nil
}
