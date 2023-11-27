package model

import "time"

// CREATE TABLE digital_person_info (
// 	id BIGINT NOT NULL,
// 	dp_id CHAR(32) NOT NULL COMMENT '数字人ID',
// 	dp_name VARCHAR(255) NOT NULL DEFAULT '' COMMENT '数字人名称',
// 	dg_status INT NOT NULL DEFAULT 0 COMMENT '状态, 枚举值(0-待生成；1-生成中；2-成功; 3-失败)',
// 	owner_id VARCHAR(255) NOT NULL COMMENT '所有者ID, 即用户账号',
// 	published BOOLEAN NOT NULL DEFAULT false COMMENT '发布状态',
// 	audited BOOLEAN NOT NULL DEFAULT true COMMENT '审核状态',
// 	content VARCHAR(65535) NOT NULL DEFAULT '' COMMENT '内容',
// 	cover_image_link VARCHAR(255) NOT NULL COMMENT '封面图片存储链接',
// 	dp_link VARCHAR(255) DEFAULT '' COMMENT '数字人存储链接',
// 	deleted BOOLEAN NOT NULL DEFAULT false COMMENT '删除标志',
// 	hot_score INT NOT NULL DEFAULT 0 COMMENT '热度评分',
// 	create_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
// 	update_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
// 	PRIMARY KEY (id),
// 	INDEX idx_dp_id_dp_status_dp_link (dp_id, dg_status, dp_link),
// 	INDEX idx_owner_id_dp_id_dp_name_cover_image_link (owner_id, dp_id, dp_name, cover_image_link),
// 	INDEX idx_hot_score_create_time (hot_score, create_time)
//   ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储数字人信息的数据表';

type DigitalPersonInfo struct {
	ID             int        `gorm:"primaryKey;autoIncrement"`
	DpId           string     `gorm:"column:dp_id"`
	DpName         string     `gorm:"column:dp_name"`
	DpStatus       int        `gorm:"column:dp_status"`
	OwnerId        string     `gorm:"column:owner_id"`
	Published      bool       `gorm:"column:published"`
	Audited        bool       `gorm:"column:audited"`
	Content        string     `gorm:"column:dp_content"`
	CoverImageLink string     `gorm:"column:cover_image_link"`
	DpLink         string     `gorm:"column:dp_link"`
	Deleted        bool       `gorm:"column:deleted"`
	HotScore       int        `gorm:"column:hot_score"`
	CreateTime     *time.Time `gorm:"autoCreateTime;column:create_time"`
	UpdateTime     *time.Time `gorm:"autoUpdateTime;column:update_time"`
}

func (dp *DigitalPersonInfo) TableName() string {
	return "digital_person_info"
}
