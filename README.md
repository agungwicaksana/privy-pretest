# CakeStore App
#### v.1.0.0

Developed by 

:sunglasses: Agung Wicaksana

:mailbox_with_mail: agung.wicaksaana@gmail.com

---

CakeStore App provides CRUD for cake data.

Postman collection: **CakeStore.postman_collection.json**


### Running on Docker
Docker version 20.10.13, build a224086.
Make sure the docker daemon is running.

```
docker build -t agungwicaksana/cakestore:latest .
docker compose up
```

Apps will be ready on App will be ready on http://localhost:8482

### Running on Minikube
###### Minikube version: v1.26.0

#### 1. Start Minikube
Start minikube if it doesn't started yet
```
minikube start
```

#### 2. Setup MySQL
Apply resources API below.
```
kubectl apply -f kubernetes/mysql/config-map.yml
kubectl apply -f kubernetes/mysql/service.yml
kubectl apply -f kubernetes/mysql/deployment.yml
```

#### 3. Setup CakeStore AppL
Apply resources API below. 

Make sure the docker image 'agungwicaksana/cakestore:latest' is still exists
```
minikube image load agungwicaksana/cakestore:latest
kubectl apply -f kubernetes/mysql/config-map.yml
kubectl apply -f kubernetes/mysql/service.yml
kubectl apply -f kubernetes/mysql/deployment.yml
```

#### 4. Check resources status
Make sure all the resources are running well
``` 
kubectl get all
```

#### 5. Start tunneling
Tunnels kubernetes service to localhost
```
minikube service cakestore-service
```

A new browser tab might open. Exposed port is random.

##### Notes:
Currently can't test the ingress through minikube because it can't tunnels on M1 machines.

### Running directly
Run Cakestore app directly on local machine
#### 1. Setup .env
Copy .env-example and rename it to .env
Fill the env variables based on your needs

#### 2. Downlaod packages
Download packages needed by the app
````
go mod download
go mod tidy
````


#### 3. Run it!
Make sure you are using go 1.18
```
go run main.go
# For hot reloading feature, you could run through nodemon
# You have to install nodemon globally
nodemon --exec go run main.go --signal SIGTERM
```
