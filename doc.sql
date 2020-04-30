/*
 Navicat MySQL Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 100137
 Source Host           : localhost:3306
 Source Schema         : doc

 Target Server Type    : MySQL
 Target Server Version : 100137
 File Encoding         : 65001

 Date: 30/04/2020 11:45:25
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for doc_admin
-- ----------------------------
DROP TABLE IF EXISTS `doc_admin`;
CREATE TABLE `doc_admin`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `last_login_ip` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `last_login_time` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `nickname` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `img` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '个人介绍',
  `desc` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '简介',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of doc_admin
-- ----------------------------
INSERT INTO `doc_admin` VALUES (1, 'admin', 'e10adc3949ba59abbe56e057f20f883e', '127.0.0.2', 0, '后台帅气小伙', 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif', NULL, '');

-- ----------------------------
-- Table structure for doc_article
-- ----------------------------
DROP TABLE IF EXISTS `doc_article`;
CREATE TABLE `doc_article`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `desc` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `look_times` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '查看次数',
  `stars` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '点赞次数',
  `cate_id` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `is_hot` tinyint(4) UNSIGNED NOT NULL DEFAULT 0,
  `is_recommend` tinyint(4) UNSIGNED NOT NULL DEFAULT 0,
  `status` tinyint(4) UNSIGNED NOT NULL DEFAULT 1,
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `add_time` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of doc_article
-- ----------------------------
INSERT INTO `doc_article` VALUES (17, 'vue中使用mavon-editor编辑器', 'vue中使用markendow编辑器的简介', 0, 0, 11, 1, 1, 1, '### 安装\n```npm\n$ npm install mavon-editor --save\n```\n\n### 引入和使用\n```vue\n<template>\n    <mavon-editor \n        v-model=\"content\"\n        placeholder=\"请输入文章内容\" style=\"height:500px\">\n</template>\n<script>\nimport mavonEditor from \'mavon-editor\'\nimport \'mavon-editor/dist/css/index.css\'\nexport default{\n    name: \"MavonEditor\",\n    data(){\n        content:\"\"\n    }\n    components: {\n        MavonEditor\n    }\n}\n</script>\n\n```\n[官方文档以及地址](https://github.com/hinesboy/mavonEditor/blob/master/README.md)\n\n', 1588217813, 1588217813);

-- ----------------------------
-- Table structure for doc_cate
-- ----------------------------
DROP TABLE IF EXISTS `doc_cate`;
CREATE TABLE `doc_cate`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `cate_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '分类名',
  `status` tinyint(4) UNSIGNED NOT NULL DEFAULT 1,
  `lock_times` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `add_time` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `is_index` tinyint(4) UNSIGNED NOT NULL DEFAULT 0,
  `sort` tinyint(4) UNSIGNED NOT NULL DEFAULT 100,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of doc_cate
-- ----------------------------
INSERT INTO `doc_cate` VALUES (8, 'php', 1, 0, 1587533816, 1587563937, 1, 120);
INSERT INTO `doc_cate` VALUES (9, 'golang', 1, 0, 1587533867, 1587562146, 1, 100);
INSERT INTO `doc_cate` VALUES (10, 'swoole', 1, 0, 1587539293, 1587539293, 1, 100);
INSERT INTO `doc_cate` VALUES (11, 'vue', 1, 0, 1588217275, 1588217275, 1, 100);

-- ----------------------------
-- Table structure for doc_token
-- ----------------------------
DROP TABLE IF EXISTS `doc_token`;
CREATE TABLE `doc_token`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `value` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `uid` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `expire_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '过期时间',
  `add_time` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of doc_token
-- ----------------------------
INSERT INTO `doc_token` VALUES (7, '93b7dcc0b8e1159dac481e37e680fee6', 1, 1587738303, 1587564609, 1587651903);
INSERT INTO `doc_token` VALUES (8, '1ce5eb7510acb967fb0d1aeea8e3335f', 1, 1587892928, 1587796612, 1587806528);
INSERT INTO `doc_token` VALUES (9, '7e7011ddf24f32bc15ae04cf960effa1', 1, 1588140136, 1588065128, 1588140136);
INSERT INTO `doc_token` VALUES (10, '5898b1cc121e90c299e03841b0b0f238', 1, 1588140366, 1588140362, 1588140366);
INSERT INTO `doc_token` VALUES (11, 'e8f80f90b4d0c2c1db449bf9e40945e8', 1, 1588140441, 1588140437, 1588140441);
INSERT INTO `doc_token` VALUES (12, 'f7b0662a281d737bc177ba8a01188610', 1, 1588140471, 1588140468, 1588140471);
INSERT INTO `doc_token` VALUES (13, '19cc10ef16ff1953da598023a56af350', 1, 1588302916, 1588140480, 1588216516);

SET FOREIGN_KEY_CHECKS = 1;
