package handler

import (
	"image"
	"image/jpeg"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
)

func ResizeImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(400, gin.H{
			"error": "No image uploaded",
		})
		return
	}

	processedImageChan := make(chan *os.File)

	go processImage(file, processedImageChan)

	processedImage := <-processedImageChan

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+processedImage.Name())
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(processedImage.Name())

	processedImage.Close()
}

func processImage(file *multipart.FileHeader, processedImageChan chan<- *os.File) {
	defer close(processedImageChan)

	fileStream, err := file.Open()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer fileStream.Close()

	img, _, err := image.Decode(fileStream)
	if err != nil {
		log.Fatal(err)
		return
	}

	resizedImg := resize.Resize(1000, 0, img, resize.Lanczos3)

	outFileName := "see_albania_resized_" + filepath.Base(file.Filename)
	outFile, err := os.Create(outFileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer outFile.Close()

	err = jpeg.Encode(outFile, resizedImg, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("Image resized and saved as: %s\n", outFileName)

	processedImage, err := os.Open(outFileName)
	if err != nil {
		log.Fatal(err)
		return
	}

	processedImageChan <- processedImage
}
