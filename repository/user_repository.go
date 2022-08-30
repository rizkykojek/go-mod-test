package repository

import "github.com/rizkykojek/go-mod-test/v2/entity"

type UserRepository interface {
	FindById(id int) *entity.User
}
