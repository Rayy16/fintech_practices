package test

import (
	Init "fintechpractices/init"
	"fintechpractices/internal/dao"
	"fintechpractices/internal/model"
	"testing"
)

func TestDigitalPersonCRUD(t *testing.T) {
	Init.Initialization()
	dps := []model.DigitalPersonInfo{
		// {
		// 	DpId:           "test_dp_1.mp4",
		// 	DpName:         "test_dp_1",
		// 	DpStatus:       dao.StatusCreatable.Int(),
		// 	OwnerId:        "test_rliu",
		// 	Content:        "hello world",
		// 	CoverImageLink: "test_dp_1.png",
		// 	DpLink:         "test_dp_1.mp4",
		// },

		{
			DpId:           "test_dp_2.mp4",
			DpName:         "test_dp_2",
			DpStatus:       dao.StatusCreatable.Int(),
			OwnerId:        "test_admin",
			Content:        "hello world",
			CoverImageLink: "test_dp_2.png",
			DpLink:         "test_dp_2.mp4",
		},

		{
			DpId:           "test_dp_3.mp4",
			DpName:         "test_dp_3",
			DpStatus:       dao.StatusCreatable.Int(),
			OwnerId:        "test_admin",
			Content:        "hello world",
			CoverImageLink: "test_dp_3.png",
			DpLink:         "test_dp_3.mp4",
		},

		{
			DpId:           "test_dp_4.mp4",
			DpName:         "test_dp_4",
			DpStatus:       dao.StatusCreating.Int(),
			OwnerId:        "test_admin",
			Content:        "hello world",
			CoverImageLink: "test_dp_4.png",
			DpLink:         "test_dp_4.mp4",
		},

		{
			DpId:           "test_dp_5.mp4",
			DpName:         "test_dp_5",
			DpStatus:       dao.StatusSuccess.Int(),
			OwnerId:        "test_admin",
			Content:        "hello world",
			CoverImageLink: "test_dp_5.png",
			DpLink:         "test_dp_5.mp4",
		},

		{
			DpId:           "test_dp_6.mp4",
			DpName:         "test_dp_6",
			DpStatus:       dao.StatusFailed.Int(),
			OwnerId:        "test_admin",
			Content:        "hello world",
			CoverImageLink: "test_dp_6.png",
			DpLink:         "test_dp_6.mp4",
		},
	}

	for _, dp := range dps {
		if err := dao.CreateDigitalPerson(&dp); err != nil {
			t.Error(err.Error())
		}
	}

	if rows, _, err := dao.GetDigitalPersonsBy(dao.OwnerBy("test_rliu")); err != nil {
		t.Error(err.Error())
	} else if len(rows) != 1 {
		t.Error("get dp by owner failed, beacuse we can't get dp which owned by rliu")
	}

	dao.DeleteDigitalPersonByLink("test_dp_1.mp4")
	if rows, _, err := dao.GetDigitalPersonsBy(dao.DpLinkBy("test_dp_1.mp4")); err != nil {
		t.Error(err.Error())
	} else if len(rows) != 0 {
		t.Error("delete by link failed, because we get dp by link")
	}

	if rows, _, err := dao.GetDigitalPersonsBy(dao.StatusBy(dao.StatusCreatable), dao.OwnerBy("test_admin")); err != nil {
		t.Error(err.Error())
	} else if len(rows) != 2 {
		t.Errorf("filter error, expected 2 rows, but we got %d", len(rows))
	}

	if rows, _, err := dao.GetDigitalPersonsBy(dao.PageBy(1, 2)); err != nil {
		t.Error(err.Error())
	} else if len(rows) != 2 {
		t.Errorf("filter error, expected 2 rows, but we got %d", len(rows))
	}

	if err := dao.UpdateDigitalPersonByLink("test_dp_1.mp4", map[string]interface{}{"DpName": "testUpdate"}); err != nil {
		t.Error(err.Error())
	}

	if err := dao.UpdateDigitalPersonByLink("test_dp_2.mp4", map[string]interface{}{"DpName": "testUpdate"}); err != nil {
		t.Error(err.Error())
	} else {
		rows, _, _ := dao.GetDigitalPersonsBy(dao.DpLinkBy("test_dp_2.mp4"))
		if rows[0].DpName != "testUpdate" {
			t.Error("update dp failed")
		}
	}
}

func TestCount(t *testing.T) {
	Init.Initialization()

	if cnt, err := dao.Count((&model.DigitalPersonInfo{}).TableName(), dao.OwnerBy("test_admin")); err != nil {
		t.Error(err.Error())
	} else if cnt != 5 {
		t.Errorf("count or owner filter failed, because we expected 5 rows, but got %d", cnt)
	}
}

func TestOrder(t *testing.T) {
	Init.Initialization()
	rows1, _, err := dao.GetDigitalPersonsBy(dao.OwnerBy("test_admin"), dao.OrderBy("dp_name desc"))
	if err != nil {
		t.Error(err.Error())
	}
	rows2, _, err := dao.GetDigitalPersonsBy(dao.OwnerBy("test_admin"), dao.OrderBy("dp_name asc"))
	if err != nil {
		t.Error(err.Error())
	}
	n := len(rows1)
	for i := 0; i < n; i++ {
		if rows1[i].ID != rows2[n-1-i].ID {
			t.Errorf("order by dp_name failed: \ndesc: %v\n\n asc: %v", rows1[i], rows2[n-1-i])
		}
	}

	rows3, _, err := dao.GetDigitalPersonsBy(dao.OwnerBy("test_admin"), dao.OrderBy("dp_status, dp_name"))
	if err != nil {
		t.Error(err.Error())
	}
	expectedID := []int{
		15, 14, 16, 17, 18,
	}
	n = len(rows3)
	for i := 0; i < n; i++ {
		if rows3[i].ID != expectedID[i] {
			t.Errorf("order by dp_name failed: \ndesc: %v\n\n asc: %v", rows3[i], expectedID[i])
		}
	}
}
