package benchmark

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := NewMockDB(ctrl)
	// 打桩
	db.EXPECT().Get(gomock.Eq("Tom")).Return(100, nil)
	// 使用db
	val, _ := db.Get("Tom")
	if val != 100 {
		t.Fatal("expect 100 get ", val)
	}
}
