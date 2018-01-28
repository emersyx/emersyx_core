package main

func main() {
	parseFlags()
	initLogging()
	loadConfig()

	rtr := newRouter()
	initRouter(
		rtr,
		initGateways(),
		initProcessors(rtr),
		initRoutes(),
	)

	rtr.Run()
}
