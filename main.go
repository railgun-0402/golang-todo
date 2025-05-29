package main

import (
	"database/sql"
	"fmt"
	"todo/models"

	_ "github.com/go-sql-driver/mysql"
)

func insertDB(db *sql.DB) {
	todo := models.Todo {
		ID: "3",
		Title: "Third Task",
		Done: false,
	}

	const sqlStr = `
		insert into todos (id, title, done, created_at) values(?, ?, ?, now());
	`

	result, err := db.Exec(sqlStr, todo.ID, todo.Title, todo.Done)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}

func update(db *sql.DB) {
	// トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	id := 2
	const sqlGetDone = `
		select done
		from todos
		where id = ?;
	`

	row := tx.QueryRow(sqlGetDone, id)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	var done bool
	err = row.Scan(&done)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	// タスクの完了を更新する
	const sqlUpdate = `update todos set done = ? where id = ?`
	_, err = tx.Exec(sqlUpdate, true, id)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	tx.Commit()
}

func main() {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	update(db)

	// for rows.Next() {
	// 	var todo models.Todo
	// 	var createdTime sql.NullTime
	// 	err := rows.Scan(&todo.ID, &todo.Title, &todo.Done, &createdTime)

	// 	if createdTime.Valid {
	// 		todo.CreatedAt = createdTime.Time
	// 	}

	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		todoArr = append(todoArr, todo)
	// 	}
	// }


	// r := mux.NewRouter()

	// r.HandleFunc("/get", handlers.GetTodos).Methods("GET")
	// r.HandleFunc("/get/{id:[0-9]+}", handlers.GetTodoById).Methods("GET")
	// r.HandleFunc("/create", handlers.CreateTodo).Methods("POST")
	// r.HandleFunc("/update/{id:[0-9]+}", handlers.Update).Methods("PUT")
	// r.HandleFunc("/delete/{id:[0-9]+}", handlers.Delete).Methods("DELETE")

	// log.Println("Server running on :8080")
	// log.Fatal(http.ListenAndServe(":8080", r))
}
