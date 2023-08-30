CREATE DATABASE  IF NOT EXISTS `devbook` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `devbook`;
-- MySQL dump 10.13  Distrib 8.0.33, for macos13 (arm64)
--
-- Host: 127.0.0.1    Database: devbook
-- ------------------------------------------------------
-- Server version	8.0.33

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `devices`
--

DROP TABLE IF EXISTS `devices`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `devices` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `address` varchar(50) NOT NULL,
  `latitude` float NOT NULL,
  `longitude` float NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `address` (`address`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `devices`
--

LOCK TABLES `devices` WRITE;
/*!40000 ALTER TABLE `devices` DISABLE KEYS */;
INSERT INTO `devices` VALUES (1,'Romulo','Rosendo Chagas',-5.10344,-38.3687),(2,'Ingrid','Rosendo Chagas 291',-5.10324,-38.3689),(3,'Olga','Doutor Viana 305',-5.10367,-38.3688),(5,'Fulano 2','Rosendo Chagas 330',-5.10345,-38.3686),(6,'Fulano 3','Rosendo Chagas 304',-5.10341,-38.3688);
/*!40000 ALTER TABLE `devices` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `records`
--

DROP TABLE IF EXISTS `records`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `records` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `value` double NOT NULL,
  `createdAt` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `device_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_device_id` (`device_id`),
  CONSTRAINT `FK_device_id` FOREIGN KEY (`device_id`) REFERENCES `devices` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=228 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `records`
--

LOCK TABLES `records` WRITE;
/*!40000 ALTER TABLE `records` DISABLE KEYS */;
INSERT INTO `records` VALUES (193,100,'2023-08-29 19:26:09',1),(194,1000,'2023-08-29 21:32:24',2),(195,100,'2023-08-29 21:32:47',2),(196,0,'2023-08-29 21:33:02',2),(197,0,'2023-08-29 21:33:20',1),(198,1,'2023-08-29 21:33:49',2),(199,200,'2023-08-29 21:41:26',1),(200,1000,'2023-08-29 21:41:32',2),(201,100,'2023-08-29 21:42:00',2),(202,1,'2023-08-29 21:48:43',3),(203,100,'2023-08-29 21:48:49',3),(204,200,'2023-08-29 21:48:57',3),(205,300,'2023-08-29 21:49:05',3),(206,1000,'2023-08-29 21:49:14',3),(207,1,'2023-08-29 21:49:19',3),(208,0,'2023-08-29 21:49:24',3),(209,1000,'2023-08-29 21:49:33',3),(210,1,'2023-08-29 22:41:11',3),(211,0,'2023-08-29 22:41:18',1),(214,0,'2023-08-29 22:48:53',1),(215,1000,'2023-08-29 22:49:38',5),(216,1000,'2023-08-29 22:49:47',6),(217,0,'2023-08-29 23:39:48',1),(218,100,'2023-08-29 23:39:52',1),(219,0,'2023-08-29 23:40:02',1),(220,100,'2023-08-29 23:40:14',2),(221,250,'2023-08-29 23:40:23',6),(222,0,'2023-08-29 23:40:36',1),(223,0,'2023-08-29 23:40:42',2),(224,0,'2023-08-29 23:40:49',3),(226,0,'2023-08-29 23:40:59',5),(227,0,'2023-08-29 23:41:03',6);
/*!40000 ALTER TABLE `records` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `usuarios`
--

DROP TABLE IF EXISTS `usuarios`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `usuarios` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nome` varchar(50) NOT NULL,
  `nick` varchar(50) NOT NULL,
  `email` varchar(50) NOT NULL,
  `senha` varchar(100) NOT NULL,
  `criadoEm` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `nick` (`nick`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `usuarios`
--

LOCK TABLES `usuarios` WRITE;
/*!40000 ALTER TABLE `usuarios` DISABLE KEYS */;
INSERT INTO `usuarios` VALUES (1,'Romulo','Romulohsc','romulohsc7@gmail.com','$2a$10$w0wnAiJCZdAF1ncZlT3l6e8wFpS0dxWmNsngvVxcb6EIQB2jWSNQ2','2023-08-06 18:10:15');
/*!40000 ALTER TABLE `usuarios` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-08-30 17:31:28
