package vo_test

import (
	"app/domain/model/vo"
	"testing"
)

func TestNewUuId(t *testing.T) {
	_, err := vo.NewUuId("aaaaaaa")
	if err != nil {
		t.Errorf("Failed Make Instance")
	}
}

func TestUuIdValue(t *testing.T) {
	got, _ := vo.NewUuId("test")
	if string(*got) != "test" {
		t.Errorf("Can Not Value")
	}
}
