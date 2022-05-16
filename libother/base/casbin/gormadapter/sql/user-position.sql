#用户-岗位设计


###岗位表
CREATE TABLE `position`
(
    `id`          int(10)      NOT NULL AUTO_INCREMENT,
    `name`        varchar(50)  NOT NULL DEFAULT '' COMMENT '岗位名称',
    `creator`     varchar(32)  NOT NULL DEFAULT '' COMMENT '创建者',
    `updater`     varchar(32)  NOT NULL DEFAULT '' COMMENT '更新者',
    `created_at`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='岗位表';

####岗位~角色关联表
CREATE TABLE `position_role`
(
    `id`              int(10)     NOT NULL AUTO_INCREMENT,
    `position_id`     int(10)     NOT NULL DEFAULT '0' COMMENT '岗位id',
    `ps_id`           int(10)     NOT NULL DEFAULT '0' COMMENT '产品系统id',
    `actions`         text        NOT NULL COMMENT 'action集合',
    `creator`         varchar(32) NOT NULL DEFAULT '' COMMENT '创建者',
    `updater`         varchar(32) NOT NULL DEFAULT '' COMMENT '更新者',
    `created_at`      timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`      timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `data_updated_at` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `position_role` (`position_id`, `ps_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='管理后台岗位角色';
