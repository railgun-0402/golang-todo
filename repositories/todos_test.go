// 「テストファイル xxx_test.go では [ディレクトリ名]_test というパッケージ名を使っても良い」という例外がある
package repositories_test

import (
	"strconv"
	"testing"
	"todo/models"
	"todo/repositories"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

// SelectTodos関数のテスト
func TestSelectTodos(t *testing.T) {

	expectedNum := 3
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
			expected: models.Todo{
				ID:    "1",
				Title: "firstTodo",
				Done:  false,
			},
		},
		{
			testTitle: "subtest2",
			expected: models.Todo{
				ID:    "2",
				Title: "secondTodo",
				Done:  true,
			},
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
			where id = ? title = ? and done = ?
		`
		testDB.Exec(sqlStr, test.ID, test.Title, test.Done)
	})
}