package main

import "github.com/StefanAtanaskovic/imgtool/imgtool"

func main() {
	img := imagetool.ReadImgFromFile("img.jpg")

	newImg := imagetool.GrayScaleFilter(img)

	imagetool.SaveImg(newImg, "newImg.jpg")
}
