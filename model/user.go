package model

type User struct {
	Id       int    `form:"id"`
	Name     string `form:"name"`
	Password string `form:"password"`
}
