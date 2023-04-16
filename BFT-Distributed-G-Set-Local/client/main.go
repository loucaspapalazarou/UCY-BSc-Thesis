package main

import (
	"flag"
	"frontend/modules"
	"frontend/tools"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	tools.ResetLogFile()

	zctx, _ := zmq.NewContext()

	var auto bool
	var reqs int
	var clients int

	flag.BoolVar(&auto, "auto", false, "Automated")
	flag.IntVar(&reqs, "reqs", 5, "Amount of requests")
	flag.IntVar(&clients, "clients", 5, "Amount of clients (if given)")

	flag.Parse()

	if auto {
		modules.StartAutomated(zctx, clients, reqs)
		return
	}
	modules.StartInteractive(zctx)
}
