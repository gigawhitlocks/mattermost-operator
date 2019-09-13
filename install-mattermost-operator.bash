#!/bin/bash
set -ex

kind delete cluster || true
kind create cluster --config ./test/kind-config-local.yaml --wait 5m
export KUBECONFIG="$(kind get kubeconfig-path --name="kind")"

kubectl delete storageclasses.storage.k8s.io --all
kubectl apply -f ./test/local-path-provisioner.yaml

kubectl create ns mysql-operator
kubectl apply -n mysql-operator -f https://raw.githubusercontent.com/mattermost/mattermost-operator/master/docs/mysql-operator/mysql-operator.yaml

kubectl create ns minio-operator
kubectl apply -n minio-operator -f https://raw.githubusercontent.com/mattermost/mattermost-operator/master/docs/minio-operator/minio-operator.yaml

kubectl create ns mattermost-operator
make build-image
kind load docker-image mattermost/mattermost-operator:test
kubectl apply -n mattermost-operator -f <(cat ./docs/mattermost-operator/mattermost-operator.yaml | sed -e 's|image: mattermost/mattermost-operator:latest|image: mattermost/mattermost-operator:test|g')
#kubectl apply -n mattermost-operator -f https://raw.githubusercontent.com/mattermost/mattermost-operator/master/docs/mattermost-operator/mattermost-operator.yaml

# nginx
kubectl create ns ingress-nginx
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/mandatory.yaml
kubectl patch deployment -n ingress-nginx nginx-ingress-controller -p "$(cat test/kind-nginx-ingress.yaml)"

kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/cloud-generic.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/baremetal/service-nodeport.yaml

# mattermost
kubectl apply -f <(sed -e 's/mattermost.example.com/mattermost.localhost/g' docs/examples/simple_anywhere.yaml)

# add the mattermost hostname to /etc/hosts
grep mattermost.localhost /etc/hosts || echo "127.0.0.1 mattermost.localhost" | sudo tee -a /etc/hosts
