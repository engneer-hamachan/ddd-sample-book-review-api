package query_service

import (
	"app/usecase/query_service"
	"app/usecase/query_service/dto"
	"github.com/jinzhu/gorm"
)

type userQueryService struct {
	Conn *gorm.DB
}

func NewUserQueryService(conn *gorm.DB) query_service.UserQueryService {
	return &userQueryService{Conn: conn}
}

func (uqs *userQueryService) GetUsers(current_page int) (*dto.Users, error) {

	PAGE_SIZE := 10

	offset := (current_page - 1) * PAGE_SIZE
	var result_users dto.Users
	if result := uqs.Conn.Offset(offset).
		Limit(PAGE_SIZE).
		Find(&result_users.Users); result.Error != nil {
		err := result.Error
		return nil, err
	}

	var ct int
	if result := uqs.Conn.Table("users").
		Count(&ct); result.Error != nil {
		err := result.Error
		return nil, err
	}
	var err error
	result_users.Page, err = dto.NewPage("/users", current_page, PAGE_SIZE, ct)
	if err != nil {
		return nil, err
	}

	return &result_users, nil
}

func (ur *userQueryService) GetUserByID(user_id string) (*dto.UserForAuth, error) {

	result_user := dto.UserForAuth{}

	if result := ur.Conn.Table("users").Where("user_id = ?", user_id).First(&result_user); result.Error != nil {
		err := result.Error
		return nil, err
	}

	return &result_user, nil

}
