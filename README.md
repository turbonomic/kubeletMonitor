# kubeletMonitor
A kubelet client to get monitor information (node.info, node resource usage, container resource usage).

## Differences between Openshift and Standard Kubernetes
|-|Kubernetes|Openshfit|
|-|-|-|
|kubeletPort| 10255 | 10250|
|http.scheme| HTTP | HTTPS|

Note: Openshift version is 3.4, Kubernetes is 1.7.

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
