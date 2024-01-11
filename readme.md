# Clone Project
```
git clone https://github.com/rizama/favorite-book-tracker.git
...

cd favorite-book-tracker
```

# Migration
## Install golang-migrate
```
go install -tags 'mysql,postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```
You can add other database to tags just add separated by comma, like this `-tags 'mysql,postgres,mongo'`. <br>
Make sure folder 'migrate' already exists in **$GOPATH/bin**
Make sure the command `migrate` are works in your cli. You can type `migrate -version`

## Create database
Create a new database.

## Run Migration
### Migration Up
Go in to project folder
```
migrate -path db/migrations -database "postgres://username:password@localhost:5432/your_database_name?sslmode=disable" up
```

### Migration Down
Go in to project folder
```
migrate -path db/migrations -database "postgres://username:password@localhost:5432/your_database_name?sslmode=disable" up
```