build:
	sudo docker-compose build avito-app

run:
	docker-compose up -d avito-app

stop:
	docker-compose stop avito-app db

test:
	go test -v ./...

migrate:
	migrate -path ./schema -database 'postgres://postgres:postgres@localhost:5436/postgres?sslmode=disable' up