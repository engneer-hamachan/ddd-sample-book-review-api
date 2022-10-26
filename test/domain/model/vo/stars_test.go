package vo_test

import (
	"app/domain/model/vo"
	"testing"
)

func TestNewStars(t *testing.T) {
	_, err := vo.NewStars(1)
	if err != nil {
		t.Errorf("Failed Make Instance")
	}
}

func TestNewStarsError(t *testing.T) {
	_, err := vo.NewStars(10)
	if err == nil {
		t.Errorf("Failed Make Instance Error Test")
	}
}

func TestStarsValue(t *testing.T) {
	got, _ := vo.NewStars(5)
	if int(*got) != 5 {
		t.Errorf("Can Not Value")
	}
}
