package model

import "time"

// CREATE TABLE metadata_market (
// 	id BIGINT NOT NULL,
// 	resource_id CHAR(32) NOT NULL COMMENT '资源ID',
// 	resource_describe VARCHAR(128) NOT NULL DEFAULT '' COMMENT '资源描述',
// 	resource_link VARCHAR(255) COMMENT '资源存储链接',
// 	resource_type INT COMMENT '资源类型',
// 	dp_link char(32) COMMENT '数字人link 唯一标识符',
// 	owner_id VARCHAR(255) COMMENT '所有者ID, 即用户账户',
// 	create_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
// 	update_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
// 	PRIMARY KEY (id),
// 	INDEX idx_source_id_source_link (resource_id, resource_link),
// 	INDEX idx_owner_id_source_type_source_id (owner_id, resource_type, resource_id)
//   ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储数字人元数据的数据表，即素材库';

type MetadataMarket struct {
	ID               int        `gorm:"primaryKey;autoIncrement"`
	ResourceId       string     `gorm:"column:resource_id"`
	ResourceDescribe string     `gorm:"column:resource_describe"`
	ResourceLink     string     `gorm:"column:resource_link"`
	ResourceType     string     `gorm:"column:resource_type"`
	DpLink           string     `gorm:"column:dp_link"`
	CoverImageLink   string     `gorm:"column:cover_image_link"`
	OwnerId          string     `gorm:"column:owner_id"`
	Deleted          bool       `gorm:"column:deleted"`
	IsReady          bool       `gorm:"column:is_ready"`
	CreateTime       *time.Time `gorm:"autoCreateTime; column:create_time"`
	UpdateTime       *time.Time `gorm:"autoUpdateTime; column:update_time"`
}

func (m *MetadataMarket) TableName() string {
	return "metadata_market"
}
