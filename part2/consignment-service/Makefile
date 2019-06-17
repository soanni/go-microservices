build:
	protoc -I. --go_out=plugins=micro:. \
		proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-service-micro .

run:
	docker run --name="consignment-service-micro" -d -p 50051:50051 consignment-service-micro