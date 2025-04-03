package main

import (
	server "github.com/juheth/Go-Clean-Arquitecture/src/infrastructure/server"
)

func main() {
	app := server.ProvidersStore{}
	app.Init()
	app.Up()

}
