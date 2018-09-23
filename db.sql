-- MySQL dump 10.13  Distrib 5.7.18, for Win64 (x86_64)
--
-- Host: jw0ch9vofhcajqg7.cbetxkdyhwsb.us-east-1.rds.amazonaws.com    Database: kzi50q30ap322cl4
-- ------------------------------------------------------
-- Server version	5.6.34-log

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
-- Table structure for table `agendamento`
--

DROP TABLE IF EXISTS `agendamento`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `agendamento` (
  `codigo` int(11) NOT NULL AUTO_INCREMENT,
  `codigopaciente` int(11) DEFAULT NULL,
  `codigomedico` int(11) DEFAULT NULL,
  `data` date DEFAULT NULL,
  `hora` int(11) DEFAULT NULL,
  `motivo` varchar(150) COLLATE utf8_unicode_ci DEFAULT NULL,
  `alergias` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `status` char(2) COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`codigo`),
  KEY `codigopaciente_idx` (`codigopaciente`),
  KEY `codigomedico_idx` (`codigomedico`),
  CONSTRAINT `codigomedico` FOREIGN KEY (`codigomedico`) REFERENCES `medico` (`codigo`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `codigopaciente` FOREIGN KEY (`codigopaciente`) REFERENCES `paciente` (`codigo`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `agendamento`
--

LOCK TABLES `agendamento` WRITE;
/*!40000 ALTER TABLE `agendamento` DISABLE KEYS */;
INSERT INTO `agendamento` VALUES (1,1,1,'2018-09-17',7,'teste','teste','a'),(5,1,1,'2018-09-17',8,'kkk','1','a'),(7,1,1,'2018-09-17',9,'kk','1','a'),(8,1,1,'2018-09-17',10,'ddasas','1,2,3','f');
/*!40000 ALTER TABLE `agendamento` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `consulta`
--

DROP TABLE IF EXISTS `consulta`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `consulta` (
  `codigo` int(11) NOT NULL AUTO_INCREMENT,
  `codigoagendamento` int(11) DEFAULT NULL,
  `descricao` varchar(45) COLLATE utf8_unicode_ci DEFAULT NULL,
  `exames` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `remedios` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `status` tinyint(2) DEFAULT NULL,
  PRIMARY KEY (`codigo`),
  KEY `codigoagendamento_idx` (`codigoagendamento`),
  CONSTRAINT `codigoagendamento` FOREIGN KEY (`codigoagendamento`) REFERENCES `agendamento` (`codigo`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `consulta`
--

LOCK TABLES `consulta` WRITE;
/*!40000 ALTER TABLE `consulta` DISABLE KEYS */;
/*!40000 ALTER TABLE `consulta` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `medico`
--

DROP TABLE IF EXISTS `medico`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `medico` (
  `codigo` int(11) NOT NULL AUTO_INCREMENT,
  `nome` varchar(45) DEFAULT NULL,
  `email` varchar(45) DEFAULT NULL,
  `senha` varchar(45) DEFAULT NULL,
  `data_nascimento` date DEFAULT NULL,
  `especializacao` varchar(45) DEFAULT NULL,
  `hospital` varchar(45) DEFAULT NULL,
  `crm` varchar(45) DEFAULT NULL,
  `role` varchar(1) DEFAULT NULL,
  `ativo` char(2) DEFAULT NULL,
  PRIMARY KEY (`codigo`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `medico`
--

LOCK TABLES `medico` WRITE;
/*!40000 ALTER TABLE `medico` DISABLE KEYS */;
INSERT INTO `medico` VALUES (1,'MEDICO','MEDICO@CLINICA.COM','medico','2018-09-04','Oftalmologista','regional','123','m','a');
/*!40000 ALTER TABLE `medico` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `paciente`
--

DROP TABLE IF EXISTS `paciente`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `paciente` (
  `codigo` int(11) NOT NULL AUTO_INCREMENT,
  `nome` varchar(45) DEFAULT NULL,
  `email` varchar(45) DEFAULT NULL,
  `senha` varchar(45) DEFAULT NULL,
  `data_nascimento` date DEFAULT NULL,
  `hospital` varchar(45) DEFAULT NULL,
  `carteira` varchar(45) DEFAULT NULL,
  `role` varchar(45) DEFAULT NULL,
  `ativo` char(2) DEFAULT NULL,
  PRIMARY KEY (`codigo`)
) ENGINE=InnoDB AUTO_INCREMENT=81 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `paciente`
--

LOCK TABLES `paciente` WRITE;
/*!40000 ALTER TABLE `paciente` DISABLE KEYS */;
INSERT INTO `paciente` VALUES (1,'PACIENTE','PACIENTE@CLINICA.COM','paciente','2018-08-08','unimed','123456789','p','i'),(80,'teste','teste@teste.com','123','2018-09-18','unimed','123','p','a');
/*!40000 ALTER TABLE `paciente` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'kzi50q30ap322cl4'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-09-18 12:19:08
