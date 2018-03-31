package main

import (
	"emersyx.net/emersyx_apis/emcomapi"
	"emersyx.net/emersyx_apis/emircapi"
	"emersyx.net/emersyx_apis/emtgapi"
	"emersyx.net/emersyx_log/emlog"
	"flag"
	"fmt"
	"io"
	"os"
	"plugin"
)

// flLogStdout holds the value of the command line flag which specifies whether to print logging messages to standard
// output or not.
var flLogStdout *bool

// flLogFile holds the value of the command line flag which specifies the file to write logging messages to.
var flLogFile *string

// flLogLevel holds the value of the command line flag which specifies the logging level.
var flLogLevel *uint

// flConfFile holds the value of the command line flag which specifies the emersyx configuration file.
var flConfFile *string

// elSinks is an io.MultiWriter instance containing all io.Writer instances to which logging messages must be sent.
var elSinks io.Writer

// el is the emlog.EmersyxLogger global instance used throughout the emcore component.
var el *emlog.EmersyxLogger

// plugins is a map object. The keys are filesystem paths to go plugin files and the values are pointers to
// plugin.Plugin objects. This map is used by the getPlugin function.
var plugins = make(map[string]*plugin.Plugin)

// parseFlags parses the command line arguments given to the emersyx binary.
func parseFlags() {
	// set the expected flags
	flLogStdout = flag.Bool("logstdout", false, "log to standard output")
	flLogFile = flag.String("logfile", "", "file to store logs into")
	flLogLevel = flag.Uint("loglevel", 0, "logging verbosity level")
	flConfFile = flag.String("conffile", "", "file to read configuration parameters from")

	// parse the flags
	flag.Parse()
}

// initLogging configures the logger (i.e. the el global variable). The parseFlags function needs to be called before
// this one.
func initLogging() {
	var err error
	var sinks []io.Writer

	if flLogStdout != nil && *flLogStdout == true {
		sinks = append(sinks, os.Stdout)
	}

	if len(*flLogFile) > 0 {
		f, err := os.OpenFile(*flLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("error occured while initializing the logger")
			os.Exit(1)
		}
		sinks = append(sinks, f)
	}

	elSinks = io.MultiWriter(sinks...)
	el, err = emlog.NewEmersyxLogger(elSinks, "emcore", *flLogLevel)
	if err != nil {
		// do not use the logger here since it might have not been initialized
		fmt.Println(err.Error())
		fmt.Println("error occured while initializing the logger")
		os.Exit(1)
	}
}

// getPlugin loads a go plugin file at a specified path and returns a *plugin.Plugin object. The function caches
// previously opened plugins in the global "plugins" map object.
func getPlugin(path string) *plugin.Plugin {
	var p *plugin.Plugin
	p, ok := plugins[path]

	// check if the plugin was previously opened and cached, and if not then open it now
	if ok != true {
		var err error
		p, err = plugin.Open(path)
		if err != nil {
			el.Errorln(err.Error())
			el.Fatalf("error occured while loading go plugin at path \"%s\"\n", path)
			return nil
		}
		// if a new plugin has been opened, then save it into the "plugins" global map
		plugins[path] = p
	}

	return p
}

// newIRCGateway creates a new emircapi.IRCGateway object using the provided ircGatewayConfig argument. Under the hood,
// the emircapi.NewIRCGateway function is used.
func newIRCGateway(cfg ircGatewayConfig) emircapi.IRCGateway {
	p := getPlugin(*cfg.PluginPath)

	optarg := make([]func(emircapi.IRCGateway) error, 0)
	opt, err := emircapi.NewIRCOptions(p)
	optarg = append(optarg, opt.Logging(elSinks, *flLogLevel))
	optarg = append(optarg, opt.Identifier(*cfg.Identifier))

	if cfg.Nick != nil {
		optarg = append(optarg, opt.Nick(*cfg.Nick))
	}
	if cfg.Ident != nil {
		optarg = append(optarg, opt.Ident(*cfg.Ident))
	}
	if cfg.Name != nil {
		optarg = append(optarg, opt.Name(*cfg.Name))
	}
	if cfg.Version != nil {
		optarg = append(optarg, opt.Version(*cfg.Version))
	}
	if cfg.ServerAddress != nil && cfg.ServerPort != nil && cfg.ServerUseSSL != nil {
		optarg = append(optarg, opt.Server(*cfg.ServerAddress, *cfg.ServerPort, *cfg.ServerUseSSL))
	}
	if cfg.QuitMessage != nil {
		optarg = append(optarg, opt.QuitMessage(*cfg.QuitMessage))
	}

	gw, err := emircapi.NewIRCGateway(p, optarg...)
	if err != nil {
		el.Errorln(err.Error())
		el.Fatalln("error occured while creating a new IRC gateway")
	}

	// automatically connect to the IRC server
	if err = gw.Connect(); err != nil {
		el.Errorln(err.Error())
		el.Fatalf("could not connect to the \"%s\" IRC server", *cfg.Identifier)
		return nil
	}

	return gw
}

// loadIRCGateways creates and initializez emircapi.IRCGateway objects for all IRC gateways specified in the emersyx
// configuration file. The objects are returned in an array of type []emcomapi.Identifiable.
func loadIRCGateways() []emcomapi.Gateway {
	gws := make([]emcomapi.Gateway, len(ec.IRCGateways))

	for i, cfg := range ec.IRCGateways {
		gws[i] = newIRCGateway(cfg)
	}

	return gws
}

