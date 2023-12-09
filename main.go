package main

import (
	"image"
	"image/jpeg"
	"log"
	"os"
	"github.com/StefanAtanaskovic/imgtool/filters"
)

func readImgFromFile(imgPath string) image.Image {
	reader, err := os.Open(imgPath)

	if err != nil {
		log.Fatal("fail to open image")
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)

	if err != nil {
		log.Fatal("failed to decode image")
	}

	return img
}

func saveImg(img image.Image, fileName string) {
	file, err := os.Create(fileName)

	if err != nil {
		log.Fatal("fail to create new file")
	}
	defer file.Close()

	err = jpeg.Encode(file, img, nil)
	if err != nil {
		log.Fatal("fail to save new image")
	}
}

func main() {
	img := readImgFromFile("img.jpg")

	newImg := filters.GrayScaleFilter(img)

	saveImg(newImg, "newImg.jpg")
}
