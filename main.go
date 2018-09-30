package main

import (
	"github.com/gin-gonic/gin"
	gallery "github.com/gin-practice/weddin/gallery"
	uploader "github.com/gin-practice/weddin/uploader"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/api")

	gallery.GalleryRegister(v1.Group("/gallery"))
	uploader.UploaderRegister(v1.Group("/uploader"))

	r.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})

	// for test
	r.GET("/image", gallery.CurrentImage)

	r.Run(":12345")
}
