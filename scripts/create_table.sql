DROP Table digital_person_info;

DROP TABLE metadata_market;

CREATE TABLE user_info (
  id BIGINT NOT NULL AUTO_INCREMENT,
  uni_account VARCHAR(255) NOT NULL COMMENT '用户账号，唯一约束',
  user_name VARCHAR(255) NOT NULL COMMENT '用户名称',
  password CHAR(32) NOT NULL COMMENT '用户密码，采用md5+盐方式存储',
  create_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  update_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (id),
  UNIQUE INDEX uni_account (uni_account)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储用户数据的数据表';

CREATE TABLE digital_person_info (
  id BIGINT NOT NULL AUTO_INCREMENT,
  dp_id CHAR(64) NOT NULL DEFAULT '' COMMENT '数字人ID',
  dp_name VARCHAR(32) NOT NULL DEFAULT '' COMMENT '数字人名称',
  dp_status INT NOT NULL DEFAULT 0 COMMENT '状态, 枚举值(0-待生成；1-生成中；2-成功; 3-失败)',
  owner_id VARCHAR(255) NOT NULL COMMENT '所有者ID, 即用户账号',
  published BOOLEAN NOT NULL DEFAULT false COMMENT '发布状态',
  audited BOOLEAN NOT NULL DEFAULT false COMMENT '审核状态',
  dp_content VARCHAR(5120) NOT NULL DEFAULT '' COMMENT '内容',
  cover_image_link VARCHAR(255) NOT NULL COMMENT '封面图片存储链接',
  dp_link VARCHAR(255) NOT NULL DEFAULT '' COMMENT '数字人存储链接',
  deleted BOOLEAN NOT NULL DEFAULT false COMMENT '删除标志',
  hot_score INT NOT NULL DEFAULT 0 COMMENT '热度评分',
  create_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  update_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (id),
  INDEX idx_owner_id_dp_link_dp_status (owner_id, dp_link, dp_status),
  INDEX idx_dp_link_dp_name_cover_image_link (dp_link, dp_name, cover_image_link),
  INDEX idx_hot_score_create_time (hot_score, create_time)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储数字人信息的数据表';

CREATE TABLE metadata_market (
  id BIGINT NOT NULL AUTO_INCREMENT,
  resource_id CHAR(64) NOT NULL DEFAULT '' COMMENT '资源ID',
  resource_describe VARCHAR(128) NOT NULL DEFAULT '' COMMENT '资源描述',
  resource_link VARCHAR(255) NOT NULL DEFAULT '' COMMENT '资源存储链接',
  resource_type VARCHAR(16) DEFAULT 'unknown' COMMENT '资源类型',
  dp_link char(64) NOT NULL DEFAULT '' COMMENT '数字人link 唯一标识符',
  cover_image_link VARCHAR(255) NOT NULL DEFAULT '' COMMENT '封面图片存储链接',
  owner_id VARCHAR(255) COMMENT '所有者ID, 即用户账户',
  deleted BOOLEAN NOT NULL DEFAULT false COMMENT '删除标志',
  is_ready BOOLEAN NOT NULL DEFAULT false COMMENT '资源是否已准备',
  create_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  update_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (id),
  INDEX idx_owner_id_resource_type_resource_link (owner_id, resource_type, resource_link)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='存储数字人元数据的数据表，即素材库';

