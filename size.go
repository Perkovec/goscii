package goscii

import (
	"image"
)

func (c *goSCIIConverter) getASCIIOutputSize(imgBounds image.Rectangle) (image.Point, image.Point) {
	imgWidth := imgBounds.Dx()
	imgHeight := imgBounds.Dy()

	imageAspectRatio := float64(imgWidth) / float64(imgHeight)
	asciiAspectRatio := float64(c.columns) / float64(c.rows)

	asciiSize := image.Point{}
	imageSize := image.Point{}

	switch c.fit {
	case FitWidth:
		asciiSize.X = c.columns
		asciiSize.Y = int(float64(c.columns) / imageAspectRatio * c.fontAspectRatio)

		imageSize.X = imgWidth
		imageSize.Y = imgHeight
	case FitHeight:
		asciiSize.X = int(float64(c.rows) * imageAspectRatio)
		asciiSize.Y = c.rows

		imageSize.X = imgWidth
		imageSize.Y = imgHeight
	case FitCover:
		asciiSize.X = c.columns
		asciiSize.Y = c.rows

		if imageAspectRatio > 1.0 {
			imageSize.X = imgWidth
			imageSize.Y = int(float64(imgWidth) / asciiAspectRatio)
		} else {
			imageSize.X = int(float64(imgHeight) * asciiAspectRatio)
			imageSize.Y = imgHeight
		}
	case FitContain:
		imageSize.X = imgWidth
		imageSize.Y = imgHeight
		if imageAspectRatio > asciiAspectRatio {
			asciiSize.X = c.columns
			asciiSize.Y = int(float64(c.columns) / imageAspectRatio * c.fontAspectRatio)
		} else {
			asciiSize.X = int(float64(c.rows) * imageAspectRatio / c.fontAspectRatio)
			asciiSize.Y = c.rows
		}

	case FitFill:
		asciiSize.X = c.columns
		asciiSize.Y = c.rows

		imageSize.X = imgWidth
		imageSize.Y = imgHeight
	}

	return asciiSize, imageSize
}
