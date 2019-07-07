build:
	protoc -I. --go_out=plugins=micro:. \
		proto/vessel/vessel.proto
	GOOS=linux GOARCH=amd64 go build -o vessel-service -a -installsuffix cgo main.go repository.go handler.go datastore.go
	docker rmi vessel-service-micro
	docker build --no-cache -t vessel-service-micro .

run:
	# docker rm -f vessel-service-micro
	docker run --name="vessel-service-micro" -d -e MICRO_SERVER_ADDRESS=:50051 -p 50052:50051 vessel-service-micro