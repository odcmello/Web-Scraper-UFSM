## CRIAÇÃO DO BANCO DE DADOS
DROP DATABASE IF EXISTS dadosEvasao;
CREATE DATABASE dadosEvasao; 
USE dadosEvasao;

CREATE TABLE dadosevasaopessoa(
	id int auto_increment not null,
	nome_curso varchar(256) not null,
    ano_evasao int,
    nome_pessoa varchar(256),
    primary key(id)
) ENGINE = InnoDB;