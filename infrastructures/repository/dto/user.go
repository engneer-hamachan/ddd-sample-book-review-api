package dto

import (
	"app/domain/model/user"
	"time"
)

type User struct {
	ID         int
	UserId     string
	Name       string
	Mail       string
	Password   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func ConvertUser(u *user.User) *User {
	return &User{
		UserId:     string(u.GetUserId()),
		Name:       string(u.GetName()),
		Mail:       string(u.GetMail()),
		Password:   string(u.GetPassword()),
	}
}

func AdaptUser(converted_user *User) (*user.User, error) {
	user, err := user.New(converted_user.UserId, converted_user.Name, converted_user.Mail, converted_user.Password)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func AdaptUsers(converted_users []*User) ([]*user.User, error) {
	var users []*user.User

	for _, converted_user := range converted_users {
		user, err := user.New(converted_user.UserId, converted_user.Name, converted_user.Mail, converted_user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
