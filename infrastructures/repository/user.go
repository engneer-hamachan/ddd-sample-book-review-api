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
