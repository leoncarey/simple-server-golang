
.PHONY: up
up:
	docker-compose up -d --build

.PHONY: down
down:
	docker-compose down --remove-orphans


.PHONY: reload
reload:
	make down
	make up