package test

import (
	"golang_test/model"
	"golang_test/repository"
	"golang_test/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = &repository.UserCategoryMock{Mock: mock.Mock{}}
var userService = service.UserService{Repository: userRepository}

func TestUserDetail(t *testing.T) {

	user := model.User{
		Id:       1,
		Name:     "hai",
		Password: "hai",
	}

	userRepository.Mock.On("FindById", "2").Return(user)

	result, err := userService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.Id, result.Id)
	assert.Equal(t, user.Name, result.Name)
	assert.Equal(t, user.Password, result.Password)

	println("user ", user.Id)
	println("user ", user.Name)

	println("result ", result.Id)
	println("result ", result.Name)
}

func TestUser_GetNotFound(t *testing.T) {

	userRepository.Mock.On("FindById", "1").Return(nil)
	category, err := userService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}
