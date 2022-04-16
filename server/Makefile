run:
	go run cmd/main.go

test: unit integration e2e

unit:
	go test -tags=unit -v -cover ./...

integration:

e2e:

postgres:
	docker start postgres14 ||\
	docker run --name postgres14 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=pw -d postgres:14-alpine

dbcreate:
	docker exec -it postgres14 createdb --username=root --owner=root db_test

dbdrop:
	docker exec -it postgres14 dropdb db_test

migrateup:
	migrate -path schema -database "postgresql://root:pw@localhost:5433/db_test?sslmode=disable" -verbose up
