package repository

import (
	"app/domain/model/review"
	"app/domain/model/review/comment"
	"app/domain/model/review/comment_like"
	"app/domain/model/review/review_like"
	"app/domain/repository"
	"app/infrastructures/repository/dto"
	"github.com/jinzhu/gorm"
)

type reviewRepository struct {
	Conn *gorm.DB
}

func NewReviewRepository(conn *gorm.DB) repository.ReviewRepository {
	return &reviewRepository{Conn: conn}
}

func (rr *reviewRepository) GetReviewByID(review_id string) (*review.Review, error) {

	var review dto.Review

	if result := rr.Conn.Table("reviews").
		Where("review_id = ?", review_id).
		First(&review); result.Error != nil {
		err := result.Error
		return nil, err
	}

	result_review, err := dto.AdaptReview(&review)
	if err != nil {
		return nil, err
	}

	return result_review, nil
}

func (rr *reviewRepository) InsertReview(review *review.Review) (review_id *string, err error) {

	converted_review := dto.ConvertReview(review)

	if result := rr.Conn.Save(&converted_review); result.Error != nil {
		err := result.Error
		return nil, err
	}

	return &converted_review.ReviewId, nil
}

func (rr *reviewRepository) PutReview(review *review.Review) error {

	converted_review := dto.ConvertReview(review)

	if result := rr.Conn.Table("review").
		Where("review_id = ?", string(review.GetReviewId())).
		Update("public_flg", bool(review.GetPublicFlg())).
		Updates(&converted_review); result.Error != nil {
		err := result.Error
		return err
	}

	return nil
}

func (rr *reviewRepository) DeleteReview(review *review.Review) error {

	converted_review := dto.ConvertReview(review)

	if result := rr.Conn.Where("review_id = ?", string(review.GetReviewId())).
		Delete(&converted_review); result.Error != nil {
		err := result.Error
		return err
	}

	return nil
}

func (rr *reviewRepository) InsertComment(comment *comment.Comment) (comment_id *string, err error) {
	converted_comment := dto.ConvertComment(comment)

	if result := rr.Conn.Save(&converted_comment); result.Error != nil {
		err := result.Error
		return nil, err
	}

	return &converted_comment.CommentId, nil
}

func (rr *reviewRepository) GetCommentByID(comment_id string) (*comment.Comment, error) {

	var comment dto.Comment

	if result := rr.Conn.Table("comments").
		Where("comment_id = ?", comment_id).
		First(&comment); result.Error != nil {
		err := result.Error
		return nil, err
	}

	result_comment, err := dto.AdaptComment(&comment)
	if err != nil {
		return nil, err
	}

	return result_comment, nil
}

func (rr *reviewRepository) DeleteComment(comment *comment.Comment) error {

	converted_comment := dto.ConvertComment(comment)

	if result := rr.Conn.Where("comment_id = ?", string(comment.GetCommentId())).
		Delete(&converted_comment); result.Error != nil {
		err := result.Error
		return err
	}

	return nil
}

func (rr *reviewRepository) GetReviewLikeByReviewIdAndUserId(review_id string, user_id string) (*review_like.ReviewLike, error) {

	var review_like dto.ReviewLike

	if result := rr.Conn.Table("review_likes").
		Where("review_id = ?", review_id).
		Where("user_id = ?", user_id).
		First(&review_like); result.Error != nil {
		err := result.Error
		return nil, err
	}

	result_review_like, err := dto.AdaptReviewLike(&review_like)
	if err != nil {
		return nil, err
	}

	return result_review_like, nil
}

func (rr *reviewRepository) InsertReviewLike(review_like *review_like.ReviewLike) error {

	converted_review_like := dto.ConvertReviewLike(review_like)

	if result := rr.Conn.Save(&converted_review_like); result.Error != nil {
		err := result.Error
		return err
	}

	return nil
}

func (rr *reviewRepository) DeleteReviewLike(review_like *review_like.ReviewLike) error {

	converted_review_like := dto.ConvertReviewLike(review_like)

	if result := rr.Conn.Where("review_like_id = ?", string(review_like.GetReviewLikeId())).
		Delete(&converted_review_like); result.Error != nil {
		err := result.Error
		return err
	}

	return nil
}

func (rr *reviewRepository) GetCommentLikeByCommentIdAndUserId(comment_id string, user_id string) (*comment_like.CommentLike, error) {

	var comment_like dto.CommentLike

	if result := rr.Conn.Table("comment_likes").
		Where("comment_id = ?", comment_id).
		Where("user_id = ?", user_id).
		First(&comment_like); result.Error != nil {
		err := result.Error
		return nil, err
	}

	result_comment_like, err := dto.AdaptCommentLike(&comment_like)
	if err != nil {
		return nil, err
	}

	return result_comment_like, nil
}

func (rr *reviewRepository) InsertCommentLike(comment_like *comment_like.CommentLike) error {

	converted_comment_like := dto.ConvertCommentLike(comment_like)

	if result := rr.Conn.Save(&converted_comment_like); result.Error != nil {
		err := result.Error
		return err
	}

	return nil
}

func (rr *reviewRepository) DeleteCommentLike(comment_like *comment_like.CommentLike) error {

	converted_comment_like := dto.ConvertCommentLike(comment_like)

	if result := rr.Conn.Where("comment_like_id = ?", string(comment_like.GetCommentLikeId())).
		Delete(&converted_comment_like); result.Error != nil {
		err := result.Error
		return err
	}

	return nil
}
