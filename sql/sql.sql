CREATE DATABASE IF NOT EXISTS cadastro;
USE cadastro;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    cpf varchar(20) not null unique,
    endereco varchar(150) not null,
    telefone varchar(50) not null,
    dataNascimento varchar(20) not null
) ENGINE=INNODB;