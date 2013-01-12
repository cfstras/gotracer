package trace

import (
	"vec"
	"image"
	"image/color"
)

type Canvas struct {
	Size  vec.V2i
	Model color.Model

	Pixels    []vec.C3d
	Exposures int
}

func NewCanvas(size vec.V2i) *Canvas {
	c := Canvas{Size: size,
		Model:     color.RGBAModel,
		Pixels:    make([]vec.C3d, size.X*size.Y),
		Exposures: 0}
	return &c
}

func (c Canvas) ColorModel() color.Model {
	return c.Model
}

func (c Canvas) Bounds() image.Rectangle {
	return image.Rect(0, 0, c.Size.X, c.Size.Y)
}

func (c Canvas) At(x, y int) color.Color {
	pixel := c.Pixels[x+y*c.Size.X]
	pixel.DivTo(float64(c.Exposures))
	pixel.MulTo(255)
	return color.RGBA{uint8(pixel.R), uint8(pixel.G), uint8(pixel.B), uint8(255)}
}
