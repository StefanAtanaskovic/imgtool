package imagetool

import "image"
import "image/color"

func GrayScaleFilter(img image.Image) image.Image {
	bounds := img.Bounds()
	gray := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			gray.Set(x, y, img.At(x, y))
		}
	}

	return gray
}

func SepiaFilter(img image.Image) image.Image {
	bounds := img.Bounds()
	sepiaImg := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			initalColor := img.At(x, y)
			r, g, b, a := initalColor.RGBA()

			r >>= 8
			g >>= 8
			b >>= 8

			newR := 0.393*float64(r) + 0.769*float64(g) + 0.189*float64(b)
			newG := 0.349*float64(r) + 0.686*float64(g) + 0.168*float64(b)
			newB := 0.272*float64(r) + 0.534*float64(g) + 0.131*float64(b)

			newR = Clamp(newR, 0, 255)
			newG = Clamp(newG, 0, 255)
			newB = Clamp(newB, 0, 255)


			sepiaColor := color.RGBA{
				R: uint8(newR),
				G: uint8(newG),
				B: uint8(newB),
				A: uint8(a),
			}

			sepiaImg.Set(x, y, sepiaColor)
		}
	}

	return sepiaImg
}

func InvertColorsFilter(src image.Image) image.Image {
	bounds := src.Bounds()
	invertedImg := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			oldColor := src.At(x, y)
			r, g, b, a := oldColor.RGBA()

			newColor := color.RGBA{
				R: 255 - uint8(r>>8),
				G: 255 - uint8(g>>8),
				B: 255 - uint8(b>>8),
				A: uint8(a>>8),
			}

			invertedImg.Set(x, y, newColor)
		}
	}

	return invertedImg
}
