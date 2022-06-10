package repository

import (
	"app/domain/model/review"
	"app/domain/model/review/comment"
	"app/domain/model/review/review_like"
	"app/domain/model/review/comment_like"
)

type ReviewRepository interface {
	GetReviewByID(review_id string) (*review.Review, error)
	InsertReview(review *review.Review) (review_id *string, err error)
	PutReview(review *review.Review) error
	DeleteReview(review *review.Review) error
	InsertComment(comment *comment.Comment) (comment_id *string, err error)
	GetCommentByID(comment_id string) (*comment.Comment, error)
	DeleteComment(comment *comment.Comment) error
	GetReviewLikeByReviewIdAndUserId(review_id string, user_id string)(review_like *review_like.ReviewLike, err error)
	InsertReviewLike(review_like *review_like.ReviewLike) error
	DeleteReviewLike(review_like *review_like.ReviewLike) error

	GetCommentLikeByCommentIdAndUserId(comment_id string, user_id string)(comment_like *comment_like.CommentLike, err error)
	InsertCommentLike(comment_like *comment_like.CommentLike) error
	DeleteCommentLike(comment_like *comment_like.CommentLike) error

}
