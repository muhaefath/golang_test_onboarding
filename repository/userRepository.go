package repository

import "golang_test/model"

type UserRepository interface {
	FindById(id string) *model.User
}
