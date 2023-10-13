include .env
export $(shell sed 's/=.*//' .env)

debug:
	docker compose -f deployment/docker-compose.yaml up mongo -d 
	go run cmd/server/main.go
	
deploy:
	docker compose -f deployment/docker-compose.yaml up --build -d

prod:
	ssh -i $(SSH_KEY) $(SERVER_USERNAME)@$(SERVER_URL) "cd more-tech && git pull && docker compose -f deployment/docker-compose.yaml up -d --build"
