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

	var bdso string
	var auto bool
	var clients int
	var reqs int
	var auto_atomic_adds bool

	flag.StringVar(&bdso, "net", "", "Bdso network")
	flag.BoolVar(&auto, "auto", false, "Automated")
	flag.BoolVar(&auto_atomic_adds, "auto_atomic", false, "Test atomic adds")
	flag.IntVar(&clients, "clients", 1, "Amount of automated clients")
	flag.IntVar(&reqs, "reqs", 5, "Amount of requests")

	flag.Parse()

	if auto_atomic_adds {
		modules.StartAutomatedAtomicAdds(zctx, reqs)
		return
	}

	if bdso != "sbdso" && bdso != "bdso-1" && bdso != "bdso-2" {
		panic("Invalid network")
	}

	if auto {
		modules.StartAutomated(zctx, clients, reqs, bdso)
		return
	}
	modules.StartInteractive(zctx, bdso)
}
