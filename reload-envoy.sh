kubectl -n demo delete configmap envoy-config
sleep 5

kubectl apply -f envoy-config-blocked.yaml
sleep 5

kubectl -n demo rollout restart deployment/envoy-proxy

sleep 5
kubectl -n demo get pods

# add sleep for 5 seconds
sleep 10

kubectl -n demo logs -l app=envoy-proxy -f