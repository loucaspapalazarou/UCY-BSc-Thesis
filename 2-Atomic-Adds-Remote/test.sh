#!/bin/bash    
    
./start.sh
ssh node1 "cd /users/loukis/Thesis/2-Atomic-Adds-Remote/client && /usr/local/go/bin/go run . -auto_atomic -my_id loukas -peer_id marios -my_msg hello -peer_msg world -dest bdso-1 -reqs 200" &
ssh node2 "cd /users/loukis/Thesis/2-Atomic-Adds-Remote/client && /usr/local/go/bin/go run . -auto_atomic -my_id marios -peer_id loukas -my_msg world -peer_msg hello -dest bdso-2 -reqs 200" &
wait

echo "Done"