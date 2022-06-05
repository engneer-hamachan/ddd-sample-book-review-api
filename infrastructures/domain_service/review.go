package domain_service

import (
	"app/domain/domain_service"
	"github.com/jinzhu/gorm"
)

type reviewDomainService struct {
	Conn *gorm.DB
}

func NewReviewDomainService(conn *gorm.DB) domain_service.ReviewDomainService {
	return &reviewDomainService{Conn: conn}
}

func (rds *reviewDomainService) IsInsertReviewLike(review_id string, current_user_id string) bool {

	var ct int
	rds.Conn.Table("review_likes").Where("review_id = ?", review_id).Where("user_id = ?", current_user_id).Count(&ct)
	if ct == 0 {
		return true
	}
	return false
}
