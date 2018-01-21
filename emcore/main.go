package main

import (
	"fmt"
	"os"
)

func main() {
	parseFlags()

	err := initLogging()
	if err != nil {
		fmt.Println("error occured while initializing the logger")
		// do not use the logger here since it might have not been initialized
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = loadConfig()
	if err != nil {
		el.Errorln(err.Error())
		el.Fatalln("error occured while loading the configuration file")
	}

	gws, err := initGateways()
	if err != nil {
		el.Errorln(err.Error())
		el.Fatalln("error occured while initializing the receptors")
	}

	procs, err := initProcessors()
	if err != nil {
		el.Errorln(err.Error())
		el.Fatalln("error occured while initializing the processors")
	}

	rtr, err := initRouter()
	if err != nil {
		el.Errorln(err.Error())
		el.Fatalln("error occured while initializing the router")
	}

	rtr.LoadGateways(gws...)
	if err != nil {
		el.Errorln(err.Error())
		el.Fatalln("error occured while loading the gateways into the router")
	}

	rtr.LoadProcessors(procs...)
	if err != nil {
		el.Errorln(err.Error())
		el.Fatalln("error occured while loading the processors into the router")
	}

	for _, cfg := range ec.Routes {
		rtr.NewRoute(cfg.Source, cfg.Destination...)
		if err != nil {
			el.Errorln(err.Error())
			el.Fatalln("error occured while loading a new route")
		}
	}

	rtr.Run()
}
