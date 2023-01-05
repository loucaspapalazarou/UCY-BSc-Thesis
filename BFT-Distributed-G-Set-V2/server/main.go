package main

import (
	"backend/server"
	"os"
	"strings"
	"fmt"
)

func main() {

	// the only thing i know is what i have to do
	// and the servers in the network
	data, err := os.ReadFile("/users/loukis/Thesis/BFT-Distributed-G-Set-V2/server/hosts")
	if err != nil {
		panic(err)
	}
	hosts := strings.Split(strings.ReplaceAll(string(data), "\n\n", "\n"), "\n")
	peers := hosts[:len(hosts)-1]
	for i := 0; i < len(peers); i++ {
		if peers[i] == "[servers]" {
			peers = peers[i+1:]
			break
		}
	}

	server := server.Create(peers)

	for {
		msg, _ := server.Receive_socket.RecvMessage(0)
		fmt.Println(msg)
		server.Receive_socket.SendMessage([]string{msg[0], "World"})
	}

}
