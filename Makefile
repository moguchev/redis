.PHONY: up
up:
	docker-compose build
	docker-compose up -d redis

.PHONY: down
down:
	docker-compose down