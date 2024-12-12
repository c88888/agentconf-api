-- MySQL dump 10.13  Distrib 8.0.35, for Linux (x86_64)
--
-- Host: localhost    Database: ga
-- ------------------------------------------------------
-- Server version	8.0.35

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
SET @MYSQLDUMP_TEMP_LOG_BIN = @@SESSION.SQL_LOG_BIN;
SET @@SESSION.SQL_LOG_BIN= 0;

--
-- GTID state at the beginning of the backup 
--

SET @@GLOBAL.GTID_PURGED=/*!80000 '+'*/ 'ee22a674-17ee-11ef-9d7c-fa163e5ef853:1-158268';

--
-- Table structure for table `casbin_rule`
--

DROP TABLE IF EXISTS `casbin_rule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=2745 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `casbin_rule`
--

LOCK TABLES `casbin_rule` WRITE;
/*!40000 ALTER TABLE `casbin_rule` DISABLE KEYS */;
INSERT INTO `casbin_rule` VALUES (1774,'p','1','/jwt/jsonInBlacklist','POST','','',''),(1777,'p','1','/menu/getMenu','POST','','',''),(1778,'p','1','/system/getServerInfo','POST','','',''),(1775,'p','1','/user/getUserInfo','GET','','',''),(1776,'p','1','/user/setUserAuthority','POST','','',''),(2642,'p','2','/authority/getAuthorityList','POST','','',''),(2636,'p','2','/jwt/jsonInBlacklist','POST','','',''),(2643,'p','2','/menu/getMenu','POST','','',''),(2646,'p','2','/sysOperationRecord/findSysOperationRecord','GET','','',''),(2647,'p','2','/sysOperationRecord/getSysOperationRecordList','GET','','',''),(2644,'p','2','/system/getServerInfo','POST','','',''),(2645,'p','2','/system/getSystemConfig','POST','','',''),(2637,'p','2','/user/admin_register','POST','','',''),(2640,'p','2','/user/getUserInfo','GET','','',''),(2638,'p','2','/user/getUserList','POST','','',''),(2639,'p','2','/user/setSelfInfo','PUT','','',''),(2641,'p','2','/user/setUserAuthority','POST','','',''),(2087,'p','201','/jwt/jsonInBlacklist','POST','','',''),(2090,'p','201','/menu/getMenu','POST','','',''),(2091,'p','201','/sysOperationRecord/findSysOperationRecord','GET','','',''),(2092,'p','201','/sysOperationRecord/getSysOperationRecordList','GET','','',''),(2088,'p','201','/user/getUserInfo','GET','','',''),(2089,'p','201','/user/setUserAuthority','POST','','',''),(2659,'p','888','/api/createApi','POST','','',''),(2660,'p','888','/api/deleteApi','POST','','',''),(2665,'p','888','/api/deleteApisByIds','DELETE','','',''),(2663,'p','888','/api/getAllApis','POST','','',''),(2664,'p','888','/api/getApiById','POST','','',''),(2662,'p','888','/api/getApiList','POST','','',''),(2661,'p','888','/api/updateApi','POST','','',''),(2666,'p','888','/authority/copyAuthority','POST','','',''),(2667,'p','888','/authority/createAuthority','POST','','',''),(2668,'p','888','/authority/deleteAuthority','POST','','',''),(2670,'p','888','/authority/getAuthorityList','POST','','',''),(2671,'p','888','/authority/setDataAuthority','POST','','',''),(2669,'p','888','/authority/updateAuthority','PUT','','',''),(2735,'p','888','/authorityBtn/canRemoveAuthorityBtn','POST','','',''),(2734,'p','888','/authorityBtn/getAuthorityBtn','POST','','',''),(2733,'p','888','/authorityBtn/setAuthorityBtn','POST','','',''),(2707,'p','888','/autoCode/createPackage','POST','','',''),(2704,'p','888','/autoCode/createPlug','POST','','',''),(2701,'p','888','/autoCode/createTemp','POST','','',''),(2709,'p','888','/autoCode/delPackage','POST','','',''),(2713,'p','888','/autoCode/delSysHistory','POST','','',''),(2703,'p','888','/autoCode/getColumn','GET','','',''),(2699,'p','888','/autoCode/getDB','GET','','',''),(2710,'p','888','/autoCode/getMeta','POST','','',''),(2708,'p','888','/autoCode/getPackage','POST','','',''),(2712,'p','888','/autoCode/getSysHistory','POST','','',''),(2700,'p','888','/autoCode/getTables','GET','','',''),(2705,'p','888','/autoCode/installPlugin','POST','','',''),(2702,'p','888','/autoCode/preview','POST','','',''),(2706,'p','888','/autoCode/pubPlug','POST','','',''),(2711,'p','888','/autoCode/rollback','POST','','',''),(2673,'p','888','/casbin/getPolicyPathByAuthorityId','POST','','',''),(2672,'p','888','/casbin/updateCasbin','POST','','',''),(2696,'p','888','/customer/customer','DELETE','','',''),(2697,'p','888','/customer/customer','GET','','',''),(2695,'p','888','/customer/customer','POST','','',''),(2694,'p','888','/customer/customer','PUT','','',''),(2698,'p','888','/customer/customerList','GET','','',''),(2732,'p','888','/email/emailTest','POST','','',''),(2684,'p','888','/fileUploadAndDownload/breakpointContinue','POST','','',''),(2685,'p','888','/fileUploadAndDownload/breakpointContinueFinish','POST','','',''),(2688,'p','888','/fileUploadAndDownload/deleteFile','POST','','',''),(2689,'p','888','/fileUploadAndDownload/editFileName','POST','','',''),(2683,'p','888','/fileUploadAndDownload/findFile','GET','','',''),(2690,'p','888','/fileUploadAndDownload/getFileList','POST','','',''),(2686,'p','888','/fileUploadAndDownload/removeChunk','POST','','',''),(2687,'p','888','/fileUploadAndDownload/upload','POST','','',''),(2648,'p','888','/jwt/jsonInBlacklist','POST','','',''),(2674,'p','888','/menu/addBaseMenu','POST','','',''),(2682,'p','888','/menu/addMenuAuthority','POST','','',''),(2676,'p','888','/menu/deleteBaseMenu','POST','','',''),(2678,'p','888','/menu/getBaseMenuById','POST','','',''),(2680,'p','888','/menu/getBaseMenuTree','POST','','',''),(2675,'p','888','/menu/getMenu','POST','','',''),(2681,'p','888','/menu/getMenuAuthority','POST','','',''),(2679,'p','888','/menu/getMenuList','POST','','',''),(2677,'p','888','/menu/updateBaseMenu','POST','','',''),(2730,'p','888','/simpleUploader/checkFileMd5','GET','','',''),(2731,'p','888','/simpleUploader/mergeFileMd5','GET','','',''),(2729,'p','888','/simpleUploader/upload','POST','','',''),(2719,'p','888','/sysDictionary/createSysDictionary','POST','','',''),(2720,'p','888','/sysDictionary/deleteSysDictionary','DELETE','','',''),(2722,'p','888','/sysDictionary/findSysDictionary','GET','','',''),(2723,'p','888','/sysDictionary/getSysDictionaryList','GET','','',''),(2721,'p','888','/sysDictionary/updateSysDictionary','PUT','','',''),(2715,'p','888','/sysDictionaryDetail/createSysDictionaryDetail','POST','','',''),(2716,'p','888','/sysDictionaryDetail/deleteSysDictionaryDetail','DELETE','','',''),(2717,'p','888','/sysDictionaryDetail/findSysDictionaryDetail','GET','','',''),(2718,'p','888','/sysDictionaryDetail/getSysDictionaryDetailList','GET','','',''),(2714,'p','888','/sysDictionaryDetail/updateSysDictionaryDetail','PUT','','',''),(2736,'p','888','/sysExportTemplate/createSysExportTemplate','POST','','',''),(2737,'p','888','/sysExportTemplate/deleteSysExportTemplate','DELETE','','',''),(2738,'p','888','/sysExportTemplate/deleteSysExportTemplateByIds','DELETE','','',''),(2742,'p','888','/sysExportTemplate/exportExcel','GET','','',''),(2743,'p','888','/sysExportTemplate/exportTemplate','GET','','',''),(2740,'p','888','/sysExportTemplate/findSysExportTemplate','GET','','',''),(2741,'p','888','/sysExportTemplate/getSysExportTemplateList','GET','','',''),(2744,'p','888','/sysExportTemplate/importExcel','POST','','',''),(2739,'p','888','/sysExportTemplate/updateSysExportTemplate','PUT','','',''),(2724,'p','888','/sysOperationRecord/createSysOperationRecord','POST','','',''),(2727,'p','888','/sysOperationRecord/deleteSysOperationRecord','DELETE','','',''),(2728,'p','888','/sysOperationRecord/deleteSysOperationRecordByIds','DELETE','','',''),(2725,'p','888','/sysOperationRecord/findSysOperationRecord','GET','','',''),(2726,'p','888','/sysOperationRecord/getSysOperationRecordList','GET','','',''),(2691,'p','888','/system/getServerInfo','POST','','',''),(2692,'p','888','/system/getSystemConfig','POST','','',''),(2693,'p','888','/system/setSystemConfig','POST','','',''),(2650,'p','888','/user/admin_register','POST','','',''),(2656,'p','888','/user/changePassword','POST','','',''),(2649,'p','888','/user/deleteUser','DELETE','','',''),(2654,'p','888','/user/getUserInfo','GET','','',''),(2651,'p','888','/user/getUserList','POST','','',''),(2658,'p','888','/user/resetPassword','POST','','',''),(2653,'p','888','/user/setSelfInfo','PUT','','',''),(2655,'p','888','/user/setUserAuthorities','POST','','',''),(2657,'p','888','/user/setUserAuthority','POST','','',''),(2652,'p','888','/user/setUserInfo','PUT','','','');
/*!40000 ALTER TABLE `casbin_rule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exa_customers`
--

DROP TABLE IF EXISTS `exa_customers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exa_customers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `customer_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '客户名',
  `customer_phone_data` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '客户手机号',
  `sys_user_id` bigint unsigned DEFAULT NULL COMMENT '管理ID',
  `sys_user_authority_id` bigint unsigned DEFAULT NULL COMMENT '管理角色ID',
  PRIMARY KEY (`id`),
  KEY `idx_exa_customers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exa_customers`
--

LOCK TABLES `exa_customers` WRITE;
/*!40000 ALTER TABLE `exa_customers` DISABLE KEYS */;
/*!40000 ALTER TABLE `exa_customers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exa_file_chunks`
--

DROP TABLE IF EXISTS `exa_file_chunks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exa_file_chunks` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `exa_file_id` bigint unsigned DEFAULT NULL,
  `file_chunk_number` bigint DEFAULT NULL,
  `file_chunk_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_exa_file_chunks_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exa_file_chunks`
--

LOCK TABLES `exa_file_chunks` WRITE;
/*!40000 ALTER TABLE `exa_file_chunks` DISABLE KEYS */;
/*!40000 ALTER TABLE `exa_file_chunks` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exa_file_upload_and_downloads`
--

