.PHONY: up
up:
	docker-compose build --pull
	docker-compose run --publish 8000:8080 app


.PHONY: down
down:
	docker-compose down --remove-orphans