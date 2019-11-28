drop database if exists egressos_informatica; 
create database egressos_informatica; 
use egressos_informatica;

create table egressos(
	id int auto_increment primary key, 
    nome varchar(256) not null, 
    job_title varchar(256), 
    company varchar(256),
    location varchar(256),
    url varchar(256),
    curso varchar(256) not null, 
    ano_evasao int 
)ENGINE = InnoDb; 

load data local infile 'C:/Users/gabri/Documents/Projetos/Web-Scraper-UFSM/web-scrape/egressos_encontrados.csv'
into table egressos 
character set utf8mb4 
fields terminated by ';'
optionally enclosed by '"'
lines terminated by '\n'
ignore 1 lines 
(nome, job_title, company, location, url, curso, ano_evasao); 
