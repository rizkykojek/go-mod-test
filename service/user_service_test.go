package service

import (
	"github.com/rizkykojek/go-mod-test/v2/entity"
	"github.com/rizkykojek/go-mod-test/v2/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var userRepository = &repository.UserRepositoryMock{Mock: mock.Mock{}}
var userService = UserService{Repository: userRepository}

func TestUserService_GetNotFound(t *testing.T) {
	userRepository.Mock.On("FindById", 1).Return(nil)
	user, err := userService.Get(1)
	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestUserService_GetSuccess(t *testing.T) {
	user := entity.User{
		Id:   2,
		Name: "kojek",
	}

	userRepository.Mock.On("FindById", 2).Return(user)
	result, err := userService.Get(2)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.Id, result.Id)
	assert.Equal(t, user.Name, result.Name)
}
