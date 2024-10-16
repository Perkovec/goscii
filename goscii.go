package goscii

import (
	"image"
)

type GOSCIIConverter interface {
	Convert(img image.Image) []byte
}

type goSCIIConverter struct {
	charset         Charset
	columns         int
	rows            int
	fit             Fit
	fontAspectRatio float64
}

/*
Create new converter
*/
func NewConverter(options ...GOSCIIConverterOptions) (*goSCIIConverter, error) {
	opts := defaultOptions

	if len(options) > 0 {
		opts.Merge(options[0])
	}

	return &goSCIIConverter{
		charset:         opts.Charset,
		columns:         opts.Columns,
		rows:            opts.Rows,
		fit:             opts.Fit,
		fontAspectRatio: opts.FontAspectRatio,
	}, nil
}

func (c *goSCIIConverter) normalizeCharIndex(i int) int {
	return int(float64(i) / 255.0 * float64(len(c.charset)))
}

/*
Convert image to ASCII art
*/
func (c *goSCIIConverter) Convert(img image.Image) []string {
	asciiSize, imageSize := c.getASCIIOutputSize(img.Bounds())

	deltaX := imageSize.X / asciiSize.X
	deltaY := imageSize.Y / asciiSize.Y

	artRows := make([]string, 0, asciiSize.Y)
	var artRow string
	for row := 0; row < asciiSize.Y; row++ {
		artRow = ""
		for col := 0; col < asciiSize.X; col++ {
			r, g, b, _ := img.At(col*deltaX, row*deltaY).RGBA()

			r = r >> 8
			g = g >> 8
			b = b >> 8

			gray := (float64(r) * 0.2126) + (float64(g) * 0.7152) + (float64(b) * 0.0722)

			artRow += string(c.charset[c.normalizeCharIndex(int(gray))])
		}

		artRows = append(artRows, artRow)
	}

	return artRows
}
