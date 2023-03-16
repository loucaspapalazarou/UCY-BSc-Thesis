package main

import (
	"backend/config"
	"backend/modules"
	"backend/tools"
	"fmt"
	"os"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	tools.ResetLogFile()

	zctx, err := zmq.NewContext()
	if err != nil {
		panic(err)
	}
	server_nodes := config.SetServerNodes()

	if len(os.Args) < 2 {
		modules.Start(server_nodes, "NORMAL", zctx)
		select {}
	}
	if len(os.Args) == 2 && os.Args[1] == "mutes" || os.Args[1] == "m" {
		modules.Start(server_nodes, "MUTES", zctx)
		select {}
	}
	if len(os.Args) > 2 {
		fmt.Println("Wrong arguments")
		return
	}

}
