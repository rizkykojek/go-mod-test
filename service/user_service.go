package service

import (
	"errors"
	"github.com/rizkykojek/go-mod-test/v2/entity"
	"github.com/rizkykojek/go-mod-test/v2/repository"
)

type UserService struct {
	Repository repository.UserRepository
}

func (service UserService) Get(id int) (*entity.User, error) {
	user := service.Repository.FindById(id)
	if user == nil {
		return user, errors.New("user not found")
	} else {
		return user, nil
	}
}
