package main

import (
	"manage_user/config"
	m "manage_user/middlewares"
	"manage_user/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8080"))
}
