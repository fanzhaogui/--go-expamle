/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50709
Source Host           : localhost:3306
Source Database       : gz2017

Target Server Type    : MYSQL
Target Server Version : 50709
File Encoding         : 65001

Date: 2019-07-09 23:17:47
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `user_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `user_name` varchar(32) NOT NULL DEFAULT '' COMMENT '姓名',
  `tel` char(12) NOT NULL DEFAULT '' COMMENT '电话号码',
  `addr` varchar(225) NOT NULL DEFAULT '' COMMENT '地址',
  `update_time` int(10) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `add_time` int(10) NOT NULL DEFAULT '0' COMMENT '添加时间',
  `gender` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0未知 1男 2女',
  `age` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '年龄',
  `is_del` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态值',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES ('1', 'zhangsan1', '13455556661', '广州市海珠区 13号302房', '0', '0', '1', '10', '1');
INSERT INTO `users` VALUES ('2', 'zhangsan2', '13455556662', '广州市海珠区 13号302房', '0', '0', '2', '20', '1');
INSERT INTO `users` VALUES ('3', 'zhangsan3', '13455556663', '广州市海珠区 13号302房', '0', '0', '2', '30', '1');
INSERT INTO `users` VALUES ('4', 'zhangsan4', '13455556664', '广州市海珠区 13号302房', '0', '0', '1', '8', '1');
