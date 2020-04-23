/*
 Navicat MySQL Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 100138
 Source Host           : localhost:3306
 Source Schema         : doc

 Target Server Type    : MySQL
 Target Server Version : 100138
 File Encoding         : 65001

 Date: 23/04/2020 22:47:50
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
  `cid` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `is_hot` tinyint(4) UNSIGNED NOT NULL DEFAULT 0,
  `is_recommend` tinyint(4) UNSIGNED NOT NULL DEFAULT 0,
  `status` tinyint(4) UNSIGNED NOT NULL DEFAULT 1,
  `contents` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `add_time` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

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
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of doc_cate
-- ----------------------------
INSERT INTO `doc_cate` VALUES (8, 'php', 1, 0, 1587533816, 1587563937, 1, 120);
INSERT INTO `doc_cate` VALUES (9, 'golang', 1, 0, 1587533867, 1587562146, 1, 100);
INSERT INTO `doc_cate` VALUES (10, 'swoole', 1, 0, 1587539293, 1587539293, 1, 100);

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
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of doc_token
-- ----------------------------
INSERT INTO `doc_token` VALUES (7, '93b7dcc0b8e1159dac481e37e680fee6', 1, 1587738303, 1587564609, 1587651903);

SET FOREIGN_KEY_CHECKS = 1;
