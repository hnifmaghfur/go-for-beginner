-- MySQL dump 10.13  Distrib 5.7.17, for macos10.12 (x86_64)
--
-- Host: localhost    Database: tokoijah
-- ------------------------------------------------------
-- Server version	5.7.22

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
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `products` (
  `sku` varchar(20) NOT NULL,
  `product_name` varchar(100) DEFAULT NULL,
  `stocks` int(11) DEFAULT NULL,
  PRIMARY KEY (`sku`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES ('ffffff-ccc-ikik','Zalekia Plain Casual Jeans (L,Broken White)',35),('SSI-D00791015-LL-BWH','Zalekia Plain Casual Blouse (L,Broken White)',154),('SSI-D00791077-MM-BWH','Zalekia Plain Casual Blouse (M,Broken White)',138),('SSI-D00791091-XL-BWH','Zalekia Plain Casual Blouse (XL,Broken White)',137),('SSI-D00864612-LL-NAV','Deklia Plain Casual Blouse (L,Navy)',8),('SSI-D00864614-XL-NAV','Deklia Plain Casual Blouse (XL,Navy)',97),('SSI-D00864652-SS-NAV','Deklia Plain Casual Blouse (S,Navy)',2),('SSI-D00864661-MM-NAV','Deklia Plain Casual Blouse (M,Navy)',13),('SSI-D01037807-X3-BWH','Dellaya Plain Loose Big Blouse (XXXL,Broken White)',74),('SSI-D01037812-X3-BLA','Dellaya Plain Loose Big Blouse (XXXL,Black)',54),('SSI-D01037822-XX-BLA','Dellaya Plain Loose Big Blouse (XXL,Black)',8),('SSI-D01220307-XL-SAL','Devibav Plain Trump Blouse (XL,Salem)',182),('SSI-D01220322-MM-YEL','Devibav Plain Trump Blouse (M,Yellow)',121),('SSI-D01220334-XL-YEL','Devibav Plain Trump Blouse (XL,Yellow)',110),('SSI-D01220338-XX-SAL','Devibav Plain Trump Blouse (XXL,Salem)',65),('SSI-D01220346-LL-SAL','Devibav Plain Trump Blouse (L,Salem)',151),('SSI-D01220349-LL-YEL','Devibav Plain Trump Blouse (L,Yellow)',101),('SSI-D01220355-XX-YEL','Devibav Plain Trump Blouse (XXL,Yellow)',140),('SSI-D01220357-SS-YEL','Devibav Plain Trump Blouse (S,Yellow)',74),('SSI-D01220388-MM-SAL','Devibav Plain Trump Blouse (M,Salem)',216),('SSI-D01322234-LL-WHI','Thafqya Plain Raglan Blouse (L,White)',105),('SSI-D01322275-XL-WHI','Thafqya Plain Raglan Blouse (XL,White)',116),('SSI-D01326201-XL-KHA','Siunfhi Ethnic Trump Blouse (XL,Khaki)',186),('SSI-D01326205-MM-NAV','Siunfhi Ethnic Trump Blouse (M,Navy)',143),('SSI-D01326223-MM-KHA','Siunfhi Ethnic Trump Blouse (M,Khaki)',209),('SSI-D01326286-LL-KHA','Siunfhi Ethnic Trump Blouse (L,Khaki)',210),('SSI-D01326299-LL-NAV','Siunfhi Ethnic Trump Blouse (L,Navy)',127),('SSI-D01401050-MM-RED','Zeomila Zipper Casual Blouse (M,Red)',73),('SSI-D01401064-XL-RED','Zeomila Zipper Casual Blouse (XL,Red)',44),('SSI-D01401071-LL-RED','Zeomila Zipper Casual Blouse (L,Red)',76),('SSI-D01466013-XX-BLA','Salyara Plain Casual Big Blouse (XXL,Black)',77),('SSI-D01466064-X3-BLA','Salyara Plain Casual Big Blouse (XXXL,Black)',52);
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `stock_ins`
--

DROP TABLE IF EXISTS `stock_ins`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `stock_ins` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_date` datetime DEFAULT NULL,
  `sku` varchar(20) NOT NULL,
  `buy_price` int(11) DEFAULT NULL,
  `qty` int(11) DEFAULT NULL,
  `kwitansi` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `stock_ins`
--

LOCK TABLES `stock_ins` WRITE;
/*!40000 ALTER TABLE `stock_ins` DISABLE KEYS */;
INSERT INTO `stock_ins` VALUES (33,'2018-05-06 22:03:31','ffffff-ccc-ikik',120000,5,'1234-1234-4321'),(34,'2018-05-06 22:05:56','ffffff-ccc-ikik',120000,2,'1234-1234-4322'),(35,'2018-05-06 23:08:49','ffffff-ccc-ikik',120000,7,'1234-1234-4322'),(36,'2018-05-06 23:23:12','ffffff-ccc-ikik',120000,35,'1234-1234-4322');
/*!40000 ALTER TABLE `stock_ins` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `stock_outs`
--

DROP TABLE IF EXISTS `stock_outs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `stock_outs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `transaction_id` varchar(20) DEFAULT NULL,
  `sku` varchar(20) DEFAULT NULL,
  `qty` int(11) DEFAULT NULL,
  `note` varchar(100) DEFAULT NULL,
  `created_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `stock_outs`
--

LOCK TABLES `stock_outs` WRITE;
/*!40000 ALTER TABLE `stock_outs` DISABLE KEYS */;
INSERT INTO `stock_outs` VALUES (4,'ID-12345-9091','ffffff-ccc-ikik',1,'Pesanan ID-12345-9091','2018-05-06 22:18:18'),(5,'ID-12345-9092','ffffff-ccc-ikik',1,'Pesanan ID-12345-9092','2018-05-06 22:19:18'),(6,'ID-12345-9093','ffffff-ccc-ikik',3,'Pesanan ID-12345-9093','2018-05-06 22:20:57'),(7,'ID-12345-9094','ffffff-ccc-ikik',1,'Pesanan ID-12345-9094','2018-05-06 22:22:30'),(8,'ID-12345-9095','ffffff-ccc-ikik',1,'Pesanan ID-12345-9095','2018-05-06 22:22:48'),(9,'ID-20180506-699523','ffffff-ccc-ikik',1,'Pesanan ID-20180506-699523','2018-05-06 23:09:11'),(10,'','ffffff-ccc-ikik',1,'Barang Hilang','2018-05-06 23:10:35'),(11,'ID-20180506-130131','ffffff-ccc-ikik',1,'Pesanan ID-20180506-130131','2018-05-06 23:13:57'),(12,'ID-20180506-633432','ffffff-ccc-ikik',1,'Pesanan ID-20180506-633432','2018-05-06 23:14:10'),(13,'ID-20180506-385103','ffffff-ccc-ikik',1,'Pesanan ID-20180506-385103','2018-05-06 23:14:12'),(14,'ID-20180506-203094','ffffff-ccc-ikik',1,'Pesanan ID-20180506-203094','2018-05-06 23:23:21'),(15,'ID-20180507-699523','ffffff-ccc-ikik',1,'Pesanan ID-20180507-699523','2018-05-07 23:09:04');
/*!40000 ALTER TABLE `stock_outs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactions`
--

DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transactions` (
  `id` varchar(20) NOT NULL,
  `created_date` datetime DEFAULT NULL,
  `sku` varchar(20) DEFAULT NULL,
  `qty` int(11) DEFAULT NULL,
  `buy_price` int(11) DEFAULT NULL,
  `sell_price` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactions`
--

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;
INSERT INTO `transactions` VALUES ('ID-12345-9091','2018-05-06 22:18:18','ffffff-ccc-ikik',1,120000,130000),('ID-12345-9092','2018-05-06 22:19:18','ffffff-ccc-ikik',1,120000,130000),('ID-12345-9093','2018-05-06 22:20:57','ffffff-ccc-ikik',3,120000,130000),('ID-12345-9094','2018-05-06 22:22:30','ffffff-ccc-ikik',1,120000,130000),('ID-12345-9095','2018-05-06 22:22:48','ffffff-ccc-ikik',1,120000,130000),('ID-20180506-130131','2018-05-06 23:13:57','ffffff-ccc-ikik',1,120000,130000),('ID-20180506-203094','2018-05-06 23:23:21','ffffff-ccc-ikik',1,120000,130000),('ID-20180506-385103','2018-05-06 23:14:12','ffffff-ccc-ikik',1,120000,130000),('ID-20180506-633432','2018-05-06 23:14:10','ffffff-ccc-ikik',1,120000,130000),('ID-20180506-699523','2018-05-06 23:09:11','ffffff-ccc-ikik',1,120000,130000),('ID-20180507-699523','2018-05-07 23:09:04','ffffff-ccc-ikik',1,120000,130000);
/*!40000 ALTER TABLE `transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'tokoijah'
--

--
-- Dumping routines for database 'tokoijah'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-06-07 21:48:04
