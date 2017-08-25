package main

import (
	"flag"
	"github.com/golang/glog"
)

var (
	kubeletHost = "localhost"
	masterUrl = ""
	kubeConfig = ""
)

func process_flags(c *kubeletConfig) {
	flag.StringVar(&masterUrl, "masterUrl", "", "kubernetes master url")
	flag.StringVar(&kubeConfig, "kubeConfig", "", "absolute path to the kubeconfig file.")
	flag.IntVar(&(c.timeout), "timeout", c.timeout, "the timeout when connecting to kubelet to get metrics.")
	flag.IntVar(&(c.port), "kubeletPort", c.port, "the Port of kubelet to get metrics.")
	flag.BoolVar(&(c.enableHttps), "kubeletHttps", c.enableHttps, "Access the kubelet with https or http.")
	flag.StringVar(&kubeletHost, "host", kubeletHost, "the host of the kubelet")

	flag.Set("alsologtostderr", "true")
	flag.Parse()
}

func testHost(client *KubeletClient, host string) {
	client.GetMachineInfo(host)
	client.GetSummary(host)
}

func main() {
	config := &kubeletConfig{
		enableHttps: true,
		port:        defaultKubeletPort,
		timeout:     defaultTimeOut,
	}
	process_flags(config)

	glog.V(2).Infof("hello")
	restConfig := GetKubeConfig(masterUrl, kubeConfig)
	if restConfig == nil {
		glog.Errorf("failed to get kubeClient config.")
		return
	}

	client, err := NewKubeletClient(config, restConfig)
	if err != nil {
		glog.Errorf("failed to create a kubelet client: %v", err)
		return
	}

	testHost(client, kubeletHost)
}
