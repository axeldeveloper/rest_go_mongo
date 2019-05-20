package main

import (
	"rest.gorn.mongo/api/app"
)

func main() {
	//config := config.GetConfig()

	//config *config.Config

	app := &app.App{}

	//fmt.Println("Database Name:", config.DB)
	app.Initialize()
	app.Run(":3000")
}
