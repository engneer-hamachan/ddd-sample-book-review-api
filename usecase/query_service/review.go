package query_service

import (
	"app/usecase/query_service/dto"
)

type ReviewQueryService interface {
	GetReviews(current_page int) (*dto.Reviews, error)
	GetReviewWithComments(review_id string, current_page int) (*dto.ReviewWithComments, error)
}
