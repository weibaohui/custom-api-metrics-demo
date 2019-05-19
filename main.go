package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/component-base/logs"
	"k8s.io/klog"


	httprequest "github.com/kubernetes-incubator/custom-metrics-apiserver/http-request/pkg"
	basecmd "github.com/kubernetes-incubator/custom-metrics-apiserver/pkg/cmd"
	"github.com/kubernetes-incubator/custom-metrics-apiserver/pkg/provider"
)

type HttpRequestAdapter struct {
	basecmd.AdapterBase

	// Message is printed on succesful startup
	Message string
}
func (a *HttpRequestAdapter) makeProviderOrDie() provider.CustomMetricsProvider {
	client, err := a.DynamicClient()
	if err != nil {
		klog.Fatalf("unable to construct dynamic client: %v", err)
	}

	mapper, err := a.RESTMapper()
	if err != nil {
		klog.Fatalf("unable to construct discovery REST mapper: %v", err)
	}

	return httprequest.NewProvider(client, mapper)
}


func main() {
	fmt.Println("HttpRequestAdapterHttpRequestAdapterHttpRequestAdapterHttpRequestAdapterHttpRequestAdapter")
	fmt.Println("HttpRequestAdapterHttpRequestAdapterHttpRequestAdapterHttpRequestAdapterHttpRequestAdapter")
	fmt.Println("HttpRequestAdapterHttpRequestAdapterHttpRequestAdapterHttpRequestAdapterHttpRequestAdapter")
	fmt.Println("HttpRequestAdapterHttpRequestAdapterHttpRequestAdapterHttpRequestAdapterHttpRequestAdapter")
	fmt.Println("HttpRequestAdapterHttpRequestAdapterHttpRequestAdapterHttpRequestAdapterHttpRequestAdapter")
	fmt.Println("HttpRequestAdapterHttpRequestAdapterHttpRequestAdapterHttpRequestAdapterHttpRequestAdapter")
	fmt.Println("HttpRequestAdapterHttpRequestAdapterHttpRequestAdapterHttpRequestAdapterHttpRequestAdapter")
	fmt.Println("HttpRequestAdapterHttpRequestAdapterHttpRequestAdapterHttpRequestAdapterHttpRequestAdapter")
	fmt.Println("HttpRequestAdapterHttpRequestAdapterHttpRequestAdapterHttpRequestAdapterHttpRequestAdapter")
	fmt.Println("HttpRequestAdapterHttpRequestAdapterHttpRequestAdapterHttpRequestAdapterHttpRequestAdapter")
	logs.InitLogs()
	defer logs.FlushLogs()

	cmd := &HttpRequestAdapter{}
	cmd.Flags().StringVar(&cmd.Message, "msg", "starting adapter...", "startup message")
	cmd.Flags().AddGoFlagSet(flag.CommandLine) // make sure we get the klog flags
	cmd.Flags().Parse(os.Args)

	metricsProvider := cmd.makeProviderOrDie()
	cmd.WithCustomMetrics(metricsProvider)

	klog.Infof(cmd.Message)

	go func() {
		// Open port for POSTing fake metrics
		klog.Fatal(http.ListenAndServe(":8080", nil))
	}()
	if err := cmd.Run(wait.NeverStop); err != nil {
		klog.Fatalf("unable to run custom metrics adapter: %v", err)
	}
}
