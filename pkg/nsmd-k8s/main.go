package main

import (
	"github.com/dtornow/cnns-nsr/k8stools"
	"github.com/dtornow/cnns-nsr/nsrlogger"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/flowcontrol"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	nsrlogger.LogInfoLevel("Starting nsmd-k8s")
	c := make(chan os.Signal, 1)
	signal.Notify(c,
		os.Interrupt,
		// More Linux signals here
		syscall.SIGHUP,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	address := os.Getenv("NSMD_K8S_ADDRESS")
	if strings.TrimSpace(address) == "" {
		address = "0.0.0.0:5000"
	}

	nsmName, ok := os.LookupEnv("NODE_NAME")

	if !ok {
		nsrlogger.LogInfoLevel("Not ok.")
	}

	nsrlogger.LogInfoLevel("Starting NSMD Kubernetes on " + address + " with NsmName " + nsmName)

	configPath := k8stools.GetK8sConfigPath()

	config, err := rest.InClusterConfig()
	if err != nil {
		nsrlogger.LogInfoLevel("Unable to get in cluster config %v, , attempting to fall back to kubeconfig", err)
		config, err = clientcmd.BuildConfigFromFlags("", configPath)
		if err != nil {
			nsrlogger.LogFatalLevel("%v", err)
		}
	}

	configShallowCopy := *config
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
}
