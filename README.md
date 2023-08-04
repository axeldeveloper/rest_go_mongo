# Projeto rest em golang e mongodb

![My Go logo](https://github.com/axeldeveloper/rest_go_mongo/blob/master/axel-pk-go.png?raw=true "Axel alexander")




### go help list 
    - <https://golang.org/>
    - <https://gobyexample.com/>
    - <https://github.com/gorilla/mux>


## projeto
- GOLANG
- MONGODB
- REST
- MVC

# Configuração - Debian 10 (em breve)

# Configuração - Window

## Workspace - espaço de trabalho 
Configure seu espaço de trabalho Go . Isso consiste em três pastas na raiz:
- bin/
- pkg/
- src/

## GOROOT
   
   A GOROOT é uma variável de ambiente que especifica o local da instalação do go e pode ser definido como C: \ Go

## GOPATH

    A GOPATH é uma variável de ambiente que especifica o local do seu espaço de trabalho. 

    Se a GOPATH não estiver definido, presume-se que esteja $HOME/go nos sistemas Unix e %USERPROFILE%\go no Windows.

    Crie a variável de ambiente GOPATH e faça referência ao caminho do espaço de trabalho Go. Para adicionar, clique em System, Advanced system settings, Environment Variables...e clique New...em System variables:

    Windows 10 (linha de comando)
    Abra um prompt de comando ( Win + re digite cmd) ou uma janela do PowerShell ( Win+ i).
    Digite setx GOPATH %USERPROFILE%\go. (Isso definirá o GOPATHseu [home folder]\go, como C:\Users\yourusername\go.)
    Feche a janela do comando ou do PowerShell. (A variável de ambiente está disponível apenas para novas janelas de comando ou PowerShell, não para a janela atual.)

    echo %GOPATH%


# Demostração do Hello World

```go

$ go version
//go version go1.12 windows/amd64


/* Fenced code */
package main
import "fmt"
func main() {
    fmt.Println("hello world")
}
$ go run hello-world.go
// nome do objeto 
$ go build hello-world.go
$ ls
// hello-world    hello-world.go
$ ./hello-world
// hello world
```

# Craindo projeto
go mod init api-test

# Instalando pacotes 

- go get -u github.com/globalsign/mgo  (Mongodb)
- go get -u github.com/gorilla/mux  (A powerful HTTP router and URL matcher for building Go web servers with)

# Estrutura do projeto

- raiz:
    - main.go
    - config/config.go
    - api/app/handler/common.go
    - api/app/handler/employess.go
    - api/app/handler/mgodao.go
    - api/app/model/model.go      (vamos mudar para Employes)   


# Routes - Rotas do Projeto 
    GET  => http://localhost:8000/employess
    POST => http://localhost:8000/employees


# Criando Model
```go
    type Employee struct {
        _id    bson.ObjectId `bson:"_id,omitempty"`
        Name   string        `json:"name"`
        City   string        `json:"city"`
        Age    int           `json:"age"`
        Status bool          `json:"status"`
    }
```


#  Rum server
> go run main.go

## Deployment
    Axel Alexander

## My web site

* [axel](https://axe-dev.herokuapp.com/) - my site

## Contributing



## Versioning



## Authors

* **Axel Alexander ** - *web site* - [contact and contracts](http://axel-dev.herokuapp.com/)

 See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details