package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var testDB *sql.DB

var (
	dbUser = os.Getenv("MYSQL_USER")
	dbPassword = os.Getenv("MYSQL_PASSWORD")
	dbDatabase = os.Getenv("MYSQL_DATABASE")
	dbConn = fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func connectDB() error {
	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func setup() error {
	if err := connectDB(); err != nil {
		return err
	}
	return nil
}

func teardown() {
	testDB.Close()
}

func TestMain(m *testing.M) {
	setup()

	m.Run()

	teardown()
}
