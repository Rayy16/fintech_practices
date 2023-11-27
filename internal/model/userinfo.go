package model

import "time"

// CREATE TABLE user_info (
// 	id BIGINT NOT NULL AUTO_INCREMENT,
// 	uni_account VARCHAR(255) NOT NULL COMMENT '用户账号，唯一约束',
// 	user_name VARCHAR(255) NOT NULL COMMENT '用户名称',
// 	password CHAR(32) NOT NULL COMMENT '用户密码，采用md5+盐方式存储',
// 	create_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
// 	update_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
// 	PRIMARY KEY (id),
// 	UNIQUE INDEX uni_account (uni_account)
//   ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储用户数据的数据表';

type UserInfo struct {
	ID         int        `gorm:"primaryKe; autoIncrement"`
	UniAccount string     `gorm:"uniqueIndex; column:uni_account"`
	UserName   string     `gorm:"column:user_name"`
	PassWord   string     `gorm:"column:password"`
	CreateTime *time.Time `gorm:"autoCreateTime; column:create_time"`
	UpdateTime *time.Time `gorm:"autoUpdateTime; column:update_time"`
}

func (u *UserInfo) TableName() string {
	return "user_info"
}
