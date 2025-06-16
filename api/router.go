package api

import (
	"database/sql"
	"todo/api/middlewares"
	"todo/controllers"
	"todo/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterRoutes(e *echo.Echo, db *sql.DB) {
	ser := services.NewTodoService(db)
	todoCon := controllers.NewTodoController(ser)

	// 共通のミドルウェア（EchoのLoggerミドルウェアでもOK）
	e.Use(middleware.Logger())
	e.Use(middlewares.LoggingMiddleware)

	// ルーティング設定
	e.GET("/get", todoCon.GetTodos)
	e.GET("/get/:id", todoCon.GetTodoByIdHandle)
	e.POST("/create", todoCon.CreateTodo)
	e.PUT("/update/:id", todoCon.Update)
	e.DELETE("/delete/:id", todoCon.Delete)

	// r := mux.NewRouter()

	// r.HandleFunc("/get", todoCon.GetTodos).Methods("GET")
	// r.HandleFunc("/get/{id:[0-9]+}", todoCon.GetTodoByIdHandle).Methods("GET")
	// r.HandleFunc("/create", todoCon.CreateTodo).Methods("POST")
	// r.HandleFunc("/update/{id:[0-9]+}", todoCon.Update).Methods("PUT")
	// r.HandleFunc("/delete/{id:[0-9]+}", todoCon.Delete).Methods("DELETE")

	// ルータ r に登録されているハンドラの前処理・後処理として
	// LoggingMiddleware が使われるようになる
	// r.Use(middlewares.LoggingMiddleware)

	// return r
}