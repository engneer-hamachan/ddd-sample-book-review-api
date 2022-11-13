package domain_service

type ReviewDomainService interface {
	IsInsertReviewLike(review_id string, current_user_id string) bool
	IsInsertCommentLike(comment_id string, current_user_id string) bool
}
