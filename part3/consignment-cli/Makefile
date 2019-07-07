build:
	GOOS=linux GOARCH=amd64 go build -o consignment-cli
	docker rmi consignment-cli-micro
	docker build --no-cache -t consignment-cli-micro .

run:
	docker ps -aq -f name=consignment-cli-micro | xargs docker rm -f
	docker run --name="consignment-cli-micro" -d consignment-cli-micro

logs:
	docker ps -aq -f name=consignment-cli-micro | xargs docker logs