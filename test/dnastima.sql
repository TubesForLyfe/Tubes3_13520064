-- MySQL dump 10.13  Distrib 8.0.26, for Win64 (x86_64)
--
-- Host: localhost    Database: dnastima_db
-- ------------------------------------------------------
-- Server version	5.5.5-10.4.22-MariaDB

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

--
-- Current Database: `dnastima_db`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `dnastima_db` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `dnastima_db`;

--
-- Table structure for table `hasilprediksi`
--

DROP TABLE IF EXISTS `hasilprediksi`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `hasilprediksi` (
  `TanggalPrediksi` date NOT NULL,
  `NamaPasien` varchar(100) NOT NULL,
  `PenyakitPrediksi` varchar(100) NOT NULL,
  `TingkatKemiripan` int(11) NOT NULL,
  `Status` int(1) NOT NULL,
  PRIMARY KEY (`TanggalPrediksi`,`NamaPasien`,`PenyakitPrediksi`),
  KEY `PenyakitPrediksi` (`PenyakitPrediksi`),
  CONSTRAINT `hasilprediksi_ibfk_1` FOREIGN KEY (`PenyakitPrediksi`) REFERENCES `jenispenyakit` (`NamaPenyakit`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `hasilprediksi`
--

LOCK TABLES `hasilprediksi` WRITE;
/*!40000 ALTER TABLE `hasilprediksi` DISABLE KEYS */;
INSERT INTO `hasilprediksi` VALUES ('2022-04-24','Aku','COVID-19',79,0),('2022-04-24','Siapa','COVID-19',99,1);
/*!40000 ALTER TABLE `hasilprediksi` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `jenispenyakit`
--

DROP TABLE IF EXISTS `jenispenyakit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `jenispenyakit` (
  `NamaPenyakit` varchar(100) NOT NULL,
  `DNA` varchar(100) NOT NULL,
  PRIMARY KEY (`NamaPenyakit`),
  UNIQUE KEY `DNA` (`DNA`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `jenispenyakit`
--

LOCK TABLES `jenispenyakit` WRITE;
/*!40000 ALTER TABLE `jenispenyakit` DISABLE KEYS */;
INSERT INTO `jenispenyakit` VALUES ('COVID-19','AGCT');
/*!40000 ALTER TABLE `jenispenyakit` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-04-25  3:02:16
