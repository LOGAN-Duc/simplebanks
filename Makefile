postgres:
	docker run --name postgres2411 --network bank-network -p 2411:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createdb:
	 docker exec -it postgres2411 createdb --username=root --owner=root simple_banks
dropdb:
	 docker exec -it postgres2411 dropdb simple_banks
migrateup:
	 migrate -path db/migration -database "postgresql://root:secret@localhost:2411/simple_banks?sslmode=disable" -verbose up
migrateup1:
	 migrate -path db/migration -database "postgresql://root:secret@localhost:2411/simple_banks?sslmode=disable" -verbose up 1
migratedown:
	 migrate -path db/migration -database "postgresql://root:secret@localhost:2411/simple_banks?sslmode=disable" -verbose down
migratedown1:
	migrate -path ./db/migration -database "postgresql://root:secret@localhost:2411/simple_banks?sslmode=disable" -verbose down 1
new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go simplebanks/db/sqlc Store


.PHONY: new_migration db_schema postgres createdb dropdb migratedown migrateup sqlc	test mock migrateup1 migratedown1 proto evans redis
