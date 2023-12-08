package test

import (
	Init "fintechpractices/init"
	"fintechpractices/internal/dao"
	"fintechpractices/internal/model"

	"testing"
)

func TestResourceCRUD(t *testing.T) {
	Init.Initialization()
	resources := []model.MetadataMarket{
		{
			ResourceType:     dao.TypeImage.String(),
			ResourceLink:     "test_image_1.png",
			ResourceDescribe: "hello world",
			OwnerId:          "test_admin",
		},

		{
			ResourceType:     dao.TypeTone.String(),
			ResourceLink:     "test_tone_1.wav",
			ResourceDescribe: "hello world",
			OwnerId:          "test_admin",
		},

		{
			ResourceType:     dao.TypeImage.String(),
			ResourceLink:     "test_image_2.png",
			ResourceDescribe: "hello world",
			OwnerId:          "test_admin",
		},

		{
			ResourceType:     dao.TypeTone.String(),
			ResourceLink:     "test_tone_2.wav",
			ResourceDescribe: "hello world",
			OwnerId:          "test_admin",
		},

		{
			ResourceType:     dao.TypeImage.String(),
			ResourceLink:     "test_image_3.png",
			ResourceDescribe: "hello world",
			OwnerId:          "test_admin",
		},

		{
			ResourceType:     dao.TypeTone.String(),
			ResourceLink:     "test_tone_3.wav",
			ResourceDescribe: "hello world",
			OwnerId:          "test_admin",
		},
	}

	for _, resource := range resources {
		if err := dao.CreateMetadataMarket(&resource); err != nil {
			t.Error(err.Error())
		}
	}

	if rows, _, err := dao.GetResourceBy(dao.TypeBy(dao.TypeImage), dao.ResourceLinkBy("test_image_1.png")); err != nil {
		t.Error(err.Error())
	} else if len(rows) != 1 {
		t.Error("get resource by type and link failed")
	}

	dao.DeleteResourceByLink("test_image_2.png")
	if rows, _, err := dao.GetResourceBy(dao.ResourceLinkBy("test_image_2.png")); err != nil {
		t.Error(err.Error())
	} else if len(rows) != 0 {
		t.Error("delete resource by link failed")
	}

	if rows, _, err := dao.GetResourceBy(dao.TypeBy(dao.TypeTone), dao.PageBy(1, 2)); err != nil {
		t.Error(err.Error())
	} else if len(rows) != 2 {
		t.Errorf("filter error, expected 2 rows, but we got %d", len(rows))
	}

	if err := dao.UpdateResourceByLink("test_tone_1.wav", map[string]interface{}{"resource_describe": "updateTest"}); err != nil {
		t.Error(err.Error())
	} else {
		rows, _, _ := dao.GetResourceBy(dao.ResourceLinkBy("test_tone_1.wav"))
		if rows[0].ResourceDescribe != "updateTest" {
			t.Error("update resource failed")
		}
	}

	if err := dao.UpdateResourceByLink("test_tone_2.wav", map[string]interface{}{"resource_describe": "updateTest"}); err != nil {
		t.Error(err.Error())
	} else {
		rows, _, _ := dao.GetResourceBy(dao.ResourceLinkBy("test_tone_2.wav"))
		if rows[0].ResourceDescribe != "updateTest" {
			t.Error("update resource failed")
		}
	}
}
