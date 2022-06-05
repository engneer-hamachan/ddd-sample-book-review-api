package comment

import (
	"fmt"
	"github.com/google/uuid"
)

type Comment struct {
	commentId   commentId
  reviewId    reviewId
	userId      userId
	comment     comment
}

type commentId string
type reviewId string
type userId string
type comment string

func New(commentId string, reviewId string, userId string, commentVal string) (*Comment, error) {

  createdCommentId, err := newCommentId(commentId)
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

	createdComment, err := newComment(commentVal)
	if err != nil {
		return nil, err
	}

	comment := Comment{
		commentId:    *createdCommentId,
		reviewId:     *createdReviewId,
		userId:       *createdUserId,
		comment:      *createdComment,
	}
	return &comment, nil
}

func Create(reviewId string, userId string, commentVal string) (*Comment, error) {
	commentId := uuid.New().String()
	comment, err := New(commentId, reviewId, userId, commentVal)

	if err != nil {
		return nil, err
	}

	return comment, err
}

func (c Comment) GetCommentId() commentId {
	return c.commentId
}

func (c Comment) GetReviewId() reviewId {
	return c.reviewId
}

func (c Comment) GetUserId() userId {
	return c.userId
}

func (c Comment) GetComment() comment {
	return c.comment
}


func newCommentId(value string) (*commentId, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:commentId newCommentId()")
		return nil, err
	}

	commentId := commentId(value)

	return &commentId, nil
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


func newComment(value string) (*comment, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:comment newComment()")
		return nil, err
	}

	comment := comment(value)
	return &comment, nil
}
