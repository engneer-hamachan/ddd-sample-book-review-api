package review

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Review struct {
	reviewId    reviewId
	userId      userId
	bookTitle   bookTitle
	reviewTitle reviewTitle
	publisher   publisher
	review      review
	readedAt    readedAt
	stars       stars
	publicFlg   publicFlg
}

type reviewId string
type userId string
type bookTitle string
type reviewTitle string
type publisher string
type review string
type readedAt time.Time
type stars int
type publicFlg bool

func New(reviewId string, userId string, bookTitle string, reviewTitle string, publisher string, reviewVal string, readedAt time.Time, stars int, publicFlg bool) (*Review, error) {

	createdReviewId, err := newReviewId(reviewId)
	if err != nil {
		return nil, err
	}

	createdUserId, err := newUserId(userId)
	if err != nil {
		return nil, err
	}

	createdBookTitle, err := newBookTitle(bookTitle)
	if err != nil {
		return nil, err
	}

	createdReviewTitle, err := newReviewTitle(reviewTitle)
	if err != nil {
		return nil, err
	}

	createdPublisher, err := newPublisher(publisher)
	if err != nil {
		return nil, err
	}

	createdReview, err := newReview(reviewVal)
	if err != nil {
		return nil, err
	}

	createdReadedAt, err := newReadedAt(readedAt)
	if err != nil {
		return nil, err
	}

	createdStars, err := newStars(stars)
	if err != nil {
		return nil, err
	}

	createdPublicFlg, err := newPublicFlg(publicFlg)
	if err != nil {
		return nil, err
	}

	review := Review{
		reviewId:    *createdReviewId,
		userId:      *createdUserId,
		bookTitle:   *createdBookTitle,
		reviewTitle: *createdReviewTitle,
		publisher:   *createdPublisher,
		review:      *createdReview,
		readedAt:    *createdReadedAt,
		stars:       *createdStars,
		publicFlg:   *createdPublicFlg,
	}
	return &review, nil
}

func Create(userId string, bookTitle string, reviewTitle string, publisher string, reviewVal string, readedAt time.Time, stars int, publicFlg bool) (*Review, error) {
	reviewId := uuid.New().String()
	review, err := New(reviewId, userId, bookTitle, reviewTitle, publisher, reviewVal, readedAt, stars, publicFlg)

	if err != nil {
		return nil, err
	}

	return review, err
}

func (r Review) GetReviewId() reviewId {
	return r.reviewId
}

func (r Review) GetUserId() userId {
	return r.userId
}

func (r Review) GetBookTitle() bookTitle {
	return r.bookTitle
}

func (r Review) GetReviewTitle() reviewTitle {
	return r.reviewTitle
}

func (r Review) GetPublisher() publisher {
	return r.publisher
}

func (r Review) GetReview() review {
	return r.review
}

func (r Review) GetReadedAt() readedAt {
	return r.readedAt
}

func (r Review) GetStars() stars {
	return r.stars
}

func (r Review) GetPublicFlg() publicFlg {
	return r.publicFlg
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

func newBookTitle(value string) (*bookTitle, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:bookTitle newBookTitle()")
		return nil, err
	}

	bookTitle := bookTitle(value)
	return &bookTitle, nil
}

func newReviewTitle(value string) (*reviewTitle, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:reviewTitle newReviewTitle()")
		return nil, err
	}

	reviewTitle := reviewTitle(value)
	return &reviewTitle, nil
}

func newPublisher(value string) (*publisher, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:publisher newPublisher()")
		return nil, err
	}

	publisher := publisher(value)
	return &publisher, nil
}

func newReview(value string) (*review, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:review newReview()")
		return nil, err
	}

	review := review(value)
	return &review, nil
}

func newReadedAt(value time.Time) (*readedAt, error) {
	readedAt := readedAt(value)
	return &readedAt, nil
}

func newStars(value int) (*stars, error) {

	if value < 1 || value > 5 {
		err := fmt.Errorf("%s", "error arg:stars between 1 and 5")
		return nil, err
	}

	stars := stars(value)
	return &stars, nil
}

func newPublicFlg(value bool) (*publicFlg, error) {

	publicFlg := publicFlg(value)
	return &publicFlg, nil
}

func (r *Review) ChangePublicFlg(current_user_id string) (*Review, error) {

	if current_user_id != string(r.userId) {
		err := fmt.Errorf("%s", "can not change public flg. because this Review is not your review.")
		return nil, err
	}

	if r.publicFlg == true {
		r.publicFlg = false
	} else {
		r.publicFlg = true
	}

	return r, nil
}
