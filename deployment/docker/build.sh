#!/bin/bash

set -x
tag="kubelet:latest"
cp ../../_output/kubelet.linux ./
docker build -t beekman9527/$tag .
docker push beekman9527/$tag

