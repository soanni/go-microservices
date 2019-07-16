build:
	GOOS=linux GOARCH=amd64 go build -o user-cli
	docker build --no-cache -t user-cli-micro .

run:
	# docker ps -aq -f name=user-cli-micro | xargs docker rm -f
	docker run --name="user-cli-micro" -d user-cli-micro