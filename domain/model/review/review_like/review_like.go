package review_like

import (
	"fmt"
	"github.com/google/uuid"
	"app/domain/model/review"
	"app/domain/model/user"
)

type ReviewLike struct {
	reviewLikeId reviewLikeId
	reviewId     review.ReviewId
	userId       user.UserId
}

func New(reviewLikeId string, reviewId string, userId string) (*ReviewLike, error) {

	createdReviewLikeId, err := newReviewLikeId(reviewLikeId)
	if err != nil {
		return nil, err
	}

	createdReviewId, err := review.NewReviewId(reviewId)
	if err != nil {
		return nil, err
	}

	createdUserId, err := user.NewUserId(userId)
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

func (r ReviewLike) GetReviewLikeId() reviewLikeId {
	return r.reviewLikeId
}

func (r ReviewLike) GetReviewId() review.ReviewId {
	return r.reviewId
}

func (r ReviewLike) GetUserId() user.UserId {
	return r.userId
}

func (r ReviewLike) IsReviewLikeYours(current_user_id string) (b bool, err error) {

	if string(r.userId) != current_user_id {
		err := fmt.Errorf("%s", "this review like is not yours")
		return false, err
	}

	return true, nil
}
