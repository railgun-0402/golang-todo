// 「テストファイル xxx_test.go では [ディレクトリ名]_test というパッケージ名を使っても良い」という例外がある
package repositories_test

import (
	"database/sql"
	"fmt"
	"testing"
	"todo/models"
	"todo/repositories"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectDetailTodo(t *testing.T) {
	// Todo:DB共通関数を後で作る
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	expected := models.Todo {
		ID: "1",
		Title: "firstTodo",
		Done: false,
	}

	got, err := repositories.SelectDetailTodo(db, 1)
	if err != nil {
		// 失敗したらテスト終了
		t.Fatal(err)
	}

	if got.ID != expected.ID {
		t.Errorf("got %+v but want %+v\n", got, expected)
	}

	if got.Title != expected.Title {
		t.Errorf("got %+v but want %+v\n", got, expected)
	}

	if got.Done != expected.Done {
		t.Errorf("got %+v but want %+v\n", got, expected)
	}
}