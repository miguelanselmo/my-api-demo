DB_URL=file:./database.db?cache=shared&mode=rwc&doNotInterpretDatetime=1

network:
	docker network create bank-network

migrateup: 
	migrate -path infrastructure/repositories/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path infrastructure/repositores/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path infrastructure/repositores/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path infrastructure/repositores/migration -database "$(DB_URL)" -verbose down 1

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store


.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 db_docs db_schema sqlc test server mock proto evans

docker-build:
	docker build -t my_demo_api .

docker-run:
	docker run --name my_demo_api -p 8082:8082 my_demo_api:latest

docker-delete:
	docker rm --force my_demo_api
