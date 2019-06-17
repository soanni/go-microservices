build:
	protoc -I. --go_out=plugins=grpc:. \
		proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-service .

run:
	docker run --name="consignment-service" -d -p 50051:50051 consignment-service