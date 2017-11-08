# kubeletMonitor
A kubelet client to get monitor information (node.info, node resource usage, container resource usage).

## Differences between Openshift and Standard Kubernetes
|-|Kubernetes|Openshfit|GKE|
|-|-|-|-|
|kubeletPort| 10255 | 10250|10255|
|http.scheme| HTTP | HTTPS| HTTP|

Note1: Openshift version is 3.4, Kubernetes is 1.7;

Note2: GKE (google container engine) 1.6+ use [`http and port 10255`](https://github.com/prometheus/prometheus/issues/2606).

## Run it
### Get kubelet monintor information from OpenShift
```console
$ make build
$ ./_output/kubelet  --v=3 --kubeletPort=10250 --kubeletHttps=true --host=<nodeIP> --kubeConfig=./config/openshift.kubeconfig.yaml

# or

# sh script/osrun.sh
```

#### Get kubelet monintor information from Kubernets
```console
$ make build
$ ./_output/kubelet  --v=3 --kubeletPort=10255 --kubeletHttps=false --host=<nodeIP> --kubeConfig=./config/kube.kubeconfig.yaml

# or

# sh script/run.sh
```
