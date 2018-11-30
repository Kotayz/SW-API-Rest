# SW-API-Rest

API Rest para Cadastro de Planetas da franquia Star Wars

Requerimentos:

Go versão 1.6 ou maior
Glide versão 0.13
MongoDB versão 3.4.18

Instalação MongoDB:
https://www.digitalocean.com/community/tutorials/como-instalar-o-mongodb-no-ubuntu-16-04-pt

Instalação Glide:
https://github.com/Masterminds/glide

Instalação e configuração go:
https://www.digitalocean.com/community/tutorials/how-to-install-go-1-6-on-ubuntu-16-04

Dependencias do projeto:
    github.com/BurntSushi/toml
    github.com/gin-gonic/gin
    gopkg.in/mgo.v2
    gopkg.in/mgo.v2/bson

Para instalar dependencias, executar o comando: glide install

Para iniciar o serviço, executar o comando: go run main.go