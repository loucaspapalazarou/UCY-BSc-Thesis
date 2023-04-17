#!/bin/bash

./start.sh

go run . -auto_atomic -my_id loukas -peer_id marios -my_msg hello -peer_msg world -dest bdso-1 -reqs 100
go run . -auto_atomic -my_id marios -peer_id loukas -my_msg world -peer_msg hello -dest bdso-2 -reqs 100