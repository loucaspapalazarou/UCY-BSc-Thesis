// // Server

package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func server_task(port string, context *zmq.Context) {

	g_set := make(map[string]string)
	g_set["record1"] = "dog"
	g_set["record2"] = "blue"
	g_set["record3"] = "cat"
	g_set["record4"] = "red"

	// server_count := 5

	inbound_socket, _ := context.NewSocket(zmq.ROUTER)
	inbound_socket.Bind(port)
	fmt.Println("Client facing socket is up in port: ", port)

	// oubound_sockets := make([]*zmq.Socket, 0)
	// for i := 0; i < server_count; i++ {
	// 	s, _ := context.NewSocket(zmq.DEALER)
	// 	port := "tcp://*:1000" + strconv.Itoa(i)
	// 	s.Bind(port)
	// 	fmt.Println("Bound dealer to port:", port)
	// 	oubound_sockets = append(oubound_sockets, s)
	// }
	// fmt.Println("Server facing sockets are up!")

	for {
		msg, _ := inbound_socket.RecvMessage(0)
		fmt.Println(msg)
		response := []string{msg[0], "World", port}
		inbound_socket.SendMessage(response)
	}
}

func main() {
	zctx, _ := zmq.NewContext()
	go server_task("tcp://*:5555", zctx)
	go server_task("tcp://*:5556", zctx)
	for {
	}
}