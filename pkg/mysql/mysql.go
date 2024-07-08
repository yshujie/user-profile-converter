package mysql
import (
	"database/sql"
	"fmt"
	"os"

	_ "user-profile-converter/internal/repository/mysql"
)

/**
 * Connect to MySQL
 * @return *sql.DB
 */
func Connect() (*sql.DB, error) {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	passwd := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DB")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, passwd, host, port, dbname)
	db, err := sql.Open("mysql-logging", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}