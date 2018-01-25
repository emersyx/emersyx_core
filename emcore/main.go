package main

func main() {
	parseFlags()
	initLogging()
	loadConfig()

	rtr := newRouter(
		initGateways(),
		initProcessors(),
		initRoutes(),
	)
	rtr.Run()
}
