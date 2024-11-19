GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o product-service
docker build -t shubham047/product-service:latest .
docker tag shubham047/product-service:latest shubham047/product-service:latest

docker push shubham047/product-service:latest