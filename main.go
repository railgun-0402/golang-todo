package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"todo/api"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	dbUser = os.Getenv("MYSQL_USER")
	dbPassword = os.Getenv("MYSQL_PASSWORD")
	dbDatabase = os.Getenv("MYSQL_DATABASE")
	dbHost     = os.Getenv("MYSQL_HOST")
	dbConn = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbDatabase)
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}

	// Echoインスタンス作成
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig {
		AllowOrigins:   []string{"http://localhost:3000"},
        AllowMethods:   []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE, echo.OPTIONS},
        AllowHeaders:   []string{echo.HeaderContentType},
        AllowCredentials: true,
	}))

	// ルーティング登録
	api.RegisterRoutes(e, db)

	log.Println("Server running on :8080")
	// log.Fatal(http.ListenAndServe(":8080", handler))
	log.Fatal(e.Start(":8080"))
}
