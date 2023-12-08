-- MySQL dump 10.13  Distrib 8.0.35, for Linux (x86_64)
--
-- Host: recything.cm4tgjp8nfxv.us-east-1.rds.amazonaws.com    Database: recything
-- ------------------------------------------------------
-- Server version	8.0.33

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

SET @@GLOBAL.GTID_PURGED=/*!80000 '+'*/ '';

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` varchar(191) NOT NULL,
  `email` varchar(191) NOT NULL,
  `password` longtext NOT NULL,
  `fullname` longtext NOT NULL,
  `phone` longtext,
  `address` longtext,
  `date_of_birth` longtext,
  `purpose` longtext,
  `point` bigint DEFAULT '0',
  `badge` enum('platinum','gold','silver','bronze') DEFAULT 'bronze',
  `is_verified` tinyint(1) DEFAULT '0',
  `verification_token` longtext,
  `otp` longtext,
  `otp_expiration` bigint DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `delete_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  KEY `idx_users_delete_at` (`delete_at`),
  KEY `fk_users_badges` (`badge`),
  CONSTRAINT `fk_users_badges` FOREIGN KEY (`badge`) REFERENCES `achievements` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('058b69e8-416e-4a3e-b7cb-e62267c9d873','michaelbrown@gmail.com','$2a$10$nUom/ayllm/YMfJi7WX.aeon0Z.MBJuvc4iCxDiIolVZuS95ZMRKK','Michael Brown',NULL,NULL,NULL,NULL,0,'platinum',NULL,NULL,NULL,NULL,'2023-11-19 14:31:55.672',NULL,NULL),('19c2d8e2-86ee-42d7-ba81-de4548921604','madvier84@gmail.com','$2a$10$mjxZu8Eobic1H9rQdVWrT.MyMI/a0rwv5EgT0iq0q8C/77sQ23q32','Vi','','','','',0,'bronze',1,'m_M8sS8KfdXdPlpPSc_d9WjoXRNKq6mQTnAVmLGKImw=','',0,'2023-11-22 06:27:48.457','2023-11-22 06:27:48.457',NULL),('1aef6030-ad76-4d24-a664-807d26277618','jiddan@gmail.com','$2a$10$658TqH39ru6771FL3I1s5.XrO.r0nxyI2R/5SZFZxttj3FsX.UFRG','jiddan','','','','',0,'bronze',1,'haTQvusywe2Dkilpoi1615SW-mCtFLFDjx9go_LiKgg=','',0,'2023-11-25 16:55:11.767','2023-11-25 16:55:11.767',NULL),('25df2772-37de-493e-922c-c43b7e46779c','emilywilliams@gmail.com','$2y$10$c35H0r1iZO/sOJL0L45YJeg9NXX8yAXncYQMK4kjzKYQDRquwuVQS','Emily Williams',NULL,NULL,NULL,NULL,0,'bronze',1,NULL,NULL,NULL,NULL,NULL,'2023-11-21 18:54:58.375'),('2732491e-ab0c-49f6-aae0-0ed17d95d402','hilal@example.com','$2y$10$fDBo3vO7yS7rhoj0KCm2FOez.02UX7WbAvoEz/7nV7GTHRqkYcuKe','Hilal',NULL,NULL,NULL,NULL,0,'bronze',1,NULL,NULL,NULL,NULL,NULL,'2023-11-27 10:56:00.002'),('3053f71f-2bc8-4095-a323-ca96a2cc67a1','christopherwilson@gmail.com','$2y$10$x5b/zylta3zArVgxhqaQ9e9EWdeq/Vz2jQGnirmQ.9iyrI./N2IAa','Christopher Wilson',NULL,NULL,NULL,NULL,0,'bronze',1,NULL,NULL,NULL,NULL,NULL,'2023-11-27 10:57:50.904'),('3777ffd6-2215-4a00-a474-8f420a94f026','user1@gmail.com','$2a$10$6PTOtsii8Nba8uZqcLEc2.qbOo.mrz8tcQRJ22F4KYfEnSm7xhJxW','user1','','','','',0,'bronze',0,'eFc_1IHssjJ2AaEOooOe0RGJno3cTnspUUxzizgeXhs=','',0,'2023-11-22 08:33:38.574','2023-11-22 08:33:38.574','2023-11-27 11:16:02.366'),('39921a03-2df4-4fb9-8399-696d75e626e4','mamad@example.com','$2y$10$RUR96P5BnRFiRzPRNLOVsOnKFxerObYV0vqOHdlP.iIR//sxa3oly','Mamad',NULL,NULL,NULL,NULL,0,'bronze',1,NULL,NULL,NULL,NULL,NULL,'2023-11-27 11:19:49.457'),('3c2d474c-1f8f-4cf8-ac6d-5f04b5caa4ad','tara@example.com','$2y$10$Pz0MnlXUQQAqc4lDIP1oQOnR3kSkL/7yYrI870/JcBOBGItr2LwL6','Tara',NULL,NULL,NULL,NULL,0,'bronze',1,NULL,NULL,NULL,NULL,NULL,'2023-11-27 11:11:22.225'),('468bf650-fe11-4cc6-9bb7-b0e7700a7bf8','klvn@gmail.com','$2a$10$Nr/8GmKmfNKr7fNXfP2nau3V.9APfR6RDryJ2q/EvgSv885q2VFES','Kelvin Sanjaya','','','','',0,'bronze',0,'W5OzAsNRKbzsWRoSRCh5ghBk3CECjnXOe14bqIzUcuE=','',0,'2023-11-27 08:43:19.403','2023-11-27 08:43:19.403','2023-11-27 12:36:54.365'),('4a8617e7-2653-4966-9f6f-6242d3085b9e','mamang@example.com','$2a$10$J6J7SuThZ9HL44V7xFBv8.blz0hPdO3WIbZuTOA.CMHkVMRjNRW/u','mamang','','','','',0,'bronze',0,'mwCCgN5O3qQnUObrxe21EB5_w0H8b4hO-kQ3zupaXBA=','',0,'2023-11-28 19:12:50.555','2023-11-28 19:12:50.555','2023-11-29 19:58:52.752'),('4c433fb0-fa5c-454c-b85b-b78ab258df23','budi@example.com','$2y$10$j8xUn7l4qHbHtcc03m0BZ.pPhfiAVywlsa0hF56GuljQwx6J8ehaa','Budi',NULL,NULL,NULL,NULL,0,'bronze',1,NULL,NULL,NULL,NULL,NULL,'2023-11-27 14:48:56.893'),('4c76ce45-025d-445f-98af-d4d78c905041','ss@gmail.com','$2a$10$hEj1IwkqhFXEK4OC3IPfwOIqkyKeXCu1MZtLV62Bc0WU85fnRY6r6','sss','','','','',0,'bronze',0,'rgEHGU4IBiCHtubu1zWOtL3hFzzaR4EC9kgpmq0nAAU=','',0,'2023-11-28 09:01:11.481','2023-11-28 09:01:11.481','2023-11-28 13:45:13.660'),('4fc0ee1f-4556-4c18-a076-689b843bf011','jonohasim123@gmail.com','$2a$10$waSLa0BkVDCscYVZxC.R2Omk6uuMCmvPCGr54or7dz5ML.o6BJQuS','salim','','','','',0,'bronze',0,'YT4Ei4Hw0Vo85AR-ermj__rF2F3-Pfjfmun7iqOKouo=','',0,'2023-11-27 07:33:38.048','2023-11-27 07:33:38.048','2023-11-27 18:45:18.795'),('61480e82-2080-4880-8eb9-51c9ff6ad16a','hanieffathulb03@gmail.com','$2a$10$.S0eWUxFUr4Ar8VrWl/ZXuCWDqVZkzR/itucR79/ysTuuu08vBhSq','jiddan2','','','','',0,'bronze',1,'gHNDJ-ISpgF0JIBdCDWrL-PfRyxTA9oujAdYWlqwp5g=','',0,'2023-11-27 07:44:06.555','2023-11-27 07:44:29.712','2023-11-27 17:25:15.843'),('6356f39d-eff6-4bf8-9220-313913cc2c57','mpsamin27@gmail.com','$2a$10$vbcjTI1C4iEmwQ7RB8imQuPiIZPUhjE6Wqt8fmhbCShNclyswkXzy','klvnqwerty','','','','',0,'bronze',0,'e73HQHKvIekaROx_FNy2OCHxnzrxzn-DpXVnQuTginA=','',0,'2023-11-27 08:44:31.719','2023-11-27 08:44:31.719','2023-11-27 11:54:26.852'),('6e504941-5fc5-46e3-9ac1-203cbc539eb0','madvier83@gmail.com','$2a$10$VRU90Q9txQHOqGWDzUMvI.fMwTa86IBLOLvW8iz2b3ojYw1QOcEgG','madvier83@gmail.com','','','','',0,'bronze',1,'Qf71j_h6rUkTW7dKG3PHgvRnv9N6VHE17d0XGg8ouL4=','',0,'2023-12-06 14:55:21.911','2023-12-06 14:56:00.734',NULL),('70b79f67-bf33-47e1-8fc3-a333d7b1919c','melissataylor@gmail.com','$2y$10$52E8MLDJbKgqB0ETQR1XAu..WvbnJAlsXZYpilzhz53ySkg8tvhDK','Melissa Taylor',NULL,NULL,NULL,NULL,0,'bronze',1,NULL,NULL,NULL,NULL,NULL,'2023-11-28 18:10:49.198'),('714abfd8-9cd2-4457-be4f-0039bd4ae91c','belajarpuspita4@gmail.com','$2a$10$o8KarL9ANcVsGqKiUakkLeGVnM5V1ukdy9TrW6.UR1HICmW2qON6a','puspita','','solo','2002-02-09','coba coba',500,'bronze',1,'fGXScoREu7k83W8IWeoNxyLmNpOEV0-R87g5ShLfc14=','',0,'2023-11-27 14:06:42.846','2023-11-27 14:15:17.492','2023-11-29 21:30:32.697'),('77ed10fc-aad7-4783-94a3-84b83b5ece9c','tehpucuk8700@gmail.com','$2a$10$uTwgP/hZgL8pcZRp6lYIPO/02x5T4Rm5j69sOaAr4QJxsmzb6c3EO','tehpucuk','','','','',158280,'bronze',0,'ptGiT4Wye1A1kvcB42vUFnZ0HUXy75X7p60f68Capt0=','',0,'2023-11-28 09:00:46.835','2023-12-04 04:50:19.677',NULL),('7b27f83b-8512-48de-a623-109aece77a2b','samanthadavis@gmail.com','$2y$10$LEqyK2tRMmc8tZ1xRPpAe.qmAu6aYFtiJwYiGUeCu.uaXyM1WgIS.','Samantha Davis',NULL,NULL,NULL,NULL,0,'bronze',1,NULL,NULL,NULL,NULL,NULL,'2023-12-02 19:28:58.128'),('7c2aea15-6151-455a-8840-2cb38b12294a','aang@example.com','$2y$10$7bKOwV0..UQrNmxtBI9B/e6./0gsNtPNsKhSX4Xvnux56l.yEV1A.','Aang','','','','',0,'bronze',1,'','M0K1',1700661862,'0000-00-00 00:00:00.000','2023-11-22 20:59:22.999','2023-12-29 10:42:15.501'),('90b2ea80-f1f1-4e53-b44b-5a93eb94ebd0','jenniferanderson@gmail.com','$2y$10$z/xYulkgSYv0HMmZJD/UZei65uajETlZSrhTRZh8bZay.jr3avs4O','Jennifer Anderson',NULL,NULL,NULL,NULL,0,'bronze',1,NULL,NULL,NULL,NULL,NULL,'2023-11-28 18:09:32.229'),('98871198-e2df-40ca-86c0-c5e75aac408c','pus@example.com','$2a$10$SPjImD1xaxUn2ay21EubJ.MEK.CyLFm2wuLLEYoH7tKYzwZcQhDwG','puspita','','','','',0,'bronze',0,'iew6pp7vNyZ7rKeKgfLrLw97sz6lJCt6O_SbWiS0BDI=','',0,'2023-11-27 14:05:22.443','2023-11-27 14:05:22.443','0000-00-00 00:00:00.000'),('9ce9458f-cecd-42c9-8bcb-d380325d9d23','madvier86@gmail.com','$2a$10$K1tQNpD2LoeMjdOvVnPwvOEtjD6mgmc60ScUm3LN5no1ADQidvkjG','madvier85','','','','',0,'bronze',0,'JaEEm2nSXa1aZkR7TQLICv5cmBakcg_tNBYa5GInz2I=','',0,'2023-11-22 08:04:40.869','2023-11-22 08:04:40.869','2023-11-29 21:34:23.922'),('aa7da12a-2500-4742-bc21-e7b71971bb56','recythingalta@gmail.com','$2a$10$iA0GUTWFK74NvOCav45BbugC.7V5cPiID/FQkfFRoiMQw1jXja7sC','recything','','','','',0,'bronze',1,'A3Xi_c5bSXyQ8H8X8KQSUtufT1d8prjmC_3T3c0f7Ls=','',0,'2023-11-21 18:37:35.792','2023-11-21 18:37:35.792','0000-00-00 00:00:00.000'),('b92c8420-7183-4344-bdac-9b84a7041e12','janesmith@gmail.com','$2y$10$Pezj37LbWkdIQy5MOLC/RuxOETKwsNnfch35mqPgZI0W.dDNwYyJ6','Jane Smith',NULL,NULL,NULL,NULL,0,'bronze',1,NULL,NULL,NULL,NULL,NULL,NULL),('bcc1e2c4-f9d0-4f41-86c1-da7ea8e3d42a','puspita@example.com','$2a$10$/MGbzf1qsoRNzr86tbfaBed3JsqhK5MR3RSrCgPSjy6zdXZZfNRSq','puspita','','','','',0,'bronze',0,'ENBooHJkBlOPgUPOYiOksiNUGqza9gusb5gq2M4p19E=','',0,'2023-11-27 14:09:51.011','2023-11-27 14:09:51.011','2023-12-02 19:35:25.830'),('be16dcf3-b78e-45a4-a09c-fecf88998e7f','robertjohnson@gmail.com','$2y$10$r076jxiQLAvtGbWiybFif.2.V4a3b2j8FG9F4A7ZWEB0ySfMQuNoK','Robert Johnson',NULL,NULL,NULL,NULL,0,'bronze',1,NULL,NULL,NULL,NULL,'2023-12-07 14:09:51.011',NULL),('c09ea139-2ece-4894-97ab-0633d2b78968','testbudi@gmail.com','$2a$10$NxB5/qgC2FiAnHySekCApObotYxrLOLh24naQT6B32efK3ZIvJJSG','Budi','','','','',0,'bronze',0,'_iEXF0tn9vuCyQRb5xMRc13YciRfWrPjGHYM0mNrro4=','',0,'2023-11-30 07:54:11.131','2023-11-30 07:54:11.131',NULL),('c46791fc-18f3-45e0-8d10-ad82dc9b77bf','andi@example.com','$2y$10$k/WN6uGnKBPL/2O3eVHY3e.N3a46xR1AnmtomCk4yvQzAnJl0h5ba','Andi',NULL,NULL,NULL,NULL,0,'bronze',1,NULL,NULL,NULL,NULL,NULL,NULL),('c67335d9-a4ed-41e0-a18b-93603b3a75db','aaa@a.com','$2a$10$KKv8kpMsDbwSYC2D2Beq5ebQaRH/p6nRS.wVeWerA9E32zSwHf0Fy','aaaaa','','','','',0,'bronze',0,'03E42ZR8XTXb4Sc4JSx9zIP4cNG27coT8BouBIN-SF0=','',0,'2023-11-23 13:37:45.488','2023-11-23 13:37:45.488','2023-11-30 05:01:08.204'),('c84a71c9-526a-4622-9de2-22584f1690e2','recythinguser@gmail.com','$2a$10$JgBQ0dlAZwk0KTXYzlC6AurVKKVAxQWIEMifUE8U1ukMlAQro/Jqu','recything','6282287653245','kota purwokerto','2002-03-07','saya ingin melestarkan lingkungan',0,'platinum',1,'05G0','1700242696',2023,'2023-11-17 17:33:17.559',NULL,'0000-00-00 00:00:00.000'),('d1fd458e-50eb-4694-90f8-654f134a9789','tiara@example.com','$2y$10$al8Ps2piH8t8xYdsYwM1mOX/6HB6P.XGrn5wkf2updF4gD6Ds4VqO','Tiara',NULL,NULL,NULL,NULL,0,'bronze',1,NULL,NULL,NULL,NULL,NULL,NULL),('df06e9f7-c0fe-4072-8c7f-1058b1471c12','ahmad@example.com','$2a$10$KJHzR6E1M9HgOokK.l.YWe4J2m1IG1AhzOZWo/SieMbEZwU0UfpDq','recything','082298567350','kota jakarta','2023-03-03','saya ingin melestarikan lingkungan',100,'bronze',1,'','44JV',1701369600,'0000-00-00 00:00:00.000','2023-12-06 09:10:25.383',NULL),('eba9c534-0aed-4331-ae07-a46de435196e','toto@example.com','$2a$10$IiwfC/zJQxvKpIAxMpnAT.UCZ5t/RKn9i6oV2MwQdjVcOBC52W9h2','i Nyoman Toto','6282287653245',' putwokerto','2023-11-28','terhubung dengan komunitas',450,'bronze',1,'','',0,'0000-00-00 00:00:00.000','2023-12-05 17:11:59.821',NULL),('f39bc0e6-efc3-4f55-99e3-0abea1d949f5','doni@example.com','$2y$10$D4D1VN76pWOzKtaXMOc6X.m2sgocuI9uJj2YGw4.i4FWnHpSrzS8m','Doni',NULL,NULL,NULL,NULL,0,'bronze',1,NULL,NULL,NULL,NULL,NULL,'2023-11-29 19:47:23.342'),('f5819de3-6343-48f1-9cde-e51f65c578c6','jiddan@example.com','$2a$10$72ObO/B1swaZrMEta1E2nOlhxBgcY9Bg8uL8FYqaodI28jr5yK90O','jiddan','','','','',0,'bronze',0,'MnMvZ8kgiJIGNqCW9iwtSU_RtlZ-lTH1PwDJqp6_bUg=','',0,'2023-12-07 23:52:58.523','2023-12-07 23:52:58.523',NULL),('f9458708-6d40-4b51-b6fe-654e44e7cb1d','davidmartinez@gmail.com','$2y$10$bF7gxWRxGs7JcWsJcTrQ4OYhyycPflDKxE7iZbnOhW0iLIYTYrisq','David Martinez',NULL,NULL,NULL,NULL,0,'bronze',1,NULL,NULL,NULL,NULL,NULL,'2023-11-29 03:48:03.957'),('fd2b1d96-8964-4080-a00a-fad50beec459','johndoe@gmail.com','$2y$10$LKcZn5iSqHJrXd6/GV0XCe0SfTnQd4kQRdr7tVMMoyp0VMDkJ4Rfm','John Doe',NULL,NULL,NULL,NULL,0,'bronze',1,NULL,NULL,NULL,NULL,NULL,'2023-11-30 05:08:02.791'),('ff55191d-f552-4acf-9c47-5a57e8932fe4','madvier85@gmail.com','$2a$10$ekdHXJSyegQjBV6Ni9CghOEfitFObrmUTBaS7UkpEi.ZTuI9qNJXm','Advie Rifaldy','','','','',0,'bronze',1,'iM-kCXdBUbuxauWm9HY5B3HSatqpX2fnptwuT991yW8=','',0,'2023-11-22 06:29:18.776','2023-11-22 06:29:18.776','2023-11-28 08:40:37.134');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
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

-- Dump completed on 2023-12-08  7:55:07
