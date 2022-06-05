package query_service

import (
	"app/usecase/query_service"
	"app/usecase/query_service/dto"
	"github.com/jinzhu/gorm"
)

type reviewQueryService struct {
	Conn *gorm.DB
}

func NewReviewQueryService(conn *gorm.DB) query_service.ReviewQueryService {
	return &reviewQueryService{Conn: conn}
}

func (rqs *reviewQueryService) GetReviews(current_page int) (*dto.Reviews, error) {

	PAGE_SIZE := 10

	offset := (current_page - 1) * PAGE_SIZE
	var result_reviews dto.Reviews

	if result := rqs.Conn.Table("reviews").
		Order("reviews.id desc").
		Offset(offset).
		Limit(PAGE_SIZE).
		Find(&result_reviews.Reviews); result.Error != nil {
			err := result.Error
			return nil, err
		}

	var ct int
	if result := rqs.Conn.Table("reviews").
		Count(&ct); result.Error != nil {
			err := result.Error
			return nil, err
		}

	var err error
	result_reviews.Page, err = dto.NewPage("/reviews", current_page, PAGE_SIZE, ct)
	if err != nil {
		return nil, err
	}
	return &result_reviews, nil
}

func (rqs *reviewQueryService) GetReviewWithComments(review_id string, current_page int) (*dto.ReviewWithComments, error) {

	PAGE_SIZE := 10

	offset := (current_page - 1) * PAGE_SIZE
	var result_review dto.ReviewWithComments

	if result := rqs.Conn.Table("reviews").
		Where("review_id = ?", review_id).
		First(&result_review); result.Error != nil {
  		err := result.Error
  		return nil, err
  	}

	if result := rqs.Conn.Table("comments").
		Select("comments.*, users.name as user_name").
		Joins("left join users on users.user_id = comments.user_id").
		Where("review_id = ?", review_id).
		Order("comments.id desc").
		Offset(offset).
		Limit(PAGE_SIZE).
		Scan(&result_review.Comments.Comments); result.Error != nil {
  		err := result.Error
  		return nil, err
  	}

	var ct int
	if result := rqs.Conn.Table("comments").
		Where("review_id = ?", review_id).
		Count(&ct); result.Error != nil {
			err := result.Error
			return nil, err
		}

	var err error
	result_review.Comments.Page, err = dto.NewPage("/comments", current_page, PAGE_SIZE, ct)
	if err != nil {
		return nil, err
	}
	return &result_review, nil
}
