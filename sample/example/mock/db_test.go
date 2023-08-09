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

	db.EXPECT().Get(gomock.Not("Sam")).Return(0, nil).Times(2) // 打桩带调用次数
	mock.GetFromDB(db, "ABC")
	mock.GetFromDB(db, "DEF")

	o1 := db.EXPECT().Get(gomock.Any()).Return(630, nil)
	o2 := db.EXPECT().Get(gomock.Any()).Do(func(key string) {
		t.Log(key)
	})
	o3 := db.EXPECT().Get(gomock.Any()).DoAndReturn(func(key string) (int, error) {
		if key == "Sam" {
			return 630, nil
		}
		return 0, errors.New("not exist")
	})
	gomock.InOrder(o1, o2, o3) // 打桩带调用顺序
	mock.GetFromDB(db, "o1")
	mock.GetFromDB(db, "o2")
	mock.GetFromDB(db, "o3")
}
