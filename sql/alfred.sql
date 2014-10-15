-- MySQL dump 10.13  Distrib 5.6.20, for osx10.9 (x86_64)
--
-- Host: 127.0.0.1    Database: alfred
-- ------------------------------------------------------
-- Server version	5.6.20

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `favor`
--

DROP TABLE IF EXISTS `favor`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `favor` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `tagset_id` int(11) DEFAULT NULL,
  `sys_id` int(11) DEFAULT NULL,
  `resource_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `favor`
--

LOCK TABLES `favor` WRITE;
/*!40000 ALTER TABLE `favor` DISABLE KEYS */;
/*!40000 ALTER TABLE `favor` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `group`
--

DROP TABLE IF EXISTS `group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `group` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `email` int(8) DEFAULT NULL,
  `comment` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  CONSTRAINT `group_ibfk_1` FOREIGN KEY (`id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group`
--

LOCK TABLES `group` WRITE;
/*!40000 ALTER TABLE `group` DISABLE KEYS */;
/*!40000 ALTER TABLE `group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `idalloc`
--

DROP TABLE IF EXISTS `idalloc`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `idalloc` (
  `id` bigint(64) unsigned NOT NULL AUTO_INCREMENT COMMENT '第三方接口分配 id',
  `api` varchar(11) NOT NULL DEFAULT '' COMMENT '接口名',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `idalloc`
--

LOCK TABLES `idalloc` WRITE;
/*!40000 ALTER TABLE `idalloc` DISABLE KEYS */;
INSERT INTO `idalloc` VALUES (1,'test'),(2,'test'),(3,'test'),(4,'test'),(5,'tree'),(6,'tree'),(7,'tree'),(8,'tree'),(9,''),(10,''),(11,'test'),(12,'test'),(13,'test'),(14,'test'),(15,'test'),(16,'test');
/*!40000 ALTER TABLE `idalloc` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `machine`
--

DROP TABLE IF EXISTS `machine`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `machine` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `ip` varchar(60) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `machine`
--

LOCK TABLES `machine` WRITE;
/*!40000 ALTER TABLE `machine` DISABLE KEYS */;
INSERT INTO `machine` VALUES (1,'web01.qq','127.0.0.1'),(4,'web04.qq','127.0.0.1');
/*!40000 ALTER TABLE `machine` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `monitor`
--

DROP TABLE IF EXISTS `monitor`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `monitor` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `metric` varchar(255) NOT NULL DEFAULT '',
  `type` int(8) NOT NULL DEFAULT '1' COMMENT '0: 机器监控, 1:业务监控',
  `comment` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`metric`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `monitor`
--

LOCK TABLES `monitor` WRITE;
/*!40000 ALTER TABLE `monitor` DISABLE KEYS */;
/*!40000 ALTER TABLE `monitor` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `permission`
--

DROP TABLE IF EXISTS `permission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `permission` (
  `taskset_id` bigint(20) NOT NULL,
  `group_id` bigint(20) NOT NULL COMMENT '用户组',
  `sys_id` bigint(20) NOT NULL COMMENT '第三方系统id',
  `perm` int(8) NOT NULL DEFAULT '1' COMMENT 'unix 权限位'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `permission`
--

LOCK TABLES `permission` WRITE;
/*!40000 ALTER TABLE `permission` DISABLE KEYS */;
/*!40000 ALTER TABLE `permission` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `subsys`
--

DROP TABLE IF EXISTS `subsys`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `subsys` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '英文名',
  `alias_name` varchar(255) DEFAULT NULL COMMENT '中文名',
  `hierarychy` varchar(512) NOT NULL DEFAULT '' COMMENT '显示层次结构',
  `author_id` int(11) NOT NULL COMMENT '负责人',
  `tree_interact` tinyint(1) DEFAULT NULL COMMENT '是否和 tree 交互',
  `comment` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `subsys`
--

LOCK TABLES `subsys` WRITE;
/*!40000 ALTER TABLE `subsys` DISABLE KEYS */;
/*!40000 ALTER TABLE `subsys` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tag_meta`
--

DROP TABLE IF EXISTS `tag_meta`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tag_meta` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `tag_key` varchar(255) NOT NULL,
  `alias_name` varchar(255) DEFAULT '' COMMENT '中文名',
  `allow_empty` tinyint(1) NOT NULL,
  `extra_info` varchar(128) DEFAULT '' COMMENT '用来标记额外信息存放表',
  `comment` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `tag_key` (`tag_key`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tag_meta`
--

LOCK TABLES `tag_meta` WRITE;
/*!40000 ALTER TABLE `tag_meta` DISABLE KEYS */;
INSERT INTO `tag_meta` VALUES (2,'corp','公司',0,'','该 tag 是用来标识公司');
/*!40000 ALTER TABLE `tag_meta` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tag_value`
--

DROP TABLE IF EXISTS `tag_value`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tag_value` (
  `id` bigint(20) unsigned NOT NULL,
  `key_id` bigint(20) unsigned NOT NULL,
  `value` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `value_index` (`value`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tag_value`
--

LOCK TABLES `tag_value` WRITE;
/*!40000 ALTER TABLE `tag_value` DISABLE KEYS */;
/*!40000 ALTER TABLE `tag_value` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tagging`
--

DROP TABLE IF EXISTS `tagging`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tagging` (
  `tagset_id` bigint(20) NOT NULL COMMENT '用来标记一个 tag 集合的 id',
  `sys_id` bigint(20) unsigned NOT NULL,
  `resource_id` bigint(20) NOT NULL,
  `tag_name` varchar(255) NOT NULL,
  `tag_value` varchar(255) NOT NULL,
  KEY `sys_id` (`sys_id`),
  KEY `tag_name` (`tag_name`),
  KEY `tag_value` (`tag_value`),
  CONSTRAINT `tagging_ibfk_1` FOREIGN KEY (`sys_id`) REFERENCES `subsys` (`id`) ON DELETE CASCADE,
  CONSTRAINT `tagging_ibfk_2` FOREIGN KEY (`tag_name`) REFERENCES `tag_meta` (`tag_key`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `tagging_ibfk_3` FOREIGN KEY (`tag_value`) REFERENCES `tag_value` (`value`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tagging`
--

LOCK TABLES `tagging` WRITE;
/*!40000 ALTER TABLE `tagging` DISABLE KEYS */;
/*!40000 ALTER TABLE `tagging` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `alias_name` varchar(255) NOT NULL DEFAULT '' COMMENT '中文名',
  `email` int(8) DEFAULT '1',
  `password` varchar(128) NOT NULL,
  `comment` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `view`
--

DROP TABLE IF EXISTS `view`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `view` (
  `uid` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户 id',
  `sys_id` int(11) NOT NULL,
  `hierachy` varbinary(512) NOT NULL DEFAULT '',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `view`
--

LOCK TABLES `view` WRITE;
/*!40000 ALTER TABLE `view` DISABLE KEYS */;
/*!40000 ALTER TABLE `view` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2014-10-13 17:05:24
