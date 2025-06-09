package services_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	"todo/services"

	_ "github.com/go-sql-driver/mysql"
)

var tSer *services.TodoService

var (
	dbUser = os.Getenv("MYSQL_USER")
	dbPassword = os.Getenv("MYSQL_PASSWORD")
	dbDatabase = os.Getenv("MYSQL_DATABASE")
	dbConn = fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func TestMain(m *testing.M) {

	var err error
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	tSer = services.NewTodoService(db)

	m.Run()
}

func BenchmarkGetTodoService(b *testing.B) {
	ID := 1

	// 前処理時間を入れず、本メソッドの実行時間だけを対象にする
	b.ResetTimer()

	// for 文を b.N 回まわす
	for i := 0; i < b.N; i++ {
		_, err := tSer.GetTodoById(ID)
		if err != nil {
			b.Error(err)
			break
		}
	}
}