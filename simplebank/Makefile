postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root bank
dropdb:
	docker exec -it postgres12 dropdb bank

startdocker_db:
	docker start postgres12

stopdocker_db:
	docker stop postgres12

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration/ -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration/ -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -destination db/mock/store.go -package mockdb simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test
.PHONY: startdocker_db stopdocker_db server mock migratedown1 migrateup1
