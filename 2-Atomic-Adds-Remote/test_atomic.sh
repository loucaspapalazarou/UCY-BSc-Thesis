#!/bin/bash

requests=10000

function run_test {
    ./start.sh
    ssh node1 "cd /users/loukis/Thesis/2-Atomic-Adds-Remote/client && /usr/local/go/bin/go run . -auto_atomic -my_id loukas -peer_id marios -my_msg hello -peer_msg world -dest bdso-1 -reqs $requests" &
    ssh node2 "cd /users/loukis/Thesis/2-Atomic-Adds-Remote/client && /usr/local/go/bin/go run . -auto_atomic -my_id marios -peer_id loukas -my_msg world -peer_msg hello -dest bdso-2 -reqs $requests" &
}

export -f run_test 
rm logfile
nohup bash -c run_test > logfile 2>&1 &

