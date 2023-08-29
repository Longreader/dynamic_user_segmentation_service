build:
	sudo docker-compose build avito-app

run:
	sudo docker-compose up avito-app

test:
	go test -v ./...

migrate:
	migrate -path ./schema -database 'postgres://postgres:postgres@localhost:5436/postgres?sslmode=disable' up