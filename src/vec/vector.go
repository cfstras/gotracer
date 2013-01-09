package vec

import (

)

type V2i struct {
	X,Y int
}

type C3d struct {
	R,G,B float64
}

func I(x, y int) V2i {
	return V2i{x,y}
}

func (c C3d) Div(d float64) C3d{
	return C3d{c.R/d, c.G/d, c.B/d}
}

func (c *C3d) DivTo(d float64){
	c.R /= d
	c.G /= d
	c.B /= d
}

func (c *C3d) MulTo(m float64){
	c.R *= m
	c.G *= m
	c.B *= m
}