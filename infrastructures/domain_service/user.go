package domain_service

import (
	"app/domain/domain_service"
	"github.com/jinzhu/gorm"
)

type userDomainService struct {
	Conn *gorm.DB
}

func NewUserDomainService(conn *gorm.DB) domain_service.UserDomainService {
	return &userDomainService{Conn: conn}
}

func (uds *userDomainService) IsUserExists(mail string) bool {

	var ct int
	uds.Conn.Table("users").Where("mail = ?", mail).Count(&ct)
	if ct != 0 {
		return true
	}
	return false
}

func (uds *userDomainService) IsUserEnabled(user_id string) bool {

	var ct int
	uds.Conn.Table("users").Where("user_id = ?", user_id).Count(&ct)
	if ct != 0 {
		return true
	}
	return false
}

func (uds *userDomainService) IsCurrentUserMailDuplicated(user_id string, mail string) bool {

	var ct int
	uds.Conn.Table("users").Where("user_id != ?", user_id).Where("mail = ?", mail).Count(&ct)
	if ct != 0 {
		return true
	}
	return false
}
