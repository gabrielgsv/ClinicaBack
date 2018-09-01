/* nome do banco */
CREATE DATABASE db_clinica

/* tabela medico */
CREATE TABLE `medico`
(
  `codigo` int
(11) NOT NULL AUTO_INCREMENT,
  `nome` varchar
(45) DEFAULT NULL,
  `email` varchar
(45) DEFAULT NULL,
  `senha` varchar
(45) DEFAULT NULL,
  `area` varchar
(45) DEFAULT NULL,
  `hospital` varchar
(45) DEFAULT NULL,
  `crm` varchar
(45) DEFAULT NULL,
  `anos_atuacao` int
(11) DEFAULT NULL,
  `role` varchar
(1) DEFAULT NULL,
  PRIMARY KEY
(`codigo`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

/* tabela paciente */
CREATE TABLE `paciente`
(
  `codigo` int
(11) NOT NULL AUTO_INCREMENT,
  `nome` varchar
(45) DEFAULT NULL,
  `email` varchar
(45) DEFAULT NULL,
  `senha` varchar
(45) DEFAULT NULL,
  `hospital` varchar
(45) DEFAULT NULL,
  `carteira` varchar
(45) DEFAULT NULL,
  `role` varchar
(45) DEFAULT NULL,
  PRIMARY KEY
(`codigo`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=latin1;
