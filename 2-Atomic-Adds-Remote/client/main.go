package main

import (
	"2-Atomic-Adds/modules"
	"2-Atomic-Adds/tools"
	"flag"

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
	var my_message string
	var my_id string
	var peer_id string
	var peer_message string
	var dest string

	flag.StringVar(&bdso, "net", "", "Bdso network")
	flag.BoolVar(&auto, "auto", false, "Automated")
	flag.IntVar(&clients, "clients", 1, "Amount of automated clients")
	flag.IntVar(&reqs, "reqs", 5, "Amount of requests")

	// Flags for testing 2-atomic-adds
	flag.BoolVar(&auto_atomic_adds, "auto_atomic", false, "Test atomic adds")
	flag.StringVar(&my_id, "my_id", "", "Client's id")
	flag.StringVar(&my_message, "my_msg", "", "The atomic message")
	flag.StringVar(&peer_id, "peer_id", "", "Client's peer id")
	flag.StringVar(&peer_message, "peer_msg", "", "The peer's atomic message")
	flag.StringVar(&dest, "dest", "", "The destination network")

	flag.Parse()

	if auto_atomic_adds {
		if my_id == "" || my_message == "" || peer_id == "" || peer_message == "" || dest == "" {
			panic("Empty arguments!")
		}
		modules.StartAutomatedAtomic(my_id, my_message, peer_id, peer_message, dest, zctx, reqs)
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
