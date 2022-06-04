package main

import (
	"ginblog/model"
	"ginblog/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