// newTelegramGateway creates a new emtgapi.TelegramGateway object using the provided telegramGatewayConfig argument.
// Under the hood, the emtgapi.NewTelegramGateway function is used.
func newTelegramGateway(cfg telegramGatewayConfig) emtgapi.TelegramGateway {
	p := getPlugin(*cfg.PluginPath)

	optarg := make([]func(emtgapi.TelegramGateway) error, 0)
	opt, err := emtgapi.NewTelegramOptions(p)
	optarg = append(optarg, opt.Logging(elSinks, *flLogLevel))
	optarg = append(optarg, opt.Identifier(*cfg.Identifier))

	if cfg.APIToken != nil {
		optarg = append(optarg, opt.APIToken(*cfg.APIToken))
	}
	if cfg.UpdatesLimit != nil {
		optarg = append(optarg, opt.UpdatesLimit(*cfg.UpdatesLimit))
	}
	if cfg.UpdatesTimeout != nil {
		optarg = append(optarg, opt.UpdatesTimeout(*cfg.UpdatesTimeout))
	}
	if cfg.UpdatesAllowed != nil {
		optarg = append(optarg, opt.UpdatesAllowed(*cfg.UpdatesAllowed...))
	}

	gw, err := emtgapi.NewTelegramGateway(p, optarg...)
	if err != nil {
		el.Errorln(err.Error())
		el.Fatalln("error occured while creating a new IRC gateway")
	}
	return gw
}

// loadTelegramGateways creates and initializez emtgapi.TelegramGateway objects for all Telegram gateways specified in
// the emersyx configuration file. The objects are returned in an array of type []emcomapi.Identifiable.
func loadTelegramGateways() []emcomapi.Gateway {
	gws := make([]emcomapi.Gateway, len(ec.TelegramGateways))

	for i, cfg := range ec.TelegramGateways {
		gws[i] = newTelegramGateway(cfg)
	}

	return gws
}

// initGateways creates and initializez objects for all gateways specified in the emersyx configuration file. The
// objects are returned in an array of type []emcomapi.Identifiable.
func initGateways() []emcomapi.Gateway {
	gws := make([]emcomapi.Gateway, 0)

	irc := loadIRCGateways()
	gws = append(gws, irc...)

	tg := loadTelegramGateways()
	gws = append(gws, tg...)

	return gws
}

// initProcessors creates and initializez emcomapi.Processor objects for all processors specified in the emersyx
// configuration file. The objects are returned in an array of type []emcomapi.Processor.
func initProcessors(rtr emcomapi.Router) []emcomapi.Processor {
	procs := make([]emcomapi.Processor, len(ec.Processors))

	for i, pcfg := range ec.Processors {
		p := getPlugin(pcfg.PluginPath)
		opts, err := emcomapi.NewProcessorOptions(p)
		if err != nil {
			el.Errorln(err.Error())
			el.Fatalln("error occured while creating a new ProcessorOptions instance")
		}

		proc, err := emcomapi.NewProcessor(p,
			opts.Identifier(pcfg.Identifier),
			opts.Config(pcfg.Config),
			opts.Router(rtr),
			opts.Logging(elSinks, *flLogLevel),
		)
		if err != nil {
			el.Errorln(err.Error())
			el.Fatalln(
				fmt.Sprintf(
					"error occured while creating a new Processor instance with identifier \"%s\"",
					pcfg.Identifier,
				),
			)
		}
		procs[i] = proc
	}

	return procs
}

// initRoutes formats the route information from the global emersyxConfig instance (initialized via loadConfig) such
// that it can be passed as argument to the emcomapi.RouterOptions.Routes method.
func initRoutes() map[string][]string {
	var m = make(map[string][]string)

	for _, cfg := range ec.Routes {
		val, ok := m[cfg.Source]
		if ok {
			val := append(val, cfg.Destination...)
			m[cfg.Source] = val
		} else {
			narr := make([]string, len(cfg.Destination))
			copy(narr, cfg.Destination)
			m[cfg.Source] = narr
		}
	}

	return m
}

// newRouter creates an emcomapi.Router object as specified in the emersyx configuration file. Under the hood, the
// emcomapi.NewRouter function is used.
func newRouter() emcomapi.Router {
	p := getPlugin(ec.Router.PluginPath)
	rtr, err := emcomapi.NewRouter(p)
	if err != nil {
		el.Errorln(err.Error())
		el.Fatalln("error occured while creating a new router")
	}
	return rtr
}

// initRouter sets the options on the provided Router object. The options include Gateways, Processors and Routes. Under
// the hood, an emcomapi.RouterOptions object is used together with emcomapi.Router.SetOptions.
func initRouter(
	rtr emcomapi.Router,
	gws []emcomapi.Gateway,
	procs []emcomapi.Processor,
	routes map[string][]string,
) {
	if gws == nil || len(gws) == 0 {
		el.Fatalln("cannot create a router without any gateways")
	}
	if procs == nil || len(procs) == 0 {
		el.Fatalln("cannot create a router without any processors")
	}
	if routes == nil || len(routes) == 0 {
		el.Fatalln("cannot create a router without any routes")
	} else {
		for key, val := range routes {
			if val == nil || len(val) == 0 {
				el.Fatalln(fmt.Sprintf("route for receptor source \"%s\" has no processor destinations", key))
			}
		}
	}

	optarg := make([]func(emcomapi.Router) error, 0)
	p := getPlugin(ec.Router.PluginPath)
	opt, err := emcomapi.NewRouterOptions(p)
	if err != nil {
		el.Errorln(err.Error())
		el.Fatalln("error occured while creating a new RouterOptions instance")
	}

	optarg = append(optarg, opt.Gateways(gws...))
	optarg = append(optarg, opt.Processors(procs...))
	optarg = append(optarg, opt.Routes(routes))
	optarg = append(optarg, opt.Logging(elSinks, *flLogLevel))

	err = rtr.SetOptions(optarg...)
	if err != nil {
		el.Errorln(err.Error())
		el.Fatalln("error occured while configuring the event router")
	}
}
