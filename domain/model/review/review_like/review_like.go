package review_like

import (
	"fmt"
	"github.com/google/uuid"
)

type ReviewLike struct {
	reviewLikeId reviewLikeId
  reviewId     reviewId
	userId       userId
}

type reviewLikeId string
type reviewId string
type userId string

func New(reviewLikeId string, reviewId string, userId string) (*ReviewLike, error) {

  createdReviewLikeId, err := newReviewLikeId(reviewLikeId)
	if err != nil {
		return nil, err
	}

  createdReviewId, err := newReviewId(reviewId)
	if err != nil {
		return nil, err
	}

	createdUserId, err := newUserId(userId)
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

func (r ReviewLike) GetReviewId() reviewId {
	return r.reviewId
}

func (r ReviewLike) GetUserId() userId {
	return r.userId
}

func newReviewLikeId(value string) (*reviewLikeId, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:reviewLikeId newReviewLikeId()")
		return nil, err
	}

	reviewLikeId := reviewLikeId(value)

	return &reviewLikeId, nil
}

func newReviewId(value string) (*reviewId, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:reviewId newReviewId()")
		return nil, err
	}

	reviewId := reviewId(value)

	return &reviewId, nil
}


func newUserId(value string) (*userId, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:userId newUserId()")
		return nil, err
	}

	userId := userId(value)
	return &userId, nil
}

func (r ReviewLike) IsReviewLikeYours(current_user_id string) (b bool, err error) {

	if string(r.userId) != current_user_id {
		err := fmt.Errorf("%s", "this review like is not yours")
		return false, err
	}

	return true, nil
}
