#用户-角色设计
###用户表
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`         int(10)     NOT NULL AUTO_INCREMENT,
    `name`       varchar(50) NOT NULL DEFAULT '' COMMENT '用户名称',
    `role_ids`   varchar(50) NOT NULL DEFAULT '' COMMENT '角色,逗号分隔',
    `creator`    varchar(32) NOT NULL DEFAULT '' COMMENT '创建者',
    `updater`    varchar(32) NOT NULL DEFAULT '' COMMENT '更新者',
    `created_at` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';

INSERT INTO `user` (`id`, `name`, `role_ids`, `creator`, `updater`, `created_at`, `updated_at`)
VALUES (1, '许磊', '1,2', '超管', '超管', '2022-04-20 22:44:43', '2022-04-20 22:44:43'),
       (2, '张三', '2,3', '超管', '超管', '2022-04-20 22:45:10', '2022-04-20 22:45:10');


###菜单表，无限极分类
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu`
(
    `id`         int(10)      NOT NULL AUTO_INCREMENT,
    `name`       varchar(50)  NOT NULL DEFAULT '' COMMENT '菜单名',
    `pid`        int(50)      NOT NULL DEFAULT '0' COMMENT '父菜单的ID',
    `path`       varchar(100) NOT NULL DEFAULT '' COMMENT '前端路由',
    `component`  varchar(100) NOT NULL DEFAULT '' COMMENT '前端组件',
    `icon`       varchar(100) NOT NULL DEFAULT '' COMMENT '菜单图标',
    `is_show`    tinyint(4)   NOT NULL DEFAULT '1' COMMENT '是否显示:1-是;0-否',
    `creator`    varchar(32)  NOT NULL DEFAULT '' COMMENT '创建者',
    `updater`    varchar(32)  NOT NULL DEFAULT '' COMMENT '更新者',
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`),
    UNIQUE KEY `path` (`path`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='菜单表';

INSERT INTO `menu` (`id`, `name`, `pid`, `path`, `component`, `icon`, `is_show`, `creator`, `updater`, `created_at`,
                    `updated_at`)
VALUES (1, '学生名册', 0, '/student/index.html', '/student/index/index.vue', '', 1, '超管', '超管', '2022-04-20 22:45:50',
        '2022-04-20 22:45:50'),
       (2, '报名学生', 1, '/student/signup/index.html', '/student/signup/index.vue', '', 1, '超管', '超管',
        '2022-04-20 22:47:39', '2022-04-20 22:47:39'),
       (3, '体验学生', 1, '/student/experience/index.html', '/student/experience/index.vue', '', 1, '超管', '超管',
        '2022-04-20 22:48:46', '2022-04-20 22:48:46');


#接口定义表，无限极分类
DROP TABLE IF EXISTS `actions`;
CREATE TABLE `actions`
(
    `id`         int(10)      NOT NULL AUTO_INCREMENT,
    `name`       varchar(50)  NOT NULL DEFAULT '' COMMENT '资源名称',
    `pid`        int(50)      NOT NULL DEFAULT '0' COMMENT '父ID，0表示组',
    `path`       varchar(100) NOT NULL DEFAULT '' COMMENT '接口定义路由',
    `method`     varchar(10)  NOT NULL DEFAULT '' COMMENT '请求方法',
    `creator`    varchar(32)  NOT NULL DEFAULT '' COMMENT '创建者',
    `updater`    varchar(32)  NOT NULL DEFAULT '' COMMENT '更新者',
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`),
    UNIQUE KEY `path_method` (`path`, `method`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='接口定义表';

INSERT INTO `actions` (`id`, `name`, `pid`, `path`, `method`, `creator`, `updater`, `created_at`, `updated_at`)
VALUES (1, '学生列表', 0, '/student/search', 'GET', '超管', '超管', '2022-04-20 22:51:52', '2022-04-20 22:51:52'),
       (2, '报名学生列表', 1, '/student/signup/search', 'GET', '超管', '超管', '2022-04-20 22:51:05', '2022-04-20 22:51:05'),
       (3, '添加学生', 1, '/student/create', 'POST', '超管', '超管', '2022-04-20 22:52:48', '2022-04-20 22:52:48');



####角色表,支持无限极分类
###创建只需要填写基本信息即可
###权限可以通过单独功能做，减少复杂性
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`
(
    `id`         int(10)      NOT NULL AUTO_INCREMENT,
    `name`       varchar(50)  NOT NULL DEFAULT '' COMMENT '角色名称',
    `pid`        int(50)      NOT NULL DEFAULT '0' COMMENT '父角色ID，0表示根角色',
    `menu_ids`   varchar(500) NOT NULL DEFAULT '' COMMENT '菜单权限,逗号分隔',
    `action_ids` varchar(500) NOT NULL DEFAULT '' COMMENT '接口权限,逗号分隔',
    `creator`    varchar(32)  NOT NULL DEFAULT '' COMMENT '创建者',
    `updater`    varchar(32)  NOT NULL DEFAULT '' COMMENT '更新者',
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='角色表';

INSERT INTO `role` (`id`, `name`, `pid`, `menu_ids`, `action_ids`, `creator`, `updater`, `created_at`, `updated_at`)
VALUES (1, '校长', 0, '1,2,3', '1,2,3', '超管', '超管', '2022-04-20 22:49:52', '2022-04-20 22:49:52'),
       (2, '老师', 0, '1', '1', '超管', '超管', '2022-04-20 22:50:38', '2022-04-20 22:50:38');



###casbin_rule表，需要把user表的role_ids转换为该表的结构
###修改角色权限的时候，需要维护这张表，通过唯一索引查询，没有就插入，有就更新
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`
(
    `id`     int(11) unsigned NOT NULL AUTO_INCREMENT,
    `ptype` varchar(255)     NOT NULL DEFAULT '' COMMENT 'p',
    `v0`     varchar(100)     NOT NULL DEFAULT '' COMMENT '角色,通常存角色ID或者其他唯一标识符都可以',
    `v1`     varchar(100)     NOT NULL DEFAULT '' COMMENT '资源路径,path',
    `v2`     varchar(100)     NOT NULL DEFAULT '' COMMENT '操作方式,method',
    `v3`     varchar(100)     NOT NULL DEFAULT '' COMMENT '',
    `v4`     varchar(100)     NOT NULL DEFAULT '' COMMENT '',
    `v5`     varchar(100)     NOT NULL DEFAULT '' COMMENT '',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='casbin_rule表';

####casbin_conf可以放到项目的conf文件中
###sql2struct或者struct2sql