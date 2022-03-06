package users_db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
	username = "root"
	password = ""
	host = "127.0.0.1:3306"
	schema = "users_db_02"
)

func init() {
	// username:password@tcp(host)/user_schema
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)
	sql.Open("mysql", dataSourceName)
}