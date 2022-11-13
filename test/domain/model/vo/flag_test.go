package vo_test

import (
	"app/domain/model/vo"
	"testing"
)

func TestNewFlag(t *testing.T) {
	_, err := vo.NewFlag(true)
	if err != nil {
		t.Errorf("Failed Make Instance")
	}
}

func TestFlagValue(t *testing.T) {
	got, _ := vo.NewFlag(true)
	if bool(*got) != true {
		t.Errorf("Can Not Value")
	}
}
