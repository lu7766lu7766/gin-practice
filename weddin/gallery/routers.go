package gallery

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	imageupload "github.com/olahol/go-imageupload"
)

// for test of showing last uploaded image
var currentImage *imageupload.Image

func GalleryRegister(router *gin.RouterGroup) {
	router.GET("/", GalleryRetrieveAll)
	router.POST("/", GalleryUpload)

	router.GET("/:id", GalleryRetrieve)
}

func GalleryRetrieveAll(c *gin.Context) {
	msg := fmt.Sprint("Query for all.")
	c.JSON(http.StatusOK, gin.H{"msg": msg})
}

func GalleryUpload(c *gin.Context) {
	uploader := c.PostForm("uploader")
	println(fmt.Sprint("uploader = ", uploader))

	img, err := imageupload.Process(c.Request, "file")
	if err != nil {
		panic(err)
	}
	currentImage = img
	// TODO: save to GOOGLE DRIVE API
	// currentImage.Save(fmt.Sprintf("upload_images/%d.png", time.Now().Unix()))
	c.Redirect(http.StatusMovedPermanently, "/")
}

func GalleryRetrieve(c *gin.Context) {
	id := c.Param("id")
	msg := fmt.Sprint("Query for id = ", id)

	c.JSON(http.StatusOK, gin.H{"msg": msg})
}

func CurrentImage(c *gin.Context) {
	if currentImage == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	t, err := imageupload.ThumbnailJPEG(currentImage, 300, 300, 80)

	if err != nil {
		panic(err)
	}

	t.Write(c.Writer)
}
