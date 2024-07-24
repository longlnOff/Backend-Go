postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

start_postgres12:
	docker start postgres12

stop_postgres12:
	docker stop postgres12

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc: 
	sqlc generate
test: 
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -destination db/mock/Store.go -package mockdb simplebank/db/sqlc Store

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1



.PHONY: postgres createdb dropdb migrateup migratedown
.PHONY: sqlc test server mock migratedown1 migrateup1 
.PHONY: start_postgres12 stop_postgres12