package server

import (
	"BFT-Distributed-G-Set-Remote/config"
	"BFT-Distributed-G-Set-Remote/gset"
	"BFT-Distributed-G-Set-Remote/tools"
	"time"

	zmq "github.com/pebbe/zmq4"
)

type Server struct {
	Zctx           *zmq.Context
	Id             string
	Peers          map[string]*zmq.Socket
	Receive_socket *zmq.Socket
	Host           string
	Port           string
	Gset           map[string]string

	// Used for reliable broadcast
	My_init    map[string]bool
	My_echo    map[string]bool
	My_vote    map[string]bool
	Peers_echo map[string]bool
	Peers_vote map[string]bool

	BRB_start_time         map[string]time.Time
	BRB_MESSAGES           int
	NORMAL_MESSAGES        int
	TOTAL_BRB_TIME         int
	COMPLETED_BRB_REQUESTS int
}

func CreateServer(me config.Node, peers []config.Node) *Server {
	zctx, _ := zmq.NewContext()
	server_sockets := make(map[string]*zmq.Socket)
	my_gset := gset.Create()
	my_init := make(map[string]bool)
	my_echo := make(map[string]bool)
	my_vote := make(map[string]bool)
	peers_echo := make(map[string]bool)
	peers_vote := make(map[string]bool)

	brb_start_time := make(map[string]time.Time)

	my_id := me.Host + ":" + me.Port
	receive_socket, err := zctx.NewSocket(zmq.ROUTER)
	if err != nil {
		panic(err)
	}
	receive_socket.Bind("tcp://*:" + me.Port)
	tools.Log(my_id, "Bound tcp://*:"+me.Port)

	// Connect my dealer sockets to all other servers' router
	for i := 0; i < len(peers); i++ {
		s, err := zctx.NewSocket(zmq.DEALER)
		if err != nil {
			panic(err)
		}
		s.SetIdentity(my_id)
		s.Connect("tcp://" + peers[i].Host + ":" + peers[i].Port)
		tools.Log(my_id, "Connected to "+peers[i].Host+":"+peers[i].Port)
		// append socket to socket list
		server_sockets[peers[i].Host+":"+peers[i].Port] = s
	}

	return &Server{
		Id:             my_id,
		Peers:          server_sockets,
		Receive_socket: receive_socket,
		Host:           me.Host,
		Port:           me.Port,
		Gset:           my_gset,
		My_init:        my_init,
		My_echo:        my_echo,
		My_vote:        my_vote,
		Peers_echo:     peers_echo,
		Peers_vote:     peers_vote,

		BRB_start_time:         brb_start_time,
		BRB_MESSAGES:           0,
		NORMAL_MESSAGES:        0,
		TOTAL_BRB_TIME:         0,
		COMPLETED_BRB_REQUESTS: 0,
	}
}
