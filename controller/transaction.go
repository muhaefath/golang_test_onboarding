package controller

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/astaxie/beego"
)

type TransactionController struct {
	beego.Controller
}

func (this *TransactionController) Get() {

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
					 VALUES (3,'halo','halo')`

	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	this.Ctx.ResponseWriter.WriteHeader(203)

}

func (this *TransactionController) Post() {

	this.Ctx.ResponseWriter.WriteHeader(204)

}
