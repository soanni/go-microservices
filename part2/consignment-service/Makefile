build:
	protoc -I. --go_out=plugins=micro:. \
		proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 go build
	docker build --no-cache -t consignment-service-micro .

run:
	docker ps -aq -f name=consignment-service-micro | xargs docker rm -f
	docker run --name="consignment-service-micro" -d -e MICRO_SERVER_ADDRESS=:50051 -p 50051:50051 consignment-service-micro