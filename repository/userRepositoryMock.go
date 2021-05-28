package repository

import (
	"fmt"
	"golang_test/model"

	"github.com/stretchr/testify/mock"
)

type UserCategoryMock struct {
	Mock mock.Mock
}

func (repository *UserCategoryMock) FindById(id string) *model.User {

	arguments := repository.Mock.Called(id)
	fmt.Println("arguments: ", arguments.Get(0))
	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(model.User)
		return &category
	}

}
