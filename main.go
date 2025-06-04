package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"todo/controllers"
	"todo/routers"
	"todo/services"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

var (
	dbUser = os.Getenv("MYSQL_USER")
	dbPassword = os.Getenv("MYSQL_PASSWORD")
	dbDatabase = os.Getenv("MYSQL_DATABASE")
	dbConn = fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}

	ser := services.NewMyAppService(db)
	con := controllers.NewTodoController(ser)

	// router層からハンドラの関係付けをする
	r := routers.NewRouter(con)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Content-Type"},
        AllowCredentials: true,
	}).Handler(r)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
