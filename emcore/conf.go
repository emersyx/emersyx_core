package main

import (
	"emersyx.net/emersyx_log/emlog"
	"flag"
	"github.com/BurntSushi/toml"
)

type receptorConfig struct {
	ReceptorType string `toml:"type"`
	Plugin       string
}

type processorConfig struct {
	Plugin string
	Config string
}

type routerConfig struct {
	Plugin string
}

type routeConfig struct {
	Destination []string
}

type ircBotConfig struct {
	Nick          *string
	Ident         *string
	Name          *string
	Version       *string
	ServerAddress *string `toml:"server_address"`
	ServerPort    *int    `toml:"server_port"`
	ServerUseSSL  *bool   `toml:"server_use_ssl"`
	QuitMessage   *string `toml:"quit_message"`
}

type telegramBotConfig struct {
	APIToken       *string   `toml:"api_token"`
	UpdatesLimit   *int      `toml:"updates_limit"`
	UpdatesTimeout *int      `toml:"updates_timeout"`
	UpdatesAllowed *[]string `toml:"updates_allowed"`
}

type emersyxConfig struct {
	Receptors    map[string]receptorConfig
	Processors   map[string]processorConfig
	Router       routerConfig
	Routes       map[string]routeConfig
	IRCBots      map[string]ircBotConfig      `toml:"IRCBot"`
	TelegramBots map[string]telegramBotConfig `toml:"TelegramBot"`
}

var ec emersyxConfig

// emlog is the emlog.EmersyxLogger instance used throughout the emcore component.
var el emlog.EmersyxLogger

// parseFlags parses the command line arguments given to the emersyx binary.
func parseFlags() error {
	var err error

	// set the expected flags
	logStdout := flag.Bool("logstdout", false, "log to standard output")
	logFile := flag.String("logfile", "", "file to store logs into")
	logLevel := flag.Uint("loglevel", 0, "logging verbosity level")
	confFile := flag.String("conffile", "", "file to read configuration parameters from")

	// parse the flags
	flag.Parse()

	// create the emlog.Emlog object
	el, err = emlog.NewEmersyxLogger(*logStdout, *logFile, "emcore", *logLevel)
	if err != nil {
		return err
	}

	// read the parameters from specified configuration file
	_, err = toml.DecodeFile(*confFile, &ec)
	if err != nil {
		return err
	}

	return nil
}
