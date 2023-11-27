package test

import (
	Init "fintechpractices/init"
	"fintechpractices/internal/dao"
	"fintechpractices/internal/model"
	"testing"
)

func TestUserCRUD(t *testing.T) {
	Init.Initialization()
	users := []model.UserInfo{
		{
			UniAccount: "test_admin1",
			UserName:   "test1",
			PassWord:   "123456",
		},

		{
			UniAccount: "test_admin2",
			UserName:   "test2",
			PassWord:   "123456",
		},
	}
	for _, user := range users {
		if err := dao.CreateUser(&user); err != nil {
			t.Error(err.Error())
		}
	}
	for _, user := range users {
		if u, err := dao.GetUserInfo(user.UniAccount); err != nil {
			if u.ID != user.ID || u.UserName != user.UserName || u.PassWord != user.PassWord {
				t.Errorf("%#v not equal to %#v", user, u)
			}
		}

		if ok, err := dao.IsAccountExisted("test_admin3"); err != nil {
			t.Error(err.Error())
		} else if ok {
			t.Error("IsAccountExisted error")
		}

		if ok, err := dao.IsAccountExisted("test_admin1"); err != nil {
			t.Error(err.Error())
		} else if !ok {
			t.Error("IsAccountExisted error")
		}
	}

	dao.UpdateUser("test_admin2", map[string]interface{}{"user_name": "testUpdate"})
	if ui, _ := dao.GetUserInfo("test_admin2"); ui.UserName != "testUpdate" {
		t.Errorf("expected %s, but got %s", "testUpdate", ui.UserName)
	}
}
