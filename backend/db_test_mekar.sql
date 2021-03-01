-- MySQL dump 10.13  Distrib 8.0.20, for Win64 (x86_64)
--
-- Host: localhost    Database: db_test_mekar
-- ------------------------------------------------------
-- Server version	8.0.20

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
-- Table structure for table `pekerjaan`
--

DROP TABLE IF EXISTS `pekerjaan`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pekerjaan` (
  `id_pekerjaan` int NOT NULL AUTO_INCREMENT,
  `nama_pekerjaan` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id_pekerjaan`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pekerjaan`
--

LOCK TABLES `pekerjaan` WRITE;
/*!40000 ALTER TABLE `pekerjaan` DISABLE KEYS */;
INSERT INTO `pekerjaan` VALUES (1,'PNS'),(2,'Wirausaha'),(3,'Wiraswasta');
/*!40000 ALTER TABLE `pekerjaan` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pendidikan`
--

DROP TABLE IF EXISTS `pendidikan`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pendidikan` (
  `id_pendidikan` int NOT NULL AUTO_INCREMENT,
  `label_pendidikan` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id_pendidikan`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pendidikan`
--

LOCK TABLES `pendidikan` WRITE;
/*!40000 ALTER TABLE `pendidikan` DISABLE KEYS */;
INSERT INTO `pendidikan` VALUES (1,'SD'),(2,'SMP'),(3,'SMA'),(4,'Diploma'),(5,'Sarjana'),(6,'Magister'),(7,'Doktor');
/*!40000 ALTER TABLE `pendidikan` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id_user` int NOT NULL AUTO_INCREMENT,
  `nama_user` varchar(50) DEFAULT NULL,
  `tanggal_lahir` date DEFAULT NULL,
  `no_ktp` varchar(16) DEFAULT NULL,
  `id_pekerjaan` int DEFAULT NULL,
  `id_pendidikan` int DEFAULT NULL,
  PRIMARY KEY (`id_user`),
  KEY `id_pekerjaan` (`id_pekerjaan`),
  KEY `id_pendidikan` (`id_pendidikan`),
  CONSTRAINT `user_ibfk_1` FOREIGN KEY (`id_pekerjaan`) REFERENCES `pekerjaan` (`id_pekerjaan`),
  CONSTRAINT `user_ibfk_2` FOREIGN KEY (`id_pendidikan`) REFERENCES `pendidikan` (`id_pendidikan`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'Viontina','1998-10-12','3327065210980009',2,3);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary view structure for view `user_detail`
--

DROP TABLE IF EXISTS `user_detail`;
/*!50001 DROP VIEW IF EXISTS `user_detail`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `user_detail` AS SELECT 
 1 AS `id_user`,
 1 AS `nama_user`,
 1 AS `tanggal_lahir`,
 1 AS `no_ktp`,
 1 AS `nama_pekerjaan`,
 1 AS `pendidikan_terakhir`*/;
SET character_set_client = @saved_cs_client;

--
-- Final view structure for view `user_detail`
--

/*!50001 DROP VIEW IF EXISTS `user_detail`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `user_detail` AS select `u`.`id_user` AS `id_user`,`u`.`nama_user` AS `nama_user`,`u`.`tanggal_lahir` AS `tanggal_lahir`,`u`.`no_ktp` AS `no_ktp`,`pk`.`nama_pekerjaan` AS `nama_pekerjaan`,`pd`.`label_pendidikan` AS `pendidikan_terakhir` from ((`user` `u` join `pekerjaan` `pk` on((`u`.`id_pekerjaan` = `pk`.`id_pekerjaan`))) join `pendidikan` `pd` on((`u`.`id_pendidikan` = `pd`.`id_pendidikan`))) */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-02-28 14:57:40
