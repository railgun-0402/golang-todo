package repositories

import (
	"database/sql"
	"fmt"
	"todo/models"
)

// タスク一覧を取得するSelect関数
func SelectTodos(db *sql.DB) ([]models.Todo, error) {
	const sqlStr = `
		select * from todos;
	`
	var todo models.Todo
	todoArr := make([]models.Todo, 0)

	// DB実行
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println(err)
		return todoArr, err
	}

	var createdTime sql.NullTime

	for rows.Next() {
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Done,  &createdTime);

		if createdTime.Valid {
			todo.CreatedAt = createdTime.Time
		}

		if err != nil {
			fmt.Println(err)
			return todoArr, err
		} else {
			todoArr = append(todoArr, todo)
		}
	}
	fmt.Printf("%+v\n", todoArr)
	return todoArr, nil
}

// IDに紐づくタスクを取得するSelect関数
func SelectDetailTodo(db *sql.DB, id int) (models.Todo, error) {
	const sqlStr = `
		select * from todos where id = ?;
	`
	var todo models.Todo

	// DB実行
	row := db.QueryRow(sqlStr, id)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		return todo, err
	}

	var createdTime sql.NullTime
	err := row.Scan(&todo.ID, &todo.Title, &todo.Done, &createdTime)
	if err != nil {
		fmt.Println(err)
		return todo, err
	}

	if createdTime.Valid {
			todo.CreatedAt = createdTime.Time
	}
	fmt.Printf("%+v\n", todo)
	return todo, nil
}

// タスクをDBに追加するinsert関数
func InsertTodo(db *sql.DB, todo models.Todo) (models.Todo, error) {
	const sqlStr = `
		insert into todos (id, title, done, created_at) values(?, ?, ?, now());
	`

	result, err := db.Exec(sqlStr, todo.ID, todo.Title, todo.Done)
	if err != nil {
		fmt.Println(err)
		return todo, err
	}

	// 追加された最後のIDと影響行数を出力
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())

	return todo, nil
}

// Todo: done以外も更新できるようにする
func UpdateTodo(db *sql.DB, id int, done bool) error {

	// タスクの完了を更新する
	const sqlUpdate = `update todos set done = ? where id = ?`
	_, err := db.Exec(sqlUpdate, done, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}