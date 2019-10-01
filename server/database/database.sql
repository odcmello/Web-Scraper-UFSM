DROP DATABASE IF EXISTS egressos_informatica; 
CREATE DATABASE egressos_informatica; 
USE egressos_informatica; 

CREATE TABLE egressos(
	id int auto_increment primary key, 
    name varchar(255) not null, 
    job_title varchar(255) not null, 
    company varchar(255) not null, 
    location varchar(255) not null, 
    url varchar(255) not null, 
    curso varchar(255) not null, 
    ano_evasao int not null
)ENGINE = InnoDB;

##TEST DATA
INSERT INTO egressos(name, job_title, company, location, url, curso, ano_evasao) VALUES("Everson Feltrin","Consultor de vendas", "Conplan Sistemas de Informática Ltda", "Santa Maria", "https://www.google.com/", "Sistemas de Informação", 2018);
INSERT INTO egressos(name, job_title, company, location, url, curso, ano_evasao) VALUES("Evandro", "Bolsista","Mestrado - UFSM", "Santa Maria", "https://www.google.com/", "Sistemas de Informação", 2018);
INSERT INTO egressos(name, job_title, company, location, url, curso, ano_evasao) VALUES("Ezequiel R. Ribeiro", "Programer",  "Conplan Sistemas de Informática Ltda", "Santa Maria", "https://www.google.com/", "Ciência da Computação", 2019);
INSERT INTO egressos(name, job_title, company, location, url, curso, ano_evasao) VALUES("Yuri Kunde Schlesner", "Site Reliability Engineer",  "Google Inc", "Seattle", "https://www.google.com/", "Ciência da Computação", 2015);
INSERT INTO egressos(name, job_title, company, location, url, curso, ano_evasao) VALUES("Roger Couto", "Técnico em informática", "Prefeitura Municipal De Restinga Sêca", "Restinga Sêca", "https://www.google.com/", "Sistemas de Informação", 2018);
INSERT INTO egressos(name, job_title, company, location, url, curso, ano_evasao) VALUES("Luis Henrique Medeiros", "Desenvolvedor", "PipeRun CRM", "Porto Alegre", "https://www.google.com/", "Sistemas de Informação", 2018);
INSERT INTO egressos(name, job_title, company, location, url, curso, ano_evasao) VALUES("Leonardo Steil", "Desenvolver back-end", "PipeRun CRM", "Porto Alegre", "https://www.google.com/", "Sistemas de Informação", 2018);
INSERT INTO egressos(name, job_title, company, location, url, curso, ano_evasao) VALUES("Anthony Tailer", "Desenvolvedor Front-end", "COWMED", "Santa Maria", "https://www.google.com/", "Sistemas de Informação" , 2018);
