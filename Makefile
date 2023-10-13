include .env
export $(shell sed 's/=.*//' .env)

# запуск локально, без докера
debug:
	docker compose -f deployment/docker-compose.yaml up mongo -d 
	go run cmd/server/main.go
	
# локально, с докером
local:
	docker compose -f deployment/docker-compose.yaml up --build -d

stop_local:
	docker compose -f deployment/docker-compose.yaml down

# запуск на сервере, с докером
deploy:
	ssh -i $(SSH_KEY) $(SERVER_USERNAME)@$(SERVER_URL) "cd more-tech && git pull && docker compose -f deployment/docker-compose.yaml up -d --build"

stop_deploy:
	ssh -i $(SSH_KEY) $(SERVER_USERNAME)@$(SERVER_URL) "cd more-tech && docker compose -f deployment/docker-compose.yaml down"

# копируем монгу на сервер
copy_db:
	scp -r $(shell pwd)/data $(SERVER_USERNAME)@$(SERVER_URL):/home/$(SERVER_USERNAME)/more-tech/data 