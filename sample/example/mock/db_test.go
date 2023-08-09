package mock_test

import (
	"errors"
	"testing"

	"example.com/mock"
	gomock "github.com/golang/mock/gomock"
)

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	db := mock.NewMockDB(ctrl)
	db.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist")) // 打桩(stubs)

	if v := mock.GetFromDB(db, "Tom"); v != -1 {
		t.Fatal("expected -1, but got", v)
	}
}
