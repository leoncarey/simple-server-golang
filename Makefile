
.PHONY: up
up:
	docker-compose up --build

.PHONY: down
down:
	docker-compose down --remove-orphans


.PHONY: reload
reload:
	make down
	make up

.PHONY: run
run:
	docker-compose build --pull
	docker-compose run --publish 8000:8080 app