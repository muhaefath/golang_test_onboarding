package service

import (
	"errors"
	"golang_test/model"
	"golang_test/repository"
)

type UserService struct {
	Repository repository.UserRepository
}

func (service UserService) Get(id string) (*model.User, error) {

	user := service.Repository.FindById(id)
	if user == nil {
		return nil, errors.New("category not found")
	} else {
		return user, nil
	}

}
