docker build -t cx-mp-gateway .
docker tag cx-mp-gateway registry.local.com/cx-mp-gateway
docker push registry.local.com/cx-mp-gateway
kubectl delete deployment cx-mp-gateway -n cx-rpc-business
kubectl apply -f cx-mp-gateway-deployment.yaml
docker rmi cx-mp-gateway
docker rmi registry.local.com/cx-mp-gateway
