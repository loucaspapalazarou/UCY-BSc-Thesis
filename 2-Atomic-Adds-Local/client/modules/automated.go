package modules

import (
	"frontend/client"
	"frontend/config"
	"frontend/messaging"
	"frontend/tools"
	"strconv"
	"sync"

	zmq "github.com/pebbe/zmq4"
)

func StartAutomated(zctx *zmq.Context, client_count, request_count int, network_name string) {
	var wg sync.WaitGroup
	wg.Add(client_count)
	for i := 0; i < client_count; i++ {
		id := "c" + strconv.Itoa(i)
		go func(id string) {
			tools.Log(id, "Id set")
			config.Initialize(network_name)
			servers := config.SERVERS
			client := client.CreateClient(id, servers, zctx)
			for r := 0; r < request_count; r++ {
				messaging.Add(client, id+"-test-"+strconv.Itoa(r))
				messaging.Get(client)
				// tools.Log(client.Id, r)
			}
			tools.Log(id, "Done")
			wg.Done()
		}(id)
	}
	wg.Wait()
}

func StartAutomatedAtomicAdds(zctx *zmq.Context, request_count int) {
	var wg sync.WaitGroup
	wg.Add(2)

	// peer_id;destination;your_message;peer_message
	go atomicAddsWorker("loukas", "hello", "marios", "world", "bdso-1", zctx, request_count, &wg)
	go atomicAddsWorker("marios", "world", "loukas", "hello", "bdso-2", zctx, request_count, &wg)

	wg.Wait()
}

func atomicAddsWorker(my_id, my_message, peer_id, peer_message, dest string, zctx *zmq.Context, request_count int, wg *sync.WaitGroup) {
	config.Initialize("sbdso")
	servers := config.SERVERS
	client := client.CreateClient(my_id, servers, zctx)
	for r := 0; r < request_count; r++ {
		message := peer_id + ";" + dest + ";" + my_message + "_" + strconv.Itoa(r) + ";" + peer_message + "_" + strconv.Itoa(r)
		messaging.AddAtomic(client, message)
		messaging.Get(client)
		// tools.Log(client.Id, r)
	}
	tools.Log(my_id, "Done")
	wg.Done()
}
