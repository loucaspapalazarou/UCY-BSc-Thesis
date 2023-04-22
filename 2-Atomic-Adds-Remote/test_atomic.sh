#!/bin/bash

run_test() {
    requests=$1
    output_file1="output_node1_${requests}.txt"
    output_file2="output_node2_${requests}.txt"
    ./start.sh
    ssh node1 "cd /users/loukis/Thesis/2-Atomic-Adds-Remote/client && /usr/local/go/bin/go run . -auto_atomic -my_id loukas -peer_id marios -my_msg hello -peer_msg world -dest bdso-1 -reqs $requests" &
    ssh node2 "cd /users/loukis/Thesis/2-Atomic-Adds-Remote/client && /usr/local/go/bin/go run . -auto_atomic -my_id marios -peer_id loukas -my_msg world -peer_msg hello -dest bdso-2 -reqs $requests" &
    wait
    scp loukis@node1:~/Thesis/2-Atomic-Adds-Remote/client/scenario_results* "/users/loukis/Thesis/2-Atomic-Adds-Remote/results/output_node1_${requests}.txt"
    scp loukis@node2:~/Thesis/2-Atomic-Adds-Remote/client/scenario_results* "/users/loukis/Thesis/2-Atomic-Adds-Remote/results/output_node2_${requests}.txt"
    # copy from servers
    scp loukis@node3:~/Thesis/2-Atomic-Adds-Remote/sbdso/scenario_results* "/users/loukis/Thesis/2-Atomic-Adds-Remote/results/output_node3_${requests}.txt"
    scp loukis@node4:~/Thesis/2-Atomic-Adds-Remote/sbdso/scenario_results* "/users/loukis/Thesis/2-Atomic-Adds-Remote/results/output_node4_${requests}.txt"
    scp loukis@node5:~/Thesis/2-Atomic-Adds-Remote/sbdso/scenario_results* "/users/loukis/Thesis/2-Atomic-Adds-Remote/results/output_node5_${requests}.txt"
    scp loukis@node6:~/Thesis/2-Atomic-Adds-Remote/sbdso/scenario_results* "/users/loukis/Thesis/2-Atomic-Adds-Remote/results/output_node6_${requests}.txt"
    scp loukis@node7:~/Thesis/2-Atomic-Adds-Remote/sbdso/scenario_results* "/users/loukis/Thesis/2-Atomic-Adds-Remote/results/output_node7_${requests}.txt"
}

export -f run_test

rm -rf results
mkdir results

for requests in 100 200 500 1000 2000 5000; do
    echo "Running test with $requests requests"
    run_test $requests
done