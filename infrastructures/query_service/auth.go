package query_service

import (
	"app/usecase/query_service"
	"app/usecase/query_service/dto"
	"github.com/jinzhu/gorm"
)

type authQueryService struct {
	Conn *gorm.DB
}

func NewAuthQueryService(conn *gorm.DB) query_service.AuthQueryService {
	return &authQueryService{Conn: conn}
}

func (aqs *authQueryService) GetUserByMailForAuth(mail string) (*dto.UserForAuth, error) {

	result_user := dto.UserForAuth{}

	if result := aqs.Conn.Table("users").Where("mail = ?", mail).First(&result_user); result.Error != nil {
		err := result.Error
		return nil, err
	}

	return &result_user, nil

}
