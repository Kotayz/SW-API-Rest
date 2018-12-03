# SW-API-Rest

API Rest para Cadastro de Planetas da franquia Star Wars

Essa API foi desenvolvida em [Golang](https://golang.org/).

### Requerimentos
* Instalação [MongoDB](https://www.digitalocean.com/community/tutorials/como-instalar-o-mongodb-no-ubuntu-16-04-pt) versão 1.6 ou superior

* Instalação [Glide](https://github.com/Masterminds/glide) versão 0.13 ou superior

* Instalação e Configuração [Golang](https://www.digitalocean.com/community/tutorials/how-to-install-go-1-6-on-ubuntu-16-04) versão 3.4.18 ou superior

### Dependencias do projeto
 * github.com/BurntSushi/toml
 * github.com/gin-gonic/gin
 * gopkg.in/mgo.v2
 * gopkg.in/mgo.v2/bson

Para instalar dependencias, executar o comando: glide install

Para iniciar o serviço, executar o comando: go run main.go

### Exemplos

Essa API utiliza a porta padrão 8080, abaixo seguem exemplos para utilização da api

> Criar planeta:

Method: POST  
url: http://localhost:8080/  
json Exemplo:  
{  
    "nome": "Yavin IV",  
    "clima": "temperate, tropical",  
    "terreno": "jungle, rainforests"  
}

> Listar planetas:

Method: GET  
url: http://localhost:8080/

> Buscar planeta por ID:

Method: GET  
url: http://localhost:8080/{id do planeta}  
url Exemplo: http://localhost:8080/5c0167ea03dea7318925137c

> Buscar planeta por Nome:

Method: POST  
url: http://localhost:8080/filter  
json Exemplo:  
{  
  "nome": "Yavin IV"  
}

> Excluir planeta:

Method: DELETE  
url: http://localhost:8080/{id do planeta}  
url Exemplo: http://localhost:8080/5c0167ea03dea7318925137c  

## Desenvolvido por

* **Elson Almeida** - [Kotayz](https://github.com/kotayz)
