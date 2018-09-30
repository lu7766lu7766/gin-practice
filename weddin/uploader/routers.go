package uploader

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func UploaderRegister(router *gin.RouterGroup) {

	router.GET("/", Index)
	router.POST("/", Upload)
}

func Index(c *gin.Context) {
	c.File("static/index_uploader.html")
}

func Upload(c *gin.Context) {
	// Accept for multi types of file
	file, header, err := c.Request.FormFile("uploader")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request you bad")
	}

	println("file name ", header.Filename)

	out, err := os.Create("upload_files/" + header.Filename)
	defer out.Close()
	if _, err := io.Copy(out, file); err != nil {
		c.String(http.StatusCreated, "upload fail")
	}
	c.String(http.StatusCreated, "upload successful heee")
}
