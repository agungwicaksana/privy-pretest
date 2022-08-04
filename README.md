########################
  CakeStore App v1.0.0
########################


Postman collection: 


========================================================


Docker

-- docker build -t agungwicaksana/cakestore:latest .
-- docker compose up

App will be ready on http://localhost:8482


========================================================


Minikube

-- kubectl apply -f kubernetes/deployment.yml

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

Start tunneling
-- minikube service cakestore-service

A new browser tab might open. Exposed port is random.


========================================================


Run directly
-- go run main.go