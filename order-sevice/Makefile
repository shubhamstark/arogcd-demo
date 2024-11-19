



build:
	GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o order-service
    docker build -t $(DOCKER_USER)/$(IMAGE_NAME):$(TAG) .
    docker push $(DOCKER_USER)/$(IMAGE_NAME):$(TAG)