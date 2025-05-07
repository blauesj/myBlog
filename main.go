package main

import (
	"myBlog/models"
	"myBlog/router"
)

func main() {

	r := router.App()
	models.Init()
	r.Run(":8080")
}
