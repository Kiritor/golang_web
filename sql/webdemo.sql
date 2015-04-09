/*
Navicat MySQL Data Transfer

Source Server         : local
Source Server Version : 50051
Source Host           : localhost:3306
Source Database       : webdemo

Target Server Type    : MYSQL
Target Server Version : 50051
File Encoding         : 65001

Date: 2015-04-09 15:05:14
*/

SET FOREIGN_KEY_CHECKS=0;
-- ----------------------------
-- Table structure for webdemo_admin
-- ----------------------------
DROP TABLE IF EXISTS `webdemo_admin`;
CREATE TABLE `webdemo_admin` (
  `admin_id` int(11) NOT NULL auto_increment,
  `admin_name` varchar(32) NOT NULL,
  `admin_password` varchar(32) NOT NULL,
  PRIMARY KEY  (`admin_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of webdemo_admin
-- ----------------------------
INSERT INTO `webdemo_admin` VALUES ('1', 'LCore', 'LCore');
