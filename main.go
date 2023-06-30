package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	files := []string{
		"./templates/index.html",
		"./templates/global/base.html",
		"./templates/global/footer.html",
		"./templates/global/navbar.html",
		"./templates/global/js.html",
	}

	router.LoadHTMLFiles(files...)
	router.StaticFS("/static", http.Dir("static"))

	router.GET("/", index)
	router.Run(":80")

}
