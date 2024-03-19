package main

import (
	"go_cloud/model"
	"go_cloud/router"
)

func main() {
	model.InitDB()
	router.InitRouter()
}
