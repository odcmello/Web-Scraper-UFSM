library("readODS")
library("dplyr")
library("RMySQL")

setwd("C:/Users/Otávio/Desktop/Dados Evasao")
filenames <- list.files(full.names=FALSE)
All <- lapply(filenames,function(i){
  teste <- read_ods(i, skip=0)
  teste <- select(teste, NOME_CURSO, ANO_EVASAO, NOME_PESSOA)
})
df <- do.call(rbind.data.frame, All)

setwd("C:/Users/Otávio/Desktop/Tabelas Extras")
sheetsname <- ods_sheets("Alunos Egressos.ods")

novoAll <- lapply(sheetsname,function(i){
  teste <- read_ods("Alunos Egressos.ods", skip=0)
  teste$NOME_CURSO<-"Pós-Graduação"
  teste$ANO_EVASAO<-NA
  names(teste)[names(teste) == "ALUNO"] <- "NOME_PESSOA"
  teste <- select(teste, NOME_CURSO, ANO_EVASAO, NOME_PESSOA)
})
df2 <- do.call(rbind.data.frame,novoAll)

df <- rbind(df,df2)

dbPass <- "password"
dbUser <- "user"
dbPort <- 3306 
dbHost <- "localhost"
dbDatabase <- "dadosEvasao"

vDB <- dbConnect(RMySQL::MySQL(), dbname=dbDatabase, password=dbPass, username=dbUser, port=dbPort)
dbWriteTable(vDB, name="dadosevasaopessoa", value=df, overwrite=FALSE, append=TRUE, sep=",", row.names=FALSE)
dbDisconnect(vDB)