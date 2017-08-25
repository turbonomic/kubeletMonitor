#!/bin/bash
# get monitor information from a Openshift cluster

mybin=./_output/kubelet

kubeconfig="./config/en119.kubeconfig.yaml"
host="enmaster.octurbo.org"
masterUrl="https://10.10.173.119:8443"

options=" --v=3"
options="$options --kubeletPort=10250"
options="$options --kubeletHttps=true"
options="$options --host=$host"
options="$options --kubeConfig=$kubeconfig"
#options="$options --masterUrl=$masterUrl"

echo "$mybin $options"
$mybin $options 