DROP TABLE IF EXISTS `exa_file_upload_and_downloads`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exa_file_upload_and_downloads` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件名',
  `url` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件地址',
  `tag` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件标签',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '编号',
  PRIMARY KEY (`id`),
  KEY `idx_exa_file_upload_and_downloads_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exa_file_upload_and_downloads`
--

LOCK TABLES `exa_file_upload_and_downloads` WRITE;
/*!40000 ALTER TABLE `exa_file_upload_and_downloads` DISABLE KEYS */;
INSERT INTO `exa_file_upload_and_downloads` VALUES (1,'2024-04-17 18:04:16.364','2024-04-17 18:04:16.364',NULL,'10.png','https://fat.abc.com/asset/person.png','png','158787308910.png'),(2,'2024-04-17 18:04:16.364','2024-04-17 18:04:16.364',NULL,'logo.png','https://fat.abc.com/asset/person.png','png','1587973709logo.png');
/*!40000 ALTER TABLE `exa_file_upload_and_downloads` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exa_files`
--

DROP TABLE IF EXISTS `exa_files`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exa_files` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `file_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `file_md5` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `file_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `chunk_total` bigint DEFAULT NULL,
  `is_finish` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_exa_files_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exa_files`
--

LOCK TABLES `exa_files` WRITE;
/*!40000 ALTER TABLE `exa_files` DISABLE KEYS */;
/*!40000 ALTER TABLE `exa_files` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `jwt_blacklists`
--

DROP TABLE IF EXISTS `jwt_blacklists`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `jwt_blacklists` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `jwt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT 'jwt',
  PRIMARY KEY (`id`),
  KEY `idx_jwt_blacklists_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=228 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `jwt_blacklists`
--

LOCK TABLES `jwt_blacklists` WRITE;
/*!40000 ALTER TABLE `jwt_blacklists` DISABLE KEYS */;
/*!40000 ALTER TABLE `jwt_blacklists` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_apis`
--

DROP TABLE IF EXISTS `sys_apis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_apis` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`),
  KEY `idx_sys_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=139 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_apis`
--

LOCK TABLES `sys_apis` WRITE;
/*!40000 ALTER TABLE `sys_apis` DISABLE KEYS */;
INSERT INTO `sys_apis` VALUES (1,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/jwt/jsonInBlacklist','jwt加入黑名单(退出，必选)','jwt','POST'),(2,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/user/deleteUser','删除用户','系统用户','DELETE'),(3,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/user/admin_register','用户注册','系统用户','POST'),(4,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/user/getUserList','获取用户列表','系统用户','POST'),(5,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/user/setUserInfo','设置用户信息','系统用户','PUT'),(6,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/user/setSelfInfo','设置自身信息(必选)','系统用户','PUT'),(7,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/user/getUserInfo','获取自身信息(必选)','系统用户','GET'),(8,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/user/setUserAuthorities','设置权限组','系统用户','POST'),(9,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/user/changePassword','修改密码（建议选择)','系统用户','POST'),(10,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/user/setUserAuthority','修改用户角色(必选)','系统用户','POST'),(11,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/user/resetPassword','重置用户密码','系统用户','POST'),(12,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/api/createApi','创建api','api','POST'),(13,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/api/deleteApi','删除Api','api','POST'),(14,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/api/updateApi','更新Api','api','POST'),(15,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/api/getApiList','获取api列表','api','POST'),(16,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/api/getAllApis','获取所有api','api','POST'),(17,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/api/getApiById','获取api详细信息','api','POST'),(18,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/api/deleteApisByIds','批量删除api','api','DELETE'),(19,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/authority/copyAuthority','拷贝角色','角色','POST'),(20,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/authority/createAuthority','创建角色','角色','POST'),(21,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/authority/deleteAuthority','删除角色','角色','POST'),(22,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/authority/updateAuthority','更新角色信息','角色','PUT'),(23,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/authority/getAuthorityList','获取角色列表','角色','POST'),(24,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/authority/setDataAuthority','设置角色资源权限','角色','POST'),(25,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/casbin/updateCasbin','更改角色api权限','casbin','POST'),(26,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/casbin/getPolicyPathByAuthorityId','获取权限列表','casbin','POST'),(27,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/menu/addBaseMenu','新增菜单','菜单','POST'),(28,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/menu/getMenu','获取菜单树(必选)','菜单','POST'),(29,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/menu/deleteBaseMenu','删除菜单','菜单','POST'),(30,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/menu/updateBaseMenu','更新菜单','菜单','POST'),(31,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/menu/getBaseMenuById','根据id获取菜单','菜单','POST'),(32,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/menu/getMenuList','分页获取基础menu列表','菜单','POST'),(33,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/menu/getBaseMenuTree','获取用户动态路由','菜单','POST'),(34,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/menu/getMenuAuthority','获取指定角色menu','菜单','POST'),(35,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/menu/addMenuAuthority','增加menu和角色关联关系','菜单','POST'),(36,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/fileUploadAndDownload/findFile','寻找目标文件（秒传）','分片上传','GET'),(37,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/fileUploadAndDownload/breakpointContinue','断点续传','分片上传','POST'),(38,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/fileUploadAndDownload/breakpointContinueFinish','断点续传完成','分片上传','POST'),(39,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/fileUploadAndDownload/removeChunk','上传完成移除文件','分片上传','POST'),(40,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/fileUploadAndDownload/upload','文件上传示例','文件上传与下载','POST'),(41,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/fileUploadAndDownload/deleteFile','删除文件','文件上传与下载','POST'),(42,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/fileUploadAndDownload/editFileName','文件名或者备注编辑','文件上传与下载','POST'),(43,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/fileUploadAndDownload/getFileList','获取上传文件列表','文件上传与下载','POST'),(44,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/system/getServerInfo','获取服务器信息','系统服务','POST'),(45,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/system/getSystemConfig','获取配置文件内容','系统服务','POST'),(46,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/system/setSystemConfig','设置配置文件内容','系统服务','POST'),(47,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/customer/customer','更新客户','客户','PUT'),(48,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/customer/customer','创建客户','客户','POST'),(49,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/customer/customer','删除客户','客户','DELETE'),(50,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/customer/customer','获取单一客户','客户','GET'),(51,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/customer/customerList','获取客户列表','客户','GET'),(52,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/autoCode/getDB','获取所有数据库','代码生成器','GET'),(53,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/autoCode/getTables','获取数据库表','代码生成器','GET'),(54,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/autoCode/createTemp','自动化代码','代码生成器','POST'),(55,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/autoCode/preview','预览自动化代码','代码生成器','POST'),(56,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/autoCode/getColumn','获取所选table的所有字段','代码生成器','GET'),(57,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/autoCode/createPlug','自动创建插件包','代码生成器','POST'),(58,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/autoCode/installPlugin','安装插件','代码生成器','POST'),(59,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/autoCode/pubPlug','打包插件','代码生成器','POST'),(60,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/autoCode/createPackage','生成包(package)','包（pkg）生成器','POST'),(61,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/autoCode/getPackage','获取所有包(package)','包（pkg）生成器','POST'),(62,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/autoCode/delPackage','删除包(package)','包（pkg）生成器','POST'),(63,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/autoCode/getMeta','获取meta信息','代码生成器历史','POST'),(64,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/autoCode/rollback','回滚自动生成代码','代码生成器历史','POST'),(65,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/autoCode/getSysHistory','查询回滚记录','代码生成器历史','POST'),(66,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/autoCode/delSysHistory','删除回滚记录','代码生成器历史','POST'),(67,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/sysDictionaryDetail/updateSysDictionaryDetail','更新字典内容','系统字典详情','PUT'),(68,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/sysDictionaryDetail/createSysDictionaryDetail','新增字典内容','系统字典详情','POST'),(69,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/sysDictionaryDetail/deleteSysDictionaryDetail','删除字典内容','系统字典详情','DELETE'),(70,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/sysDictionaryDetail/findSysDictionaryDetail','根据ID获取字典内容','系统字典详情','GET'),(71,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/sysDictionaryDetail/getSysDictionaryDetailList','获取字典内容列表','系统字典详情','GET'),(72,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/sysDictionary/createSysDictionary','新增字典','系统字典','POST'),(73,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/sysDictionary/deleteSysDictionary','删除字典','系统字典','DELETE'),(74,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/sysDictionary/updateSysDictionary','更新字典','系统字典','PUT'),(75,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/sysDictionary/findSysDictionary','根据ID获取字典','系统字典','GET'),(76,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/sysDictionary/getSysDictionaryList','获取字典列表','系统字典','GET'),(77,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/sysOperationRecord/createSysOperationRecord','新增操作记录','操作记录','POST'),(78,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/sysOperationRecord/findSysOperationRecord','根据ID获取操作记录','操作记录','GET'),(79,'2024-04-17 18:04:15.658','2024-07-02 14:52:44.328',NULL,'/sysOperationRecord/getSysOperationRecordList','获取操作记录列表','操作记录','GET'),(80,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/sysOperationRecord/deleteSysOperationRecord','删除操作记录','操作记录','DELETE'),(81,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/sysOperationRecord/deleteSysOperationRecordByIds','批量删除操作历史','操作记录','DELETE'),(82,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/simpleUploader/upload','插件版分片上传','断点续传(插件版)','POST'),(83,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/simpleUploader/checkFileMd5','文件完整度验证','断点续传(插件版)','GET'),(84,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/simpleUploader/mergeFileMd5','上传完成合并文件','断点续传(插件版)','GET'),(85,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/email/emailTest','发送测试邮件','email','POST'),(86,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/email/emailSend','发送邮件示例','email','POST'),(87,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/authorityBtn/setAuthorityBtn','设置按钮权限','按钮权限','POST'),(88,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/authorityBtn/getAuthorityBtn','获取已有按钮权限','按钮权限','POST'),(89,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/authorityBtn/canRemoveAuthorityBtn','删除按钮','按钮权限','POST'),(90,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/sysExportTemplate/createSysExportTemplate','新增导出模板','表格模板','POST'),(91,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/sysExportTemplate/deleteSysExportTemplate','删除导出模板','表格模板','DELETE'),(92,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/sysExportTemplate/deleteSysExportTemplateByIds','批量删除导出模板','表格模板','DELETE'),(93,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/sysExportTemplate/updateSysExportTemplate','更新导出模板','表格模板','PUT'),(94,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/sysExportTemplate/findSysExportTemplate','根据ID获取导出模板','表格模板','GET'),(95,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/sysExportTemplate/getSysExportTemplateList','获取导出模板列表','表格模板','GET'),(96,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/sysExportTemplate/exportExcel','导出Excel','表格模板','GET'),(97,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/sysExportTemplate/exportTemplate','下载模板','表格模板','GET'),(98,'2024-04-17 18:04:15.658','2024-04-17 18:04:15.658',NULL,'/sysExportTemplate/importExcel','导入Excel','表格模板','POST');
/*!40000 ALTER TABLE `sys_apis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_authorities`
--

DROP TABLE IF EXISTS `sys_authorities`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_authorities` (
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `authority_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `authority_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '角色名',
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父角色ID',
  `default_router` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'dashboard' COMMENT '默认菜单',
  PRIMARY KEY (`authority_id`),
  UNIQUE KEY `uni_sys_authorities_authority_id` (`authority_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9529 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_authorities`
--

LOCK TABLES `sys_authorities` WRITE;
/*!40000 ALTER TABLE `sys_authorities` DISABLE KEYS */;
INSERT INTO `sys_authorities` VALUES ('2024-06-06 10:15:08.311','2024-10-18 18:43:45.517',NULL,1,'Common',0,'dashboard'),('2024-06-06 10:15:22.300','2024-10-18 18:43:59.714',NULL,2,'SRE',0,'dashboard'),('2024-07-03 15:18:58.673','2024-07-03 15:19:26.866',NULL,201,'审计',2,'dashboard'),('2024-04-17 18:04:15.728','2024-10-18 18:44:13.712',NULL,888,'普通用户',0,'dashboard');
/*!40000 ALTER TABLE `sys_authorities` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_authority_btns`
--

DROP TABLE IF EXISTS `sys_authority_btns`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_authority_btns` (
  `authority_id` bigint unsigned DEFAULT NULL COMMENT '角色ID',
  `sys_menu_id` bigint unsigned DEFAULT NULL COMMENT '菜单ID',
  `sys_base_menu_btn_id` bigint unsigned DEFAULT NULL COMMENT '菜单按钮ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_authority_btns`
--

LOCK TABLES `sys_authority_btns` WRITE;
/*!40000 ALTER TABLE `sys_authority_btns` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_authority_btns` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_authority_menus`
--

DROP TABLE IF EXISTS `sys_authority_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_authority_menus` (
  `sys_base_menu_id` bigint unsigned NOT NULL,
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_base_menu_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_authority_menus`
--

LOCK TABLES `sys_authority_menus` WRITE;
/*!40000 ALTER TABLE `sys_authority_menus` DISABLE KEYS */;
INSERT INTO `sys_authority_menus` VALUES (1,1),(1,2),(1,201),(1,888),(3,2),(3,201),(3,888),(4,888),(5,888),(6,888),(7,2),(7,888),(8,888),(9,201),(9,888),(10,888),(15,888),(18,888),(23,2),(23,888);
/*!40000 ALTER TABLE `sys_authority_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_auto_code_histories`
--

DROP TABLE IF EXISTS `sys_auto_code_histories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_auto_code_histories` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `package` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `business_db` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `table_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `menu_id` bigint unsigned DEFAULT NULL,
  `request_meta` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `auto_code_path` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `injection_meta` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `struct_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `struct_cn_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `api_ids` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `flag` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_sys_auto_code_histories_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_auto_code_histories`
--

LOCK TABLES `sys_auto_code_histories` WRITE;
/*!40000 ALTER TABLE `sys_auto_code_histories` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_auto_code_histories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_auto_codes`
--

DROP TABLE IF EXISTS `sys_auto_codes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_auto_codes` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `package_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '包名',
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '展示名',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  KEY `idx_sys_auto_codes_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_auto_codes`
--

LOCK TABLES `sys_auto_codes` WRITE;
/*!40000 ALTER TABLE `sys_auto_codes` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_auto_codes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_base_menu_btns`
--

DROP TABLE IF EXISTS `sys_base_menu_btns`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_base_menu_btns` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '按钮关键key',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `sys_base_menu_id` bigint unsigned DEFAULT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menu_btns_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_base_menu_btns`
--

LOCK TABLES `sys_base_menu_btns` WRITE;
/*!40000 ALTER TABLE `sys_base_menu_btns` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_base_menu_btns` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_base_menu_parameters`
--

DROP TABLE IF EXISTS `sys_base_menu_parameters`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_base_menu_parameters` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `sys_base_menu_id` bigint unsigned DEFAULT NULL,
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menu_parameters_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_base_menu_parameters`
--

LOCK TABLES `sys_base_menu_parameters` WRITE;
/*!40000 ALTER TABLE `sys_base_menu_parameters` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_base_menu_parameters` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_base_menus`
--

DROP TABLE IF EXISTS `sys_base_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_base_menus` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `menu_level` bigint unsigned DEFAULT NULL,
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint DEFAULT NULL COMMENT '排序标记',
  `active_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '附加属性',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `default_menu` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `title` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '附加属性',
  `icon` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '附加属性',
  `close_tab` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=56 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_base_menus`
--

LOCK TABLES `sys_base_menus` WRITE;
/*!40000 ALTER TABLE `sys_base_menus` DISABLE KEYS */;
INSERT INTO `sys_base_menus` VALUES (1,'2024-04-17 18:04:16.013','2024-07-02 20:19:29.500',NULL,0,0,'dashboard','dashboard',0,'view/dashboard/index.vue',1,'',0,0,'快捷入口','customer-home',0),(3,'2024-04-17 18:04:16.013','2024-07-03 15:09:17.662',NULL,0,0,'admin','superAdmin',0,'view/superAdmin/index.vue',3,'',0,0,'Manage','user',0),(4,'2024-04-17 18:04:16.013','2024-04-17 18:04:16.013',NULL,0,3,'authority','authority',0,'view/superAdmin/authority/authority.vue',1,'',0,0,'角色管理','avatar',0),(5,'2024-04-17 18:04:16.013','2024-04-17 18:04:16.013',NULL,0,3,'menu','menu',0,'view/superAdmin/menu/menu.vue',2,'',1,0,'菜单管理','tickets',0),(6,'2024-04-17 18:04:16.013','2024-04-17 18:04:16.013',NULL,0,3,'api','api',0,'view/superAdmin/api/api.vue',3,'',1,0,'api管理','platform',0),(7,'2024-04-17 18:04:16.013','2024-04-17 18:04:16.013',NULL,0,3,'user','user',0,'view/superAdmin/user/user.vue',4,'',0,0,'用户管理','coordinate',0),(8,'2024-04-17 18:04:16.013','2024-04-17 18:04:16.013',NULL,0,3,'dictionary','dictionary',0,'view/superAdmin/dictionary/sysDictionary.vue',5,'',0,0,'字典管理','notebook',0),(9,'2024-04-17 18:04:16.013','2024-04-17 18:04:16.013',NULL,0,3,'operation','operation',0,'view/superAdmin/operation/sysOperationRecord.vue',6,'',0,0,'操作历史','pie-chart',0),(10,'2024-04-17 18:04:16.013','2024-04-17 18:04:16.013',NULL,0,0,'person','person',1,'view/person/person.vue',4,'',0,0,'个人信息','message',0),(11,'2024-04-17 18:04:16.013','2024-06-28 15:12:04.020',NULL,0,0,'example','example',0,'view/example/index.vue',7,'',0,0,'示例文件','management',0),(12,'2024-04-17 18:04:16.013','2024-04-17 18:04:16.013',NULL,0,11,'upload','upload',0,'view/example/upload/upload.vue',5,'',0,0,'媒体库（上传下载）','upload',0),(13,'2024-04-17 18:04:16.013','2024-04-17 18:04:16.013',NULL,0,11,'breakpoint','breakpoint',0,'view/example/breakpoint/breakpoint.vue',6,'',0,0,'断点续传','upload-filled',0),(14,'2024-04-17 18:04:16.013','2024-04-17 18:04:16.013',NULL,0,11,'customer','customer',0,'view/example/customer/customer.vue',7,'',0,0,'客户列表（资源示例）','avatar',0),(15,'2024-04-17 18:04:16.013','2024-04-17 18:04:16.013',NULL,0,0,'systemTools','systemTools',0,'view/systemTools/index.vue',5,'',0,0,'系统工具','tools',0),(16,'2024-04-17 18:04:16.013','2024-04-17 18:04:16.013',NULL,0,15,'autoCode','autoCode',0,'view/systemTools/autoCode/index.vue',1,'',1,0,'代码生成器','cpu',0),(17,'2024-04-17 18:04:16.013','2024-04-17 18:04:16.013',NULL,0,15,'formCreate','formCreate',0,'view/systemTools/formCreate/index.vue',2,'',1,0,'表单生成器','magic-stick',0),(18,'2024-04-17 18:04:16.013','2024-04-17 18:04:16.013',NULL,0,15,'system','system',0,'view/systemTools/system/system.vue',3,'',0,0,'系统配置','operation',0),(19,'2024-04-17 18:04:16.013','2024-04-17 18:04:16.013',NULL,0,15,'autoCodeAdmin','autoCodeAdmin',0,'view/systemTools/autoCodeAdmin/index.vue',1,'',0,0,'自动化代码管理','magic-stick',0),(20,'2024-04-17 18:04:16.013','2024-04-17 18:04:16.013',NULL,0,15,'autoCodeEdit/:id','autoCodeEdit',1,'view/systemTools/autoCode/index.vue',0,'',0,0,'自动化代码-${id}','magic-stick',0),(21,'2024-04-17 18:04:16.013','2024-04-17 18:04:16.013',NULL,0,15,'autoPkg','autoPkg',0,'view/systemTools/autoPkg/autoPkg.vue',0,'',0,0,'自动化package','folder',0),(23,'2024-04-17 18:04:16.013','2024-06-28 15:11:17.473',NULL,0,0,'state','state',0,'view/system/state.vue',8,'',0,0,'服务器状态','cloudy',0),(24,'2024-04-17 18:04:16.013','2024-04-17 18:38:56.571',NULL,0,0,'plugin','plugin',1,'view/routerHolder.vue',6,'',0,0,'插件系统','cherry',0),(26,'2024-04-17 18:04:16.013','2024-04-17 18:04:16.013',NULL,0,24,'installPlugin','installPlugin',0,'view/systemTools/installPlugin/index.vue',1,'',0,0,'插件安装','box',0),(27,'2024-04-17 18:04:16.013','2024-04-17 18:04:16.013',NULL,0,24,'autoPlug','autoPlug',0,'view/systemTools/autoPlug/autoPlug.vue',2,'',0,0,'插件模板','folder',0),(28,'2024-04-17 18:04:16.013','2024-04-17 18:04:16.013',NULL,0,24,'pubPlug','pubPlug',0,'view/systemTools/pubPlug/pubPlug.vue',3,'',0,0,'打包插件','files',0),(29,'2024-04-17 18:04:16.013','2024-04-17 18:04:16.013',NULL,0,24,'plugin-email','plugin-email',0,'plugin/email/view/index.vue',4,'',0,0,'邮件插件','message',0),(30,'2024-04-17 18:04:16.013','2024-04-17 18:04:16.013',NULL,0,15,'exportTemplate','exportTemplate',0,'view/systemTools/exportTemplate/exportTemplate.vue',10,'',0,0,'表格模板','reading',0);
/*!40000 ALTER TABLE `sys_base_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_data_authority_id`
--

DROP TABLE IF EXISTS `sys_data_authority_id`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_data_authority_id` (
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`,`data_authority_id_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_data_authority_id`
--

LOCK TABLES `sys_data_authority_id` WRITE;
/*!40000 ALTER TABLE `sys_data_authority_id` DISABLE KEYS */;
INSERT INTO `sys_data_authority_id` VALUES (888,888),(888,8881),(888,9528),(9528,8881);
/*!40000 ALTER TABLE `sys_data_authority_id` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_dictionaries`
--

DROP TABLE IF EXISTS `sys_dictionaries`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_dictionaries` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字典名（中）',
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字典名（英）',
  `status` tinyint(1) DEFAULT NULL COMMENT '状态',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionaries_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dictionaries`
--

LOCK TABLES `sys_dictionaries` WRITE;
/*!40000 ALTER TABLE `sys_dictionaries` DISABLE KEYS */;
INSERT INTO `sys_dictionaries` VALUES (1,'2024-04-17 18:04:15.887','2024-04-17 18:04:15.902',NULL,'性别','gender',1,'性别字典'),(2,'2024-04-17 18:04:15.887','2024-04-17 18:04:15.929',NULL,'数据库int类型','int',1,'int类型对应的数据库类型'),(3,'2024-04-17 18:04:15.887','2024-04-17 18:04:15.949',NULL,'数据库时间日期类型','time.Time',1,'数据库时间日期类型'),(4,'2024-04-17 18:04:15.887','2024-04-17 18:04:15.966',NULL,'数据库浮点型','float64',1,'数据库浮点型'),(5,'2024-04-17 18:04:15.887','2024-04-17 18:04:15.981',NULL,'数据库字符串','string',1,'数据库字符串'),(6,'2024-04-17 18:04:15.887','2024-04-17 18:04:15.995',NULL,'数据库bool类型','bool',1,'数据库bool类型');
/*!40000 ALTER TABLE `sys_dictionaries` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_dictionary_details`
--

DROP TABLE IF EXISTS `sys_dictionary_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_dictionary_details` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '展示值',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字典值',
  `extend` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '扩展值',
  `status` tinyint(1) DEFAULT NULL COMMENT '启用状态',
  `sort` bigint DEFAULT NULL COMMENT '排序标记',
  `sys_dictionary_id` bigint unsigned DEFAULT NULL COMMENT '关联标记',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionary_details_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dictionary_details`
--

LOCK TABLES `sys_dictionary_details` WRITE;
/*!40000 ALTER TABLE `sys_dictionary_details` DISABLE KEYS */;
INSERT INTO `sys_dictionary_details` VALUES (1,'2024-04-17 18:04:15.904','2024-04-17 18:04:15.904',NULL,'男','1','',1,1,1),(2,'2024-04-17 18:04:15.904','2024-04-17 18:04:15.904',NULL,'女','2','',1,2,1),(3,'2024-04-17 18:04:15.931','2024-04-17 18:04:15.931',NULL,'smallint','1','mysql',1,1,2),(4,'2024-04-17 18:04:15.931','2024-04-17 18:04:15.931',NULL,'mediumint','2','mysql',1,2,2),(5,'2024-04-17 18:04:15.931','2024-04-17 18:04:15.931',NULL,'int','3','mysql',1,3,2),(6,'2024-04-17 18:04:15.931','2024-04-17 18:04:15.931',NULL,'bigint','4','mysql',1,4,2),(7,'2024-04-17 18:04:15.931','2024-04-17 18:04:15.931',NULL,'int2','5','pgsql',1,5,2),(8,'2024-04-17 18:04:15.931','2024-04-17 18:04:15.931',NULL,'int4','6','pgsql',1,6,2),(9,'2024-04-17 18:04:15.931','2024-04-17 18:04:15.931',NULL,'int6','7','pgsql',1,7,2),(10,'2024-04-17 18:04:15.931','2024-04-17 18:04:15.931',NULL,'int8','8','pgsql',1,8,2),(11,'2024-04-17 18:04:15.951','2024-04-17 18:04:15.951',NULL,'date','','',1,0,3),(12,'2024-04-17 18:04:15.951','2024-04-17 18:04:15.951',NULL,'time','1','mysql',1,1,3),(13,'2024-04-17 18:04:15.951','2024-04-17 18:04:15.951',NULL,'year','2','mysql',1,2,3),(14,'2024-04-17 18:04:15.951','2024-04-17 18:04:15.951',NULL,'datetime','3','mysql',1,3,3),(15,'2024-04-17 18:04:15.951','2024-04-17 18:04:15.951',NULL,'timestamp','5','mysql',1,5,3),(16,'2024-04-17 18:04:15.951','2024-04-17 18:04:15.951',NULL,'timestamptz','6','pgsql',1,5,3),(17,'2024-04-17 18:04:15.968','2024-04-17 18:04:15.968',NULL,'float','','',1,0,4),(18,'2024-04-17 18:04:15.968','2024-04-17 18:04:15.968',NULL,'double','1','mysql',1,1,4),(19,'2024-04-17 18:04:15.968','2024-04-17 18:04:15.968',NULL,'decimal','2','mysql',1,2,4),(20,'2024-04-17 18:04:15.968','2024-04-17 18:04:15.968',NULL,'numeric','3','pgsql',1,3,4),(21,'2024-04-17 18:04:15.968','2024-04-17 18:04:15.968',NULL,'smallserial','4','pgsql',1,4,4),(22,'2024-04-17 18:04:15.983','2024-04-17 18:04:15.983',NULL,'char','','',1,0,5),(23,'2024-04-17 18:04:15.983','2024-04-17 18:04:15.983',NULL,'varchar','1','mysql',1,1,5),(24,'2024-04-17 18:04:15.983','2024-04-17 18:04:15.983',NULL,'tinyblob','2','mysql',1,2,5),(25,'2024-04-17 18:04:15.983','2024-04-17 18:04:15.983',NULL,'tinytext','3','mysql',1,3,5),(26,'2024-04-17 18:04:15.983','2024-04-17 18:04:15.983',NULL,'text','4','mysql',1,4,5),(27,'2024-04-17 18:04:15.983','2024-04-17 18:04:15.983',NULL,'blob','5','mysql',1,5,5),(28,'2024-04-17 18:04:15.983','2024-04-17 18:04:15.983',NULL,'mediumblob','6','mysql',1,6,5),(29,'2024-04-17 18:04:15.983','2024-04-17 18:04:15.983',NULL,'mediumtext','7','mysql',1,7,5),(30,'2024-04-17 18:04:15.983','2024-04-17 18:04:15.983',NULL,'longblob','8','mysql',1,8,5),(31,'2024-04-17 18:04:15.983','2024-04-17 18:04:15.983',NULL,'longtext','9','mysql',1,9,5),(32,'2024-04-17 18:04:15.997','2024-04-17 18:04:15.997',NULL,'tinyint','1','mysql',1,0,6),(33,'2024-04-17 18:04:15.997','2024-04-17 18:04:15.997',NULL,'bool','2','pgsql',1,0,6);
/*!40000 ALTER TABLE `sys_dictionary_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_export_template_condition`
--

DROP TABLE IF EXISTS `sys_export_template_condition`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_export_template_condition` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `template_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '模板标识',
  `from` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '条件取的key',
  `column` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '作为查询条件的字段',
  `operator` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '操作符',
  PRIMARY KEY (`id`),
  KEY `idx_sys_export_template_condition_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_export_template_condition`
--

LOCK TABLES `sys_export_template_condition` WRITE;
/*!40000 ALTER TABLE `sys_export_template_condition` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_export_template_condition` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_export_template_join`
--

DROP TABLE IF EXISTS `sys_export_template_join`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_export_template_join` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `template_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '模板标识',
  `joins` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '关联',
  `table` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '关联表',
  `on` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '关联条件',
  PRIMARY KEY (`id`),
  KEY `idx_sys_export_template_join_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_export_template_join`
--

LOCK TABLES `sys_export_template_join` WRITE;
/*!40000 ALTER TABLE `sys_export_template_join` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_export_template_join` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_export_templates`
--

DROP TABLE IF EXISTS `sys_export_templates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_export_templates` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `db_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '数据库名称',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '模板名称',
  `table_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '表名称',
  `template_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '模板标识',
  `template_info` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `limit` bigint DEFAULT NULL COMMENT '导出限制',
  `order` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `idx_sys_export_templates_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_export_templates`
--

LOCK TABLES `sys_export_templates` WRITE;
/*!40000 ALTER TABLE `sys_export_templates` DISABLE KEYS */;
INSERT INTO `sys_export_templates` VALUES (1,'2024-04-17 18:04:16.246','2024-04-17 18:04:16.246',NULL,'','api','sys_apis','api','{\n\"path\":\"路径\",\n\"method\":\"方法（大写）\",\n\"description\":\"方法介绍\",\n\"api_group\":\"方法分组\"\n}',0,'');
/*!40000 ALTER TABLE `sys_export_templates` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_operation_records`
--

DROP TABLE IF EXISTS `sys_operation_records`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_operation_records` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `ip` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求路径',
  `status` bigint DEFAULT NULL COMMENT '请求状态',
  `latency` bigint DEFAULT NULL COMMENT '延迟',
  `agent` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '代理',
  `error_message` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '错误信息',
  `body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '请求Body',
  `resp` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '响应Body',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`),
  KEY `idx_sys_operation_records_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_operation_records`
--

LOCK TABLES `sys_operation_records` WRITE;
/*!40000 ALTER TABLE `sys_operation_records` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_operation_records` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user_authority`
--

DROP TABLE IF EXISTS `sys_user_authority`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_user_authority` (
  `sys_user_id` bigint unsigned NOT NULL,
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_user_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_authority`
--

LOCK TABLES `sys_user_authority` WRITE;
/*!40000 ALTER TABLE `sys_user_authority` DISABLE KEYS */;
INSERT INTO `sys_user_authority` VALUES (1,1),(1,2),(1,201),(1,888);
/*!40000 ALTER TABLE `sys_user_authority` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_users`
--

DROP TABLE IF EXISTS `sys_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `uuid` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户UUID',
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户登录密码',
  `nick_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '系统用户' COMMENT '用户昵称',
  `side_mode` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'dark' COMMENT '用户侧边主题',
  `header_img` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '/asset/person.png' COMMENT '用户头像',
  `base_color` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '#fff' COMMENT '基础颜色',
  `active_color` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '#1890ff' COMMENT '活跃颜色',
  `authority_id` bigint unsigned DEFAULT '888' COMMENT '用户角色ID',
  `phone` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户手机号',
  `email` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户邮箱',
  `enable` bigint DEFAULT '1' COMMENT '用户是否被冻结 1正常 2冻结',
  PRIMARY KEY (`id`),
  KEY `idx_sys_users_deleted_at` (`deleted_at`),
  KEY `idx_sys_users_uuid` (`uuid`),
  KEY `idx_sys_users_username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_users`
--

LOCK TABLES `sys_users` WRITE;
/*!40000 ALTER TABLE `sys_users` DISABLE KEYS */;
INSERT INTO `sys_users` VALUES (1,'2024-04-17 18:04:16.186','2024-07-09 16:24:42.033',NULL,'34bf3134-807c-49c7-bd73-dcd5c85ea0a7','admin','$2a$10$OW4WXCU.wDf0QjvclHGgVOAnVAOSLLD1mMAjO/sVJbH6i9WCGVsk2','admin888','dark','/asset/person.png','#fff','#1890ff',888,'13995941111','13995941111@qq.com',1),(2,'2024-04-17 18:04:16.186','2024-06-28 17:48:16.773','2024-07-03 14:48:38.505','e6595ebf-14ea-4307-92a8-7604bde1e5a2','a303176530','$2a$10$FUMk04cEyZypG0mpUc12JOXVUDWo9PxgoFi.ON3Tt4otJAXCOyA6S','用户1','dark','/asset/person.png','#fff','#1890ff',888,'17611111111','333333333@qq.com',1);
/*!40000 ALTER TABLE `sys_users` ENABLE KEYS */;
UNLOCK TABLES;
SET @@SESSION.SQL_LOG_BIN = @MYSQLDUMP_TEMP_LOG_BIN;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-10-21 14:56:53
