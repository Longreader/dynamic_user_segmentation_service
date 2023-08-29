# dynamic_user_segmentation_service

Test assignment for a Backend intern

## database running

- docker run --name=avito-db -e POSTGRES_PASSWORD='postgres' -p 5436:5432 -d --rm postgres

- migrate -path ./schema -database 'postgres://postgres:postgres@localhost:5436/postgres?sslmode=disable' up

- docker strop avito-db
