package main

import (
	"database/sql"
	"fmt"
	"golang_test/controller"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq" // postgresql needed
)

func main() {

	fmt.Println("Halo")

	// CORS: allow all origins in any mode
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"https://*.local", "https://*.connexi.id"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	beego.Router("/user", &controller.UserController{})
	beego.Router("/user/detail/?:id", &controller.UserDetailController{})
	beego.Router("/transaction", &controller.TransactionController{})
	beego.BConfig.Listen.HTTPPort = 8082

	beego.Run()
}

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "sirclo_new"
)

func inits() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	if err2 := db.Ping(); err2 != nil {
		fmt.Println("Failed to keep connection alive")
	}

	sqlStatement := `INSERT INTO users (id, name, password)
					 VALUES (2,'halo','halo')`

	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}

	fmt.Println("masuk init")
}
