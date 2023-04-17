package modules

import (
	"2-Atomic-Adds/client"
	"2-Atomic-Adds/config"
	"2-Atomic-Adds/messaging"
	"2-Atomic-Adds/tools"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	zmq "github.com/pebbe/zmq4"
)

func StartAutomated(zctx *zmq.Context, client_count, request_count int, network_name string) {
	servers := config.Initialize(network_name)
	var wg sync.WaitGroup
	wg.Add(client_count)
	for i := 0; i < client_count; i++ {
		host, err := os.Hostname()
		if err != nil {
			panic(err)
		}
		host = strings.Split(host, ".")[0]
		id := host + "_client_" + strconv.Itoa(i)
		go func(id string) {
			tools.Log(id, "Id set")
			config.Initialize(network_name)
			client := client.CreateClient(id, servers, zctx)
			for r := 0; r < request_count; r++ {
				messaging.Add(client, id+"-test-"+strconv.Itoa(r))
				messaging.Get(client)
			}
			tools.Log(id, "Done")
			wg.Done()
		}(id)
	}
	wg.Wait()
}

// Starts just one client with a base message
func StartAutomatedAtomic(my_id, my_message, peer_id, peer_message, dest string, zctx *zmq.Context, request_count int) {
	servers := config.Initialize("sbdso")
	tools.Log(my_id, "Id set")
	client := client.CreateClient(my_id, servers, zctx)
	for r := 0; r < request_count; r++ {
		message := peer_id + ";" + dest + ";" + my_message + "_" + strconv.Itoa(r) + ";" + peer_message + "_" + strconv.Itoa(r)
		add_atomic_time := messaging.AddAtomic(client, message)
		_, get_time := messaging.Get(client)
		fmt.Println(add_atomic_time)
		s := tools.Stats{
			TOTAL_GET_TIME:        client.TOTAL_GET_TIME,
			TOTAL_ADD_TIME:        client.TOTAL_ADD_TIME,
			TOTAL_ADD_ATOMIC_TIME: client.TOTAL_ADD_ATOMIC_TIME,
			REQUESTS:              client.REQUESTS,
		}
		client.TOTAL_ADD_ATOMIC_TIME, client.REQUESTS = tools.IncrementAddAtomicTime(client.Id, add_atomic_time, s)
		client.TOTAL_GET_TIME, client.REQUESTS = tools.IncrementGetTime(client.Id, get_time, s)
	}
	tools.Log(my_id, "Done")
}
