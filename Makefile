debug:
	docker compose -f deployment/docker-compose.yaml up mongo -d 
	go run cmd/server/main.go

deploy:
	docker compose -f deployment/docker-compose.yaml up --build -d