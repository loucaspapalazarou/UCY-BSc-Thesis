package modules

import (
	"backend/config"
	"backend/messaging"
	"backend/server"
	"time"
)

func Leader_task(leader config.Node, peers []config.Node) {
	server := server.Create(leader, peers)
	time.Sleep(time.Second * 1)
	m, _ := messaging.StringToMessage([]string{server.Id, messaging.BRACHA_BROADCAST, "1"})
	messaging.ReliableBroadcast(server, m)
	m, _ = messaging.StringToMessage([]string{server.Id, messaging.BRACHA_BROADCAST, "2"})
	messaging.ReliableBroadcast(server, m)
	m, _ = messaging.StringToMessage([]string{server.Id, messaging.BRACHA_BROADCAST, "3"})
	messaging.ReliableBroadcast(server, m)
	for {
		message, _ := server.Receive_socket.RecvMessage(0)
		messaging.HandleMessage(server, message)
	}
}