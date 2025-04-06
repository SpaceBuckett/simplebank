postgres: 
	docker run --name postgres17 -p 5432:5432 POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgress:17-alpine

dbversion: 
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" version

forcemigrate: 
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" force 0

closeallconnections: 
	docker exec -it postgres17 psql -U root -c "SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = 'simple_bank' AND pid <> pg_backend_pid();"

createdb: 
	docker exec -it postgres17 createdb --username=root --owner=root simple_bank 

dropdb: 
	docker exec -it postgres17 dropdb --username=root simple_bank  

migrateup: 
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 

migratedown: 
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 

sqlc: 
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown closeallconnections forcemigrate dbversion sqlc test