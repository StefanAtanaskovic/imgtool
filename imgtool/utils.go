package imagetool

import (
	"os"
	"image"
	"log"
	"image/jpeg"
) 

func Clamp(value, min, max float64) float64 {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}

func ReadImgFromFile(imgPath string) image.Image {
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

func SaveImg(img image.Image, fileName string) {
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
