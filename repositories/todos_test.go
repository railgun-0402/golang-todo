// 「テストファイル xxx_test.go では [ディレクトリ名]_test というパッケージ名を使っても良い」という例外がある
package repositories_test

import (
	"database/sql"
	"strconv"
	"testing"
	"todo/models"
	"todo/repositories"
	"todo/repositories/testdata"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

// SelectTodos関数のテスト
func TestSelectTodos(t *testing.T) {

	expectedNum := len(testdata.SelectDetailTestData)
	got, err := repositories.SelectTodos(testDB)
	require.NoError(t, err)

	if num := len(got); num != expectedNum {
		t.Errorf("want %v but got %v todos\n", expectedNum, num)
	}
}

// SelectDetailTodo関数のテスト
func TestSelectDetailTodo(t *testing.T) {
	tests := []struct {
		testTitle string
		expected  models.Todo
	}{
		{
			testTitle: "subtest1",
			expected: testdata.SelectDetailTestData[0],
		},
		{
			testTitle: "subtest2",
			expected: testdata.SelectDetailTestData[1],
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			id, err := strconv.Atoi(test.expected.ID)
			require.NoError(t, err)

			got, err := repositories.SelectDetailTodo(testDB, id)
			require.NoError(t, err)

			if got.ID != test.expected.ID {
				t.Errorf("got %+v but want %+v\n", got, test.expected.ID)
			}

			if got.Title != test.expected.Title {
				t.Errorf("got %+v but want %+v\n", got, test.expected.Title)
			}

			if got.Done != test.expected.Done {
				t.Errorf("got %+v but want %+v\n", got, test.expected.Done)
			}
		})
	}
}

// InsertTodo関数のテスト
func TestInsertTodo(t *testing.T) {
	test := models.Todo {
		ID: "9999",
		Title: "insertTodo",
		Done: false,
	}

	expectedTodoNum := "9999"
	newTodo, err := repositories.InsertTodo(testDB, test)
	require.NoError(t, err)

	if newTodo.ID != expectedTodoNum {
		t.Errorf("new todo id is expected %+v but got %+v\n", expectedTodoNum, newTodo.ID)
	}

	// 後処理：追加したデータを必ず消す
	t.Cleanup(func()  {
		const sqlStr = `
			delete from todos
			where id = ?
		`
		testDB.Exec(sqlStr, test.ID)
	})
}

// UpdateTodo関数のテスト
func TestUpdateTodo(t *testing.T) {

	// テストデータを挿入
	testData := models.Todo {
		ID: "9999",
		Title: "updateTodo",
		Done: false,
	}

	_, err := testDB.Exec("INSERT INTO todos (id, title, done) VALUES(?, ?, ?)", testData.ID, testData.Title, testData.Done)
	require.NoError(t, err)

	// falseとtrueそれぞれに更新できることを確認する
	err = repositories.UpdateTodo(testDB, 9999, true)
	require.NoError(t, err)

	// 結果を確認
	var updatedDone bool
	err = testDB.QueryRow("SELECT done FROM todos WHERE id = ?", testData.ID).Scan(&updatedDone)
	require.NoError(t, err)
	require.Equal(t, true, updatedDone)

	// 後処理：追加したデータを必ず消す
	t.Cleanup(func()  {
		const sqlStr = `
			delete from todos
			where id = ?
		`
		testDB.Exec(sqlStr, testData.ID)
	})
}

// DeleteTodo関数のテスト
func TestDeleteTodo(t *testing.T) {
	// テストデータを挿入
	testData := models.Todo {
		ID: "9999",
		Title: "deleteTodo",
		Done: false,
	}

	_, err := testDB.Exec("INSERT INTO todos (id, title, done) VALUES(?, ?, ?)", testData.ID, testData.Title, testData.Done)
	require.NoError(t, err)

	// idに紐づいたタスクを削除できるかを確認する
	err = repositories.DeleteTodo(testDB, 9999)
	require.NoError(t, err)

	// 削除されたか確認：該当レコードがなくなっているか
	var id, title string
	var done bool
	err = testDB.QueryRow("SELECT id, title, done FROM todos WHERE id = ?", testData.ID).Scan(&id, &title, &done)

	require.Error(t, err)
	require.Equal(t, sql.ErrNoRows, err)
}