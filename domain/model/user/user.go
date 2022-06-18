package user

import (
	"github.com/google/uuid"
)

type User struct {
	userId     userId
	name       name
	mail       mail
	password   password
}

func New(userId string, name string, mail string, password string) (*User, error) {

	createdUserId, err := newUserId(userId)
	if err != nil {
		return nil, err
	}

	createdName, err := newName(name)
	if err != nil {
		return nil, err
	}

	createdMail, err := newMail(mail)
	if err != nil {
		return nil, err
	}

	createdPassword, err := newPassword(password)
	if err != nil {
		return nil, err
	}

	user := User{
		userId:     *createdUserId,
		name:       *createdName,
		mail:       *createdMail,
		password:   *createdPassword,
	}
	return &user, nil
}

//Create Constructor
func Create(name string, mail string, password string) (*User, error) {
	userId := uuid.New().String()
	user, err := New(userId, name, mail, password)

	if err != nil {
		return nil, err
	}

	return user, err
}

func (u User) GetUserId() userId {
	return u.userId
}

func (u User) GetName() name {
	return u.name
}

func (u User) GetMail() mail {
	return u.mail
}

func (u User) GetPassword() password {
	return u.password
}

func (u *User) ChangeMail(value string) (*User, error) {

	mail, err := newMail(value)
	if err != nil {
		return nil, err
	}
	u.mail = *mail

	return u, nil
}
