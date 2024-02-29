package main

import (
	"github.com/ccesarfp/shrine/internal/config/application"
)

func main() {
	app := application.New()
	app.S.SetupServer()
	app.Up()
}
