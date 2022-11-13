package user

import (
	"app/domain/model/vo"
	"github.com/google/uuid"
)

type User struct {
	userId   vo.UuId
	name     vo.PersonName
	mail     vo.Email
	password vo.Password
}

func New(userId string, name string, mail string, pass string) (*User, error) {

	createdUserId, err := vo.NewUuId(userId)
	if err != nil {
		return nil, err
	}

	createdName, err := vo.NewPersonName(name)
	if err != nil {
		return nil, err
	}

	createdMail, err := vo.NewEmail(mail)
	if err != nil {
		return nil, err
	}

	createdPassword, err := vo.NewPassword(pass)
	if err != nil {
		return nil, err
	}

	user := User{
		userId:   *createdUserId,
		name:     *createdName,
		mail:     *createdMail,
		password: *createdPassword,
	}
	return &user, nil
}

// Create Constructor
func Create(name string, mail string, pass string) (*User, error) {
	userId := uuid.New().String()
	user, err := New(userId, name, mail, pass)

	if err != nil {
		return nil, err
	}

	return user, err
}

func (u *User) GetUserId() vo.UuId {
	return u.userId
}

func (u *User) GetName() vo.PersonName {
	return u.name
}

func (u *User) GetMail() vo.Email {
	return u.mail
}

func (u *User) GetPassword() vo.Password {
	return u.password
}

func (u *User) ChangeMail(value string) (*User, error) {

	mail, err := vo.NewEmail(value)
	if err != nil {
		return nil, err
	}
	u.mail = *mail

	return u, nil
}
