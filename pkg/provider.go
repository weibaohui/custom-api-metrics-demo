package provider

import (
	"fmt"
	"github.com/kubernetes-incubator/custom-metrics-apiserver/pkg/provider"
	"github.com/kubernetes-incubator/custom-metrics-apiserver/pkg/provider/helpers"
	apimeta "k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/metrics/pkg/apis/custom_metrics"
	"math/rand"
	"time"
)

type httpRequestProvider struct {
	client dynamic.Interface
	mapper apimeta.RESTMapper

	values map[provider.CustomMetricInfo]int64
}

func NewProvider(client dynamic.Interface, mapper apimeta.RESTMapper) provider.CustomMetricsProvider {
	return &httpRequestProvider{
		client: client,
		mapper: mapper,
		values: make(map[provider.CustomMetricInfo]int64),
	}
}

func (h *httpRequestProvider) ListAllMetrics() []provider.CustomMetricInfo {
	fmt.Println("ListAllMetrics")
	fmt.Println("ListAllMetrics")
	fmt.Println("ListAllMetrics")
	fmt.Println("ListAllMetrics")
	fmt.Println("ListAllMetrics")

	return []provider.CustomMetricInfo{
		// these are mostly arbitrary examples
		{
			GroupResource: schema.GroupResource{Group: "", Resource: "pods"},
			Metric:        "http_requests",
			Namespaced:    true,
		},
		{
			GroupResource: schema.GroupResource{Group: "", Resource: "services"},
			Metric:        "http_requests",
			Namespaced:    true,
		},
		{
			GroupResource: schema.GroupResource{Group: "", Resource: "namespaces"},
			Metric:        "http_requests",
			Namespaced:    false,
		},
	}

}
func (h *httpRequestProvider) GetMetricBySelector(namespace string, selector labels.Selector, info provider.CustomMetricInfo) (*custom_metrics.MetricValueList, error) {
	totalValue, err := h.valueFor(info)
	if err != nil {
		return nil, err
	}

	names, err := helpers.ListObjectNames(h.mapper, h.client, namespace, selector, info)
	if err != nil {
		return nil, err
	}

	fmt.Println("GetMetricBySelector", names)
	if selector != nil {

		fmt.Println("GetMetricBySelector", names, selector.String())
	}

	fmt.Println("GetMetricBySelector", info.String())

	res := make([]custom_metrics.MetricValue, len(names))
	for i, name := range names {
		// in a real adapter, you might want to consider pre-computing the
		// object reference created in metricFor, instead of recomputing it
		// for each object.
		value, err := h.metricFor(100*totalValue/int64(len(res)), types.NamespacedName{Namespace: namespace, Name: name}, info)
		if err != nil {
			return nil, err
		}
		res[i] = *value
	}

	return &custom_metrics.MetricValueList{
		Items: res,
	}, nil
}

func (h *httpRequestProvider) GetMetricByName(name types.NamespacedName, info provider.CustomMetricInfo) (*custom_metrics.MetricValue, error) {
	fmt.Println("GetMetricByName", name.Namespace, name.Name, info.String())
	fmt.Println("GetMetricByName", name.Namespace, name.Name, info.String())
	fmt.Println("GetMetricByName", name.Namespace, name.Name, info.String())
	fmt.Println("GetMetricByName", name.Namespace, name.Name, info.String())
	fmt.Println("GetMetricByName", name.Namespace, name.Name, info.String())
	fmt.Println("GetMetricByName", name.Namespace, name.Name, info.String())
	value, err := h.valueFor(info)
	if err != nil {
		return nil, err
	}

	return h.metricFor(value, name, info)
}

// valueFor fetches a value from the fake list and increments it.
func (h *httpRequestProvider) valueFor(info provider.CustomMetricInfo) (int64, error) {
	// normalize the value so that you treat plural resources and singular
	// resources the same (e.g. pods vs pod)
	info, _, err := info.Normalized(h.mapper)
	if err != nil {
		return 0, err
	}

	value := h.values[info]
	value = rand.Int63() + 10
	if info.Metric == "tcp_conns" {
		value = 2
	}
	h.values[info] = value
	fmt.Println("value",value)
	fmt.Println("info.Metric", info.Metric)
	fmt.Println("info.GroupResource", info.GroupResource)
	fmt.Println("valueFor ", value, info.Metric, info.GroupResource, info.Namespaced, info.String())
	fmt.Println("valueFor ", value, info.Metric, info.GroupResource, info.Namespaced, info.String())
	fmt.Println("valueFor ", value, info.Metric, info.GroupResource, info.Namespaced, info.String())
	fmt.Println("valueFor ", value, info.Metric, info.GroupResource, info.Namespaced, info.String())
	return value, nil
}

func (p *httpRequestProvider) metricFor(value int64, name types.NamespacedName, info provider.CustomMetricInfo) (*custom_metrics.MetricValue, error) {
	// construct a reference referring to the described object
	objRef, err := helpers.ReferenceFor(p.mapper, name, info)
	if err != nil {
		return nil, err
	}

	return &custom_metrics.MetricValue{

		DescribedObject: objRef,
		Metric: custom_metrics.MetricIdentifier{
			Name: info.Metric,
		},
		// you'll want to use the actual timestamp in a real adapter
		Timestamp:     metav1.Time{time.Now()},
		WindowSeconds: nil,
		Value:         *resource.NewMilliQuantity(value*100, resource.DecimalSI),
	}, nil
}
