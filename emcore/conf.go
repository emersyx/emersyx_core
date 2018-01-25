package main

import (
	"github.com/BurntSushi/toml"
)

// ec is the emersyxConfig global instance which holds all values from the config file.
var ec emersyxConfig

// processorConfig is the struct for holding processor configuration values from the emersyx configuration file.
type processorConfig struct {
	Identifier string
	Config     string
	PluginPath string `toml:"plugin_path"`
}

// ircGatewayConfig is the struct for holding IRC gateway configuration values from the emersyx configuration file.
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

// telegramGatewayConfig is the struct for holding Telegram gateway configuration values from the emersyx configuration
// file.
type telegramGatewayConfig struct {
	Identifier     *string
	APIToken       *string   `toml:"api_token"`
	UpdatesLimit   *uint     `toml:"updates_limit"`
	UpdatesTimeout *uint     `toml:"updates_timeout"`
	UpdatesAllowed *[]string `toml:"updates_allowed"`
	PluginPath     *string   `toml:"plugin_path"`
}

// routerConfig is the struct for holding router configuration values from the emersyx configuration file.
type routerConfig struct {
	PluginPath string `toml:"plugin_path"`
}

// routeConfig is the struct for holding route configuration values from the emersyx configuration file.
type routeConfig struct {
	Source      string
	Destination []string
}

// emersyxConfig is the container struct for holding all configuration values from the emersyx configuration file.
type emersyxConfig struct {
	Processors       []processorConfig
	IRCGateways      []ircGatewayConfig
	TelegramGateways []telegramGatewayConfig
	Router           routerConfig
	Routes           []routeConfig
}

// loadConfig opens, reads and parses the toml configuration file specified as command line argument. The parseFlags
// function needs to be called before this one.
func loadConfig() {
	// read the parameters from the specified configuration file
	_, err := toml.DecodeFile(*flConfFile, &ec)
	if err != nil {
		el.Errorln(err.Error())
		el.Fatalln("error occured while loading the configuration file")
	}
}
