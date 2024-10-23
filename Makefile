DATABASE_URL := "root:health@tcp(orcamento_db:3306)/orcamento_db?charset=utf8&parseTime=True&loc=Local"

.PHONY: migration-status migrate-up migrate-down test

migration-status:
	docker exec -it orcamento-app goose -dir di/database/migrations mysql $(DATABASE_URL) status

migrate-up:
	docker exec -it orcamento-app goose -dir di/database/migrations mysql $(DATABASE_URL) up

migrate-down:
	docker exec -it orcamento-app goose -dir di/database/migrations mysql $(DATABASE_URL) down

test:
	docker exec -it orcamento-app go test -v ./...


