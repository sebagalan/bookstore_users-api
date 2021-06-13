package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/sebagalan/bookstore_users-api/logger"
)

const (
	bookstore_users_username string = "bookstore_users_username"
	bookstore_users_password string = "bookstore_users_password"
	bookstore_users_host     string = "bookstore_users_host"
	bookstore_users_database string = "bookstore_users_database"
)

var (
	UsersDb *sql.DB
)

/*
docker run  -p 3306:3306 --volume=/opt/storage/docker/mysql-data:/var/lib/mysql -td mysqltest2
*/
func init() {
	var err error

	err = godotenv.Load(".env")

	if err != nil {
		logger.Error("Error loading .env file", err)
	}

	username := os.Getenv(bookstore_users_username)
	password := os.Getenv(bookstore_users_password)
	host := os.Getenv(bookstore_users_host)
	database := os.Getenv(bookstore_users_database)

	datasourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, database)

	UsersDb, err = sql.Open("mysql", datasourceName)

	if err != nil {
		panic(err)
	}

	if err = UsersDb.Ping(); err != nil {
		panic(err)
	}

	log.Println("Connection successfully")

}
