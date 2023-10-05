package domain_service

import (
	"app/domain/repository"
)

type ReviewDomainService interface {
	IsInsertReviewLike(review_id string, current_user_id string) (bool, error)
	IsInsertCommentLike(comment_id string, current_user_id string) (bool, error)
}

type reviewDomainService struct {
	reviewRepository repository.ReviewRepository
}

func NewReviewDomainService(rr repository.ReviewRepository) ReviewDomainService {
	return &reviewDomainService{
		reviewRepository: rr,
	}
}

func (rds *reviewDomainService) IsInsertReviewLike(review_id string, current_user_id string) (bool, error) {
	r, err := rds.reviewRepository.GetReviewLikeByReviewIdAndUserId(review_id, current_user_id)

	if err != nil {
		return false, err
	}

	if r != nil {
		return false, nil
	}

	return true, nil
}

func (rds *reviewDomainService) IsInsertCommentLike(comment_id string, current_user_id string) (bool, error) {
	c, err := rds.reviewRepository.GetCommentLikeByCommentIdAndUserId(comment_id, current_user_id)

	if err != nil {
		return false, err
	}

	if c != nil {
		return false, nil
	}

	return true, nil
}
