#!/bin/bash

mybin=./_output/kubelet

## get monitor information from Kubernetes cluster

kubeconfig="./config/aws.kubeconfig.yaml"
host="10.10.172.33"
masterUrl="https://x.x.x.x:6443"

options=" --v=3"
options="$options --kubeletPort=10255"
options="$options --kubeletHttps=false"
options="$options --host=$host"
options="$options --kubeConfig=$kubeconfig"
#options="$options --masterUrl=$masterUrl"


echo "$mybin $options"
$mybin $options 
