#!/bin/bash

echo Enter Namespace

read namespace

check=$(kubectl get namespaces | awk {'print $1'} | column -t)
# echo Check [$namespace] in [$check]

if [ -z "$namespace" ]; then
    echo Please enter a Namespace...
    exit 1
elif [[ $check =~ $namespace ]]; then
    echo Namespace already exists...
    exit 1
else
    kubectl create namespace $namespace
    echo 
    echo -----Namespace $namespace has been created-----
    echo
    kubectl create -f . -n $namespace
    echo
    echo -----All Services and Deployments have been created-----
    echo
    echo -----Waiting 15 seconds for pods to spin up-----
    echo
    echo -----Pods will be running in a while! Thank you for your patience! -----
    echo
    
    sleep 10
    address=$(kubectl get svc -n $namespace exercise-5-service -o jsonpath='{.status.loadBalancer.ingress[0].hostname}{"\n"}')
    
    echo -----To get the host address, run the following command in some time-----
    echo 
    echo host $address
    echo
    exit 0
fi

