package repository

import (
	"app/domain/model/user"
)

type UserRepository interface {
	PostUsers(user *user.User) (user_id *string, err error)
	PutUsers(user *user.User) error
	GetUserByEmail(email string) (user *user.User, err error)
	GetUser(user_id string) (user *user.User, err error)
}
