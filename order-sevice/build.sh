GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o order-service
docker build -t shubham047/order-service:latest .
docker tag shubham047/order-service:latest shubham047/order-service:latest

docker push shubham047/order-service:latest