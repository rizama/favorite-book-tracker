run migration
migrate -path db/migrations -database "postgres://sam:sam@localhost:5432/favorite_books?sslmode=disable" up