CREATE TABLE `student` (
                           `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                           `name` varchar(50) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '名称',
                           `age` int(11) NOT NULL DEFAULT '0' COMMENT '年龄',
                           `birthday` date NOT NULL DEFAULT '0000-00-00' COMMENT '生日',
                           `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                           `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '结束时间',
                           `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '结束时间',
                           `store_id` int(11) NOT NULL DEFAULT '0' COMMENT '校区ID',
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';


CREATE TABLE `store` (
                         `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                         `name` varchar(20) NOT NULL DEFAULT '',
                         `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;