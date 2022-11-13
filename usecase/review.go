package usecase

import (
	"app/domain/domain_service"
	"app/domain/model/review"
	"app/domain/model/review/comment"
	"app/domain/model/review/comment_like"
	"app/domain/model/review/review_like"
	"app/domain/repository"
	"app/usecase/query_service"
	"app/usecase/query_service/dto"
	"fmt"
	"time"
)

type ReviewUseCase interface {
	ReviewCreate(user_id string, book_title string, review_title string, publisher string, reviewVal string, readed_at time.Time, stars int, public_flg bool) (review_id *string, err error)
	ReviewDetail(review_id string, current_page int) (*dto.ReviewWithComments, error)
	ReviewAll(current_page int) (*dto.Reviews, error)
	ChangeReviewPublicFlg(review_id string, current_user_id string) error
	ReviewDelete(review_id string, current_user_id string) error
	CommentCreate(review_id string, user_id string, comment string) (comment_id *string, err error)
	CommentDelete(comment_id string, current_user_id string) error
	ReviewLikeCreate(review_id string, current_user_id string) error
	ReviewLikeDelete(review_id string, current_user_id string) error
	CommentLikeCreate(comment_id string, current_user_id string) error
	CommentLikeDelete(comment_id string, current_user_id string) error
}

type reviewUseCase struct {
	reviewRepository    repository.ReviewRepository
	reviewDomainService domain_service.ReviewDomainService
	reviewQueryService  query_service.ReviewQueryService
}

func NewReviewUseCase(rr repository.ReviewRepository, rds domain_service.ReviewDomainService, rqs query_service.ReviewQueryService) ReviewUseCase {
	return &reviewUseCase{
		reviewRepository:    rr,
		reviewDomainService: rds,
		reviewQueryService:  rqs,
	}
}

func (ru *reviewUseCase) ReviewCreate(user_id string, book_title string, review_title string, publisher string, reviewVal string, readed_at time.Time, stars int, public_flg bool) (review_id *string, err error) {

	created_review, err := review.Create(user_id, book_title, review_title, publisher, reviewVal, readed_at, stars, public_flg)
	if err != nil {
		return nil, err
	}

	review_id, err = ru.reviewRepository.InsertReview(created_review)
	if err != nil {
		return nil, err
	}

	return review_id, nil
}

func (ru *reviewUseCase) ReviewDetail(review_id string, current_page int) (*dto.ReviewWithComments, error) {

	review, err := ru.reviewQueryService.GetReviewWithComments(review_id, current_page)
	if err != nil {
		return nil, err
	}

	return review, nil
}

func (ru *reviewUseCase) ReviewAll(current_page int) (*dto.Reviews, error) {

	reviews, err := ru.reviewQueryService.GetReviews(current_page)
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (ru *reviewUseCase) ChangeReviewPublicFlg(review_id string, current_user_id string) error {

	review, err := ru.reviewRepository.GetReviewByID(review_id)
	if err != nil {
		return err
	}

	if string(review.GetUserId()) != current_user_id {
		err := fmt.Errorf("%s", "this review is not your review")
		return err
	}

	review, err = review.ChangePublicFlg(current_user_id)
	if err != nil {
		return err
	}

	err = ru.reviewRepository.PutReview(review)
	if err != nil {
		return err
	}

	return nil
}

func (ru *reviewUseCase) ReviewDelete(review_id string, current_user_id string) error {

	review, err := ru.reviewRepository.GetReviewByID(review_id)
	if err != nil {
		return err
	}

	if string(review.GetUserId()) != current_user_id {
		err := fmt.Errorf("%s", "this review is not your review")
		return err
	}

	err = ru.reviewRepository.DeleteReview(review)
	if err != nil {
		return err
	}

	return nil
}

func (ru *reviewUseCase) CommentCreate(review_id string, user_id string, commentVal string) (comment_id *string, err error) {

	created_comment, err := comment.Create(review_id, user_id, commentVal)
	if err != nil {
		return nil, err
	}

	comment_id, err = ru.reviewRepository.InsertComment(created_comment)
	if err != nil {
		return nil, err
	}

	return comment_id, nil
}

func (ru *reviewUseCase) CommentDelete(comment_id string, current_user_id string) error {

	comment, err := ru.reviewRepository.GetCommentByID(comment_id)
	if err != nil {
		return err
	}

	if string(comment.GetUserId()) != current_user_id {
		err := fmt.Errorf("%s", "this comment is not your comment")
		return err
	}

	err = ru.reviewRepository.DeleteComment(comment)
	if err != nil {
		return err
	}

	return nil
}

func (ru *reviewUseCase) ReviewLikeCreate(review_id string, current_user_id string) error {

	b := ru.reviewDomainService.IsInsertReviewLike(review_id, current_user_id)
	if b == false {
		err := fmt.Errorf("%s", "this review is your liked")
		return err
	}

	created_review_like, err := review_like.Create(review_id, current_user_id)
	if err != nil {
		return err
	}

	err = ru.reviewRepository.InsertReviewLike(created_review_like)
	if err != nil {
		return err
	}

	return nil
}

func (ru *reviewUseCase) ReviewLikeDelete(review_id string, current_user_id string) error {

	review_like, err := ru.reviewRepository.GetReviewLikeByReviewIdAndUserId(review_id, current_user_id)
	if err != nil {
		return err
	}

	_, err = review_like.IsReviewLikeYours(current_user_id)
	if err != nil {
		return err
	}

	err = ru.reviewRepository.DeleteReviewLike(review_like)
	if err != nil {
		return err
	}

	return nil
}

func (ru reviewUseCase) CommentLikeCreate(comment_id string, current_user_id string) error {

	b := ru.reviewDomainService.IsInsertCommentLike(comment_id, current_user_id)
	if b == false {
		err := fmt.Errorf("%s", "this comment is your liked")
		return err
	}

	created_comment_like, err := comment_like.Create(comment_id, current_user_id)
	if err != nil {
		return err
	}

	err = ru.reviewRepository.InsertCommentLike(created_comment_like)
	if err != nil {
		return err
	}

	return nil
}

func (ru reviewUseCase) CommentLikeDelete(comment_id string, current_user_id string) error {

	comment_like, err := ru.reviewRepository.GetCommentLikeByCommentIdAndUserId(comment_id, current_user_id)
	if err != nil {
		return err
	}

	_, err = comment_like.IsCommentLikeYours(current_user_id)
	if err != nil {
		return err
	}

	err = ru.reviewRepository.DeleteCommentLike(comment_like)
	if err != nil {
		return err
	}

	return nil
}
