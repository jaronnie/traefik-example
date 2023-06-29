deploy:
	@GOOS=linux GOARCH=amd64 go build -o helloworld
	@docker-compose up -d --build --remove-orphans
	@rm helloworld