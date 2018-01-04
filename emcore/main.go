package main

import (
	"emersyx.net/emersyx_log/emlog"
	golog "log"
)

var log *golog.Logger

func main() {
	log := emlog.NewEmlogStdout("emcore")
	log.Print("emersyx core")
}
