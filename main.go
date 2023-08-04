package main

import (
	"fmt"
	"rest_go_mongo/api/app"
)

func main() {
	//config := config.GetConfig()
	//config *config.Config
	app := &app.App{}
	fmt.Println("Iniciando Serviço :")
	app.Initialize()
	app.Run(":8183")
}
