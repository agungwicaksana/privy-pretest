########################
  CakeStore App v1.0.0
########################


Postman collection: CakeStore.postman_collection.json


========================================================


Docker -- Docker version 20.10.13, build a224086

-- docker build -t agungwicaksana/cakestore:latest .
-- docker compose up

App will be ready on http://localhost:8482


========================================================


Minikube -- minikube version: v1.26.0

Setup MySQL

-- kubectl apply -f kubernetes/mysql/config-map.yml
-- kubectl apply -f kubernetes/mysql/service.yml
-- kubectl apply -f kubernetes/mysql/deployment.yml


Setup cakestore apps
Make sure the docker image is still exists

-- minikube image load agungwicaksana/cakestore:latest
-- kubectl apply -f kubernetes/config-map.yml
-- kubectl apply -f kubernetes/service.yml
-- kubectl apply -f kubernetes/deployment.yml

Make sure all resources are running well
-- kubectl get all

Start tunneling
-- minikube service cakestore-service

Note: Currently can't test the ingress locally because it can't tunnels on M1 machines.

A new browser tab might open. Exposed port is random.


========================================================


Run directly
Setup .env
-- go run main.go