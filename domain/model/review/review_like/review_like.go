package review_like

import (
	"app/domain/model/vo"
	"fmt"
	"github.com/google/uuid"
	"app/domain/model/review"
	"app/domain/model/user"
)

type ReviewLike struct {
	reviewLikeId vo.UuId
	reviewId     vo.UuId
	userId       vo.UuId
}

func New(reviewLikeId string, reviewId string, userId string) (*ReviewLike, error) {

	createdReviewLikeId, err := vo.NewUuId(reviewLikeId)
	if err != nil {
		return nil, err
	}

	createdReviewId, err := vo.NewUuId(reviewId)
	if err != nil {
		return nil, err
	}

	createdUserId, err := vo.NewUuId(userId)
	if err != nil {
		return nil, err
	}

	reviewLike := ReviewLike{
		reviewLikeId: *createdReviewLikeId,
		reviewId:     *createdReviewId,
		userId:       *createdUserId,
	}
	return &reviewLike, nil
}

func Create(reviewId string, userId string) (*ReviewLike, error) {
	reviewLikeId := uuid.New().String()
	review, err := New(reviewLikeId, reviewId, userId)

	if err != nil {
		return nil, err
	}

	return review, err
}

func (r *ReviewLike) GetReviewLikeId() vo.UuId {
	return r.reviewLikeId
}

func (r *ReviewLike) GetReviewId() vo.UuId {
	return r.reviewId
}

func (r *ReviewLike) GetUserId() vo.UuId {
	return r.userId
}

func (r *ReviewLike) IsReviewLikeYours(current_user_id string) (b bool, err error) {
	if string(r.userId) != current_user_id {
		err := fmt.Errorf("%s", "this review like is not yours")
		return false, err
	}

	return true, nil
}
