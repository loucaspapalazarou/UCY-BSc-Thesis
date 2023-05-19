#!/bin/bash

ansible-playbook -i ./hosts end.yml
ansible-playbook -i ./hosts start.yml -v -e "requests=10000"
