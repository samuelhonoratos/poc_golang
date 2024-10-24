.PHONY: status up down test

status:
	docker exec -it orcamento-app goose status

up:
	docker exec -it orcamento-app goose up

down:
	docker exec -it orcamento-app goose down

test:
	docker exec -it orcamento-app go test -v ./...
