package dto

import (
	"app/domain/model/review/review_like"
)

type ReviewLike struct {
	ReviewLikeId string
	ReviewId     string
	UserId       string
}

func ConvertReviewLike(r *review_like.ReviewLike) *ReviewLike {
	return &ReviewLike{
		ReviewLikeId: string(r.GetReviewLikeId()),
		ReviewId:     string(r.GetReviewId()),
		UserId:       string(r.GetUserId()),
	}
}

func AdaptReviewLike(converted_review_like *ReviewLike) (*review_like.ReviewLike, error) {

	review_like, err := review_like.New(converted_review_like.ReviewLikeId, converted_review_like.ReviewId, converted_review_like.UserId)

	if err != nil {
		return nil, err
	}
	return review_like, nil
}
