package repository

import (
	"app/domain/model/user"
	"app/domain/repository"
	"app/infrastructures/repository/dto"

	"github.com/jinzhu/gorm"
)

type userRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) repository.UserRepository {
	return &userRepository{Conn: conn}
}

func (ur *userRepository) PostUsers(user *user.User) (user_id *string, err error) {

	save_user := dto.ConvertUser(user)

	if result := ur.Conn.Save(&save_user); result.Error != nil {
		err := result.Error
		return nil, err
	}

	return &save_user.UserId, nil
}

func (ur *userRepository) PutUsers(user *user.User) error {

	save_user := dto.ConvertUser(user)

	if result := ur.Conn.Table("users").Where("user_id = ?", save_user.UserId).Updates(save_user); result.Error != nil {
		err := result.Error
		return err
	}

	return nil
}

func (ur *userRepository) GetUser(user_id string) (*user.User, error) {
	user := dto.User{}

	if result := ur.Conn.Table("users").Where("user_id = ?", user_id).First(&user); result.Error != nil {
		err := result.Error
		return nil, err
	}

	result_user, err := dto.AdaptUser(&user)
	if err != nil {
		return nil, err
	}

	return result_user, nil
}

func (ur *userRepository) GetUserByEmail(email string) (*user.User, error) {
	user := dto.User{}

	if result := ur.Conn.Table("users").Where("email = ?", email).First(&user); result.Error != nil {
		err := result.Error
		return nil, err
	}

	result_user, err := dto.AdaptUser(&user)
	if err != nil {
		return nil, err
	}

	return result_user, nil
}
