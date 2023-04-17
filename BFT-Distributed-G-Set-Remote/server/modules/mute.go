package modules

import (
	"BFT-Distributed-G-Set-Remote/config"
	"BFT-Distributed-G-Set-Remote/server"
	"BFT-Distributed-G-Set-Remote/tools"
	"os"
	"strconv"
	"strings"
)

func StartMute(servers []config.Node, default_port, num_threads int) {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	hostname = strings.Split(hostname, ".")[0]
	for i := default_port; i < default_port+num_threads; i++ {
		go func(my_port int) {
			p := strconv.Itoa(my_port)
			me := config.Node{Host: hostname, Port: p}
			s := server.CreateServer(me, servers)
			tools.Log(s.Id, "Started with MUTE behaviour")
			for {
				msg, err := s.Receive_socket.RecvMessage(0)
				if err != nil {
					tools.Log(s.Id, err.Error())
					return
				}
				tools.Log(s.Id, "Received {"+strings.Join(msg, " ")+"}, no action")
			}
		}(i)
	}
}
