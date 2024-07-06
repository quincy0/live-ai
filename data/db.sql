create database live;

DROP TABLE IF EXISTS `goods_info`;
CREATE TABLE `goods_info` (
                               `goods_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '表自增ID',
                               `name` varchar(128) NOT NULL DEFAULT '' COMMENT '商品名称',
                               `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
                               `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                               PRIMARY KEY (`goods_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品表';

DROP TABLE IF EXISTS `script_scene`;
CREATE TABLE `script_scene` (
                               `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '表自增ID',
                               `script_id` bigint unsigned NOT NULL COMMENT '剧本ID',
                               `scene_id` bigint unsigned NOT NULL COMMENT '分镜ID',
                               `goods_id` bigint unsigned NOT NULL COMMENT '商品ID',
                               `audio` varchar(255) NOT NULL COMMENT '音频地址',
                               `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
                               `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                               PRIMARY KEY (`id`),
                               UNIQUE KEY `uk_script_scene` (`script_id`,`scene_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='剧本信息';

DROP TABLE IF EXISTS `script_info`;
CREATE TABLE `script_info` (
                               `script_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '表自增ID',
                               `script_name` varchar(128) NOT NULL DEFAULT '' COMMENT '剧本名称',
                               `goods_id` bigint unsigned NOT NULL COMMENT '商品ID',
                               `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
                               `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                               PRIMARY KEY (`script_id`),
                               KEY `index_goods` (`goods_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='剧本信息';


DROP TABLE IF EXISTS `room_info`;
CREATE TABLE `room_info` (
                              `room_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '表自增ID',
                              `name` varchar(128) NOT NULL DEFAULT '' COMMENT '直播间名称',
                              `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
                              `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                              PRIMARY KEY (`room_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='直播间表';

DROP TABLE IF EXISTS `room_script`;
CREATE TABLE `room_script` (
                             `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '表自增ID',
                             `room_id` bigint unsigned NOT NULL COMMENT '直播间ID',
                             `script_id` bigint unsigned NOT NULL COMMENT '剧本ID',
                             `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
                             `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                             PRIMARY KEY (`id`),
                             UNIQUE KEY `uk_room_script` (`room_id`,`script_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='直播剧本表';