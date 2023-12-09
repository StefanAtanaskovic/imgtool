package main

import (
	"fmt"
	"image"

	"github.com/StefanAtanaskovic/imgtool/imgtool"
)

func createGray(done chan string, img image.Image) {
	grayImg := imagetool.GrayScaleFilter(img)
	imagetool.SaveImg(grayImg, "grey.jpg")

	done <- "done with gray"
}

func createSepia(done chan string, img image.Image) {
	sepiaImg := imagetool.SepiaFilter(img)
	imagetool.SaveImg(sepiaImg, "sepia.jpg")
	done <- "done with sepia"
}

func createInverted(done chan string, img image.Image) {
	invertedImg := imagetool.InvertColorsFilter(img)
	imagetool.SaveImg(invertedImg, "inverted.jpg")
	done <- "done with inverted"
}

func main() {
	img := imagetool.ReadImgFromFile("big.jpg")

	grayChan := make(chan string)
	sepiaChan := make(chan string)
	invertedChan := make(chan string)

	go createGray(grayChan, img)
	go createSepia(sepiaChan, img)
	go createInverted(invertedChan, img)


	for i := 0; i < 3; i++ {
		select {
		case grayMsg := <- grayChan:
			fmt.Println(grayMsg)
		case sepiaMsg := <- sepiaChan:
			fmt.Print(sepiaMsg)
		case invertedMsg := <- invertedChan:
			fmt.Println(invertedMsg)
		}
	}
}
