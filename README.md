# Golang-Api-Crud

go test ./... -v -cover 
migrate -path migrations -database "postgres://postgres:1234@localhost:5432/postgres?sslmode=disable" force 1