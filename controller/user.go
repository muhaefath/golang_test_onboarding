package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang_test/model"
	"log"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

type UserDetailController struct {
	beego.Controller
}

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "sirclo_new"
)

func (this *UserDetailController) Get() {

	// connect db local
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	// connect db local

	data := this.Ctx.Input.Query("id")

	sqlStatement := `SELECT id, name, password FROM users WHERE id = $1`

	myUser := model.User{}
	err = db.QueryRow(sqlStatement, data).Scan(&myUser.Id, &myUser.Name, &myUser.Password)
	if err != nil {
		panic(err)
	}

	this.Ctx.ResponseWriter.WriteHeader(206)
	this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(this.Ctx.ResponseWriter).Encode(myUser)
	fmt.Println(myUser)
}

func (this *UserController) Get() {

	this.Ctx.ResponseWriter.WriteHeader(201)

}

func (this *UserController) Post() {

	// connect db local
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	// connect db local

	user_model := model.User{}
	if err := this.ParseForm(&user_model); err != nil {
		//handle error
	}

	sqlStatement := `INSERT INTO users (id, name, password)
					 VALUES ($1,$2,$3)`

	_, err = db.Query(sqlStatement, user_model.Id, user_model.Name, user_model.Password)
	if err != nil {
		panic(err)
	}

	data := user_model
	this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(this.Ctx.ResponseWriter).Encode(data)

	this.Ctx.ResponseWriter.WriteHeader(202)
}

func (this *UserController) Put() {

	// connect db local
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	// connect db local

	user_model := model.User{}
	if err := this.ParseForm(&user_model); err != nil {
		//handle error
	}

	sqlStatement := `UPDATE users 
					 SET name = $2,password =$3
					 WHERE id = $1`

	_, err = db.Query(sqlStatement, user_model.Id, user_model.Name, user_model.Password)
	if err != nil {
		panic(err)
	}

	data := user_model
	this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(this.Ctx.ResponseWriter).Encode(data)

	this.Ctx.ResponseWriter.WriteHeader(202)
}
