# Golang hello app

### To run app locally
```
export APP_NAME="backend"
export APP_VERSION="v1"
export APP_INSTANCE="backend-v1-aksed"
export URL=""
export CONN_PORTS="8080,8081"
go run main.go
```

### To build and push app docker image
```
export REPO_PATH=""
docker build -t experiments/golang-hello-app .
docker tag experiments/golang-hello-app "${REPO_PATH}"
docker push "${REPO_PATH}"
```

### To run docker image locally

Setup:
```
docker network create golang-hello-app-network
```

Run:
```
docker run -it --rm --name golang-hello-app \
	--network golang-hello-app-network \
	-p 8080:8080 \
	-p 8081:8081 \
	-e CONN_PORTS="8080,8081" \
	-e APP_NAME="backend" \
	-e APP_VERSION="v1" \
	-e APP_INSTANCE="backend-v1-aksed" \
	-e URL="" \
	experiments/golang-hello-app
```

Cleanup:
```
docker network rm golang-hello-app-network
```