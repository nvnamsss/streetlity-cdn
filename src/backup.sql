-- MySQL dump 10.13  Distrib 5.7.30, for Linux (x86_64)
--
-- Host: localhost    Database: services
-- ------------------------------------------------------
-- Server version	8.0.19

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
-- Table structure for table `atm`
--

DROP TABLE IF EXISTS `atm`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `atm` (
  `id` int NOT NULL AUTO_INCREMENT,
  `lat` float DEFAULT NULL,
  `lon` float DEFAULT NULL,
  `note` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `bank_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `bank_id` (`bank_id`),
  CONSTRAINT `atm_ibfk_1` FOREIGN KEY (`bank_id`) REFERENCES `bank` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `atm`
--

LOCK TABLES `atm` WRITE;
/*!40000 ALTER TABLE `atm` DISABLE KEYS */;
INSERT INTO `atm` VALUES (1,1,1,'','help',5),(2,8.09734,8.40725,'','dasdasd',2),(3,8.20691,8.48151,'','asddas',2),(4,11.0112,106.625,'','dai nam',2);
/*!40000 ALTER TABLE `atm` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `atm_ucf`
--

DROP TABLE IF EXISTS `atm_ucf`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `atm_ucf` (
  `id` int NOT NULL AUTO_INCREMENT,
  `lat` float DEFAULT NULL,
  `lon` float DEFAULT NULL,
  `note` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `bank_id` int DEFAULT NULL,
  `confident` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `bank_id` (`bank_id`),
  CONSTRAINT `atm_ucf_ibfk_1` FOREIGN KEY (`bank_id`) REFERENCES `bank` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `atm_ucf`
--

LOCK TABLES `atm_ucf` WRITE;
/*!40000 ALTER TABLE `atm_ucf` DISABLE KEYS */;
/*!40000 ALTER TABLE `atm_ucf` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `bank`
--

DROP TABLE IF EXISTS `bank`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `bank` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bank`
--

LOCK TABLES `bank` WRITE;
/*!40000 ALTER TABLE `bank` DISABLE KEYS */;
INSERT INTO `bank` VALUES (1,'Agribank'),(2,'VietcomBank'),(3,'BiDV'),(4,'NaniBank'),(5,'VCB');
/*!40000 ALTER TABLE `bank` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `fuel`
--

DROP TABLE IF EXISTS `fuel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `fuel` (
  `id` int NOT NULL AUTO_INCREMENT,
  `lat` float DEFAULT NULL,
  `lon` float DEFAULT NULL,
  `note` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `fuel`
--

LOCK TABLES `fuel` WRITE;
/*!40000 ALTER TABLE `fuel` DISABLE KEYS */;
INSERT INTO `fuel` VALUES (1,1,2,'',''),(2,8.23043,8.59573,'asaaaa','asd'),(3,8.29974,8.2016,'','sdadsada'),(4,7.78424,8.49545,'','sdsada'),(5,8.2221,8.46112,'sadad','sadasd'),(6,10.7678,106.667,'','help me akhd'),(7,10.7962,106.768,'','dsasd'),(8,10.8069,106.531,'','adsad'),(9,10.7355,106.536,'','asddadww');
/*!40000 ALTER TABLE `fuel` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `fuel_ucf`
--

DROP TABLE IF EXISTS `fuel_ucf`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `fuel_ucf` (
  `id` int NOT NULL AUTO_INCREMENT,
  `lat` float DEFAULT NULL,
  `lon` float DEFAULT NULL,
  `note` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `confident` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `fuel_ucf`
--

LOCK TABLES `fuel_ucf` WRITE;
/*!40000 ALTER TABLE `fuel_ucf` DISABLE KEYS */;
/*!40000 ALTER TABLE `fuel_ucf` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `maintenance`
--

DROP TABLE IF EXISTS `maintenance`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `maintenance` (
  `id` int NOT NULL AUTO_INCREMENT,
  `lat` float DEFAULT NULL,
  `lon` float DEFAULT NULL,
  `note` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `iscallable` tinyint(1) DEFAULT NULL,
  `phone` varchar(10) DEFAULT NULL,
  `owner` varchar(32) DEFAULT NULL,
  `name` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `maintenance`
--

LOCK TABLES `maintenance` WRITE;
/*!40000 ALTER TABLE `maintenance` DISABLE KEYS */;
INSERT INTO `maintenance` VALUES (1,8.0487,8.3743,'help','pls',NULL,NULL,'thosua','idk'),(2,8.03923,8.39461,'sacas','abc xyz',NULL,NULL,'','dafas'),(3,9.40504,8.15752,'adada asd ad','Nigeria',NULL,NULL,'','sadadasdsad'),(4,8.09564,8.38707,'asdwa','asdad',NULL,NULL,'','adwa'),(5,7.99966,8.55928,'dadadadad','bro',NULL,NULL,'','adaa'),(6,7.98284,8.35346,'pls','25 cho do do',NULL,NULL,'','help me'),(7,1,5,'','near near sidney',NULL,NULL,'',''),(8,8.06539,8.36574,'gsfdsfsd','sadaswae',NULL,NULL,'','asdasdadad'),(9,7.81987,8.25176,'androidx.appcompat.widget.AppCompatEditText{386df9b VFED..CL. .F...... 15,129-465,175 #7f0a0088 app:id/edt_store_note}','asssssssss',NULL,NULL,'','assssss'),(10,10.8101,106.709,'androidx.appcompat.widget.AppCompatEditText{e110b67 VFED..CL. ........ 15,129-465,175 #7f0a0084 app:id/edt_store_note}','150 Đinh Bộ Lĩnh, Phường 26, Bình Thạnh, Hồ Chí Minh, Vietnam',NULL,NULL,'','abc');
/*!40000 ALTER TABLE `maintenance` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `maintenance_history`
--

DROP TABLE IF EXISTS `maintenance_history`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `maintenance_history` (
  `id` int NOT NULL AUTO_INCREMENT,
  `maintenance_user` varchar(32) NOT NULL,
  `common_user` varchar(32) NOT NULL,
  `timestamp` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `maintenance_history`
--

LOCK TABLES `maintenance_history` WRITE;
/*!40000 ALTER TABLE `maintenance_history` DISABLE KEYS */;
/*!40000 ALTER TABLE `maintenance_history` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `maintenance_ucf`
--

DROP TABLE IF EXISTS `maintenance_ucf`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `maintenance_ucf` (
  `id` int NOT NULL AUTO_INCREMENT,
  `lat` float DEFAULT NULL,
  `lon` float DEFAULT NULL,
  `note` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `phone` varchar(10) DEFAULT NULL,
  `iscallable` tinyint(1) DEFAULT NULL,
  `name` varchar(32) DEFAULT NULL,
  `confident` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `maintenance_ucf`
--

LOCK TABLES `maintenance_ucf` WRITE;
/*!40000 ALTER TABLE `maintenance_ucf` DISABLE KEYS */;
/*!40000 ALTER TABLE `maintenance_ucf` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `nodes`
--

DROP TABLE IF EXISTS `nodes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `nodes` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `lat` float DEFAULT NULL,
  `lon` float DEFAULT NULL,
  `streets` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `nodes`
--

LOCK TABLES `nodes` WRITE;
/*!40000 ALTER TABLE `nodes` DISABLE KEYS */;
/*!40000 ALTER TABLE `nodes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `review`
--

DROP TABLE IF EXISTS `review`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `review` (
  `id` int NOT NULL AUTO_INCREMENT,
  `score` int NOT NULL,
  `body` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `review`
--

LOCK TABLES `review` WRITE;
/*!40000 ALTER TABLE `review` DISABLE KEYS */;
/*!40000 ALTER TABLE `review` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `services`
--

DROP TABLE IF EXISTS `services`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `services` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `services`
--

LOCK TABLES `services` WRITE;
/*!40000 ALTER TABLE `services` DISABLE KEYS */;
/*!40000 ALTER TABLE `services` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `streets`
--

DROP TABLE IF EXISTS `streets`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `streets` (
  `id` varchar(15) NOT NULL,
  `generalinfo` varchar(1024) DEFAULT NULL,
  `cost` float DEFAULT NULL,
  `oneway` tinyint(1) DEFAULT '0',
  `onewayfrom` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `streets`
--

LOCK TABLES `streets` WRITE;
/*!40000 ALTER TABLE `streets` DISABLE KEYS */;
/*!40000 ALTER TABLE `streets` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `toilet`
--

DROP TABLE IF EXISTS `toilet`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `toilet` (
  `id` int NOT NULL AUTO_INCREMENT,
  `lat` float DEFAULT NULL,
  `lon` float DEFAULT NULL,
  `note` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `isworking` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `toilet`
--

LOCK TABLES `toilet` WRITE;
/*!40000 ALTER TABLE `toilet` DISABLE KEYS */;
INSERT INTO `toilet` VALUES (1,8.19907,8.58748,'asdd','dadasd',NULL),(2,8.34172,8.14528,'','asdddda',NULL);
/*!40000 ALTER TABLE `toilet` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `toilet_ucf`
--

DROP TABLE IF EXISTS `toilet_ucf`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `toilet_ucf` (
  `id` int NOT NULL AUTO_INCREMENT,
  `lat` float DEFAULT NULL,
  `lon` float DEFAULT NULL,
  `note` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `confident` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `toilet_ucf`
--

LOCK TABLES `toilet_ucf` WRITE;
/*!40000 ALTER TABLE `toilet_ucf` DISABLE KEYS */;
/*!40000 ALTER TABLE `toilet_ucf` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-05-30  3:00:01
