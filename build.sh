#!/usr/bin/env bash
CGO_ENABLED=0 GOOS=linux go build -ldflags '-d -w -s ' -a -installsuffix cgo -o adapter .
docker build -t weibh/k8s-test-metrics-adapter:latest .
docker push weibh/k8s-test-metrics-adapter:latest
kubectl scale deploy custom-metrics-apiserver -n custom-metrics --replicas=0
kubectl scale deploy custom-metrics-apiserver -n custom-metrics --replicas=1
