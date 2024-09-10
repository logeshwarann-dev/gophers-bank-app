postgres:
	docker run -d --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root gopher_bank
dropdb:
	docker exec -it postgres12 dropdb gopher_bank

migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/gopher_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/gopher_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc