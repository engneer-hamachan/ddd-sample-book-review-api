package user

import (
	"github.com/google/uuid"
	"app/domain/model/vo"
)

type User struct {
	userId     vo.UserId
	name       vo.UserName
	mail       vo.Mail
	password   vo.Password
}

type password string

func New(userId string, name string, mail string, password string) (*User, error) {

	createdUserId, err := vo.NewUserId(userId)
	if err != nil {
		return nil, err
	}

	createdName, err := vo.NewUserName(name)
	if err != nil {
		return nil, err
	}

	createdMail, err := vo.NewMail(mail)
	if err != nil {
		return nil, err
	}

	createdPassword, err := vo.NewPassword(password)
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

func (u User) GetUserId() vo.UserId {
	return u.userId
}

func (u User) GetName() vo.UserName {
	return u.name
}

func (u User) GetMail() vo.Mail {
	return u.mail
}

func (u User) GetPassword() vo.Password {
	return u.password
}

func (u *User) ChangeMail(value string) (*User, error) {

	mail, err := vo.NewMail(value)
	if err != nil {
		return nil, err
	}
	u.mail = *mail

	return u, nil
}
