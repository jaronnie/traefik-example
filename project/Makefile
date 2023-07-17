deploy:
	@GOOS=linux GOARCH=amd64 go build -o ./apps/app1/app1 ./apps/app1
	@GOOS=linux GOARCH=amd64 go build -o ./apps/app2/app2 ./apps/app2
	@GOOS=linux GOARCH=amd64 go build -o ./apps/app3/app3 ./apps/app3
	@docker-compose down	
	@docker-compose up -d --build --remove-orphans
	@rm ./apps/app1/app1
	@rm ./apps/app2/app2
	@rm ./apps/app3/app3