package main

import (
	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/hackernews-concurrency/routes"
)

func main() {
	app := gin.Default()

	app.LoadHTMLGlob("./templates/*.tmpl")

	app.GET("/", routes.GetNews)

	app.Run(":8080")

}
