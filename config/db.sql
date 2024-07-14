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

DROP TABLE IF EXISTS `script_tag`;
CREATE TABLE `script_tag` (
                                `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '表自增ID',
                                `script_tag` varchar(64) not null default '' comment '剧本标签',
                                `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
                                `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                PRIMARY KEY (`id`),
                                KEY (`script_tag`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='内容文案信息';

DROP TABLE IF EXISTS `content_info`;
CREATE TABLE `content_info` (
                               `content_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '表自增ID',
                               `content_name` varchar(128) NOT NULL DEFAULT '' COMMENT '剧本名称',
                               `product_tag` varchar(64) not null default '' comment '剧本品类',
                               `script_tag` varchar(64) not null default '' comment '剧本标签',
                               `summary` text not null default '' comment '剧本简介',
                               `content` text not null default '' comment '剧本内容',
                               `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
                               `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                               PRIMARY KEY (`content_id`),
                               KEY (`script_tag`),
                               key (`product_tag`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='内容文案';

DROP TABLE IF EXISTS `chapter`;
CREATE TABLE `chapter` (
                                `chapter_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '篇章ID',
                                `chapter_title` varchar(128) NOT NULL DEFAULT '' COMMENT '篇章名称',
                                `product_tag` varchar(128) not null default '' comment '剧本品类',
                                `script_tag` varchar(128) not null default '' comment '剧本标签',
                                `summary` varchar(512) not null default '' comment '篇章简介',
                                `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
                                `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                PRIMARY KEY (`chapter_id`),
                                KEY (`script_tag`),
                                key (`product_tag`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文案篇章';

DROP TABLE IF EXISTS `paragraph`;
CREATE TABLE `paragraph` (
                           `paragraph_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '段落ID',
                           `chapter_id` bigint unsigned NOT NULL COMMENT '篇章ID',
                           `paragraph_title` varchar(128) NOT NULL DEFAULT '' COMMENT '段落名称',
                           `content` text comment '段落内容',
                           `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
                           `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                           PRIMARY KEY (`paragraph_id`),
                           KEY (`chapter_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='段落文案';


DROP TABLE IF EXISTS `script`;
CREATE TABLE `script` (
                               `script_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '剧本ID',
                               `script_title` varchar(128) NOT NULL DEFAULT '' COMMENT '剧本名称',
                               `script_tag` varchar(128) not null default '' comment '剧本标签',
                               `product_tag` varchar(128) not null default '' comment '剧本品类',
                               `summary` varchar(512) not null default '' comment '剧本简介',
                               `timbre` varchar(128) not null default '' comment '音色',
                               `last_play` bigint unsigned NOT NULL COMMENT '上次播放时间',
                               `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
                               `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                               PRIMARY KEY (`script_id`),
                               KEY (`timbre`, `script_tag`),
                               key (`product_tag`,`timbre`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='剧本信息';

DROP TABLE IF EXISTS `scene`;
CREATE TABLE `scene` (
                               `scene_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '分镜ID',
                               `script_id` bigint unsigned NOT NULL COMMENT '剧本ID',
                               `scene_name` varchar(128) NOT NULL DEFAULT '' COMMENT '剧本名称',
                               `content` text comment '剧本内容',
                               `audio`  varchar(512) NOT NULL COMMENT '音频地址',
                               `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
                               `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                               PRIMARY KEY (`scene_id`),
                               KEY (`script_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='分镜信息';



DROP TABLE IF EXISTS `script_info`;
CREATE TABLE `script_info` (
                               `script_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '表自增ID',
                               `script_name` varchar(128) NOT NULL DEFAULT '' COMMENT '剧本名称',
                               `script_tag` varchar(64) not null default '' comment '剧本标签',
                               `product_tag` varchar(64) not null default '' comment '剧本品类',
                               `summary` text not null default '' comment '剧本简介',
                               `content` text not null default '' comment '剧本内容',
                               `timbre` varchar(64) not null default '' comment '音色',
                               `audio`  varchar(255) NOT NULL COMMENT '音频地址',
                               `last_play` bigint unsigned NOT NULL COMMENT '上次播放时间',
                               `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
                               `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                               PRIMARY KEY (`script_id`),
                               KEY (`script_tag`),
                               key (`product_tag`),
                               key (`timbre`)
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




DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info` (
                             `user_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
                             `username` varchar(64) NOT NULL COMMENT '用户名',
                             `password` char(32) NOT NULL COMMENT '密码',
                             `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
                             `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                             PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户信息';

DROP TABLE IF EXISTS `room_info`;
CREATE TABLE `room_info` (
                             `room_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '直播间ID',
                             `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
                             `room_name` varchar(128) NOT NULL DEFAULT '' COMMENT '直播间名称',
                             `product_tag` varchar(128) not null default '' comment '剧本品类',
                             `timbre` varchar(128) not null default '' comment '音色',
                             `template_id` int unsigned not null default '1' comment '模版id',
                             `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
                             `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                             PRIMARY KEY (`room_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='直播间信息';

DROP TABLE IF EXISTS `room_script`;
CREATE TABLE `room_script` (
                               `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '表自增ID',
                               `room_id` bigint unsigned NOT NULL COMMENT '直播间ID',
                               `script_id` bigint unsigned NOT NULL COMMENT '剧本ID',
                               `product_tag` varchar(128) not null default '' comment '剧本品类',
                               `timbre` varchar(128) not null default '' comment '音色',
                               `script_tag` varchar(128) not null default '' comment '剧本标签',
                               `sequence` int unsigned NOT NULL COMMENT '序号',
                               `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
                               `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                               PRIMARY KEY (`id`),
                               UNIQUE KEY `uk_room_script` (`room_id`,`script_id`, `sequence`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='直播间剧本';