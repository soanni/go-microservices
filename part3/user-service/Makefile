build:
	protoc -I. --go_out=plugins=micro:. \
		proto/user/user.proto
	GOOS=linux GOARCH=amd64 go build -o user-service -a -installsuffix cgo main.go repository.go handler.go
	# docker rmi user-service-micro
	docker build --no-cache -t user-service-micro .

run:
	docker ps -aq -f name=user-service-micro | xargs docker rm -f
	docker run --name="user-service-micro" -d -e MICRO_SERVER_ADDRESS=:50051 -p 50053:50051 user-service-micro