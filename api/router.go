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
	// 先順にすると、token検証失敗時にログが出なくなる
	e.Use(middlewares.AuthMiddleware)

	// ルーティング設定
	e.GET("/get", todoCon.GetTodos)
	e.GET("/get/:id", todoCon.GetTodoByIdHandle)
	e.POST("/create", todoCon.CreateTodo)
	e.PUT("/update/:id", todoCon.Update)
	e.DELETE("/delete/:id", todoCon.Delete)

	// Health Check
	e.GET("/health", todoCon.HealthCheck)
}