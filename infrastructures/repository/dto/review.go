package dto

import (
	"app/domain/model/review"
	"time"
)

type Review struct {
	ID          int
	ReviewId    string
	UserId      string
	BookTitle   string
	ReviewTitle string
	Publisher   string
	Review      string
	ReadedAt    time.Time
	Stars       int
	PublicFlg   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func ConvertReview(r *review.Review) *Review {
	return &Review{
		ReviewId:    string(r.GetReviewId()),
		UserId:      string(r.GetUserId()),
		BookTitle:   string(r.GetBookTitle()),
		ReviewTitle: string(r.GetReviewTitle()),
		Publisher:   string(r.GetPublisher()),
		Review:      string(r.GetReview()),
		ReadedAt:    time.Time(r.GetReadedAt()),
		Stars:       int(r.GetStars()),
		PublicFlg:   bool(r.GetPublicFlg()),
	}
}

func AdaptReview(converted_review *Review) (*review.Review, error) {

	review, err := review.New(converted_review.ReviewId, converted_review.UserId, converted_review.BookTitle, converted_review.ReviewTitle, converted_review.Publisher, converted_review.Review, converted_review.ReadedAt, converted_review.Stars, converted_review.PublicFlg)

	if err != nil {
		return nil, err
	}
	return review, nil
}
