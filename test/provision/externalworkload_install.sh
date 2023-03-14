#!/bin/bash
set -e

cd /home/vagrant/go/src/github.com/go-faster/cilium

# Build docker image
make docker-cilium-image

CLUSTER_ADDR=192.168.56.11:32379 HOST_IP=192.168.56.10 CILIUM_IMAGE=go-faster/cilium:latest contrib/k8s/install-external-workload.sh
