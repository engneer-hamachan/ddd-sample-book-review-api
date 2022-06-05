package user

import (
	"fmt"
	"github.com/google/uuid"
)

type User struct {
	userId     userId
	name       name
	mail       mail
	password   password
}

type userId string
type name string
type mail string
type password string

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

func newUserId(value string) (*userId, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:userId newUserId()")
		return nil, err
	}

	userId := userId(value)

	return &userId, nil
}

func newName(value string) (*name, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:name newName()")
		return nil, err
	}

	name := name(value)
	return &name, nil
}

func newMail(value string) (*mail, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:mail newMail()")
		return nil, err
	}

	mail := mail(value)
	return &mail, nil
}

func (u *User) ChangeMail(value string) (*User, error) {

	mail, err := newMail(value)
	if err != nil {
		return nil, err
	}
	u.mail = *mail

	return u, nil
}

func newPassword(value string) (*password, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:password newPassword()")
		return nil, err
	}

	password := password(value)
	return &password, nil
}
