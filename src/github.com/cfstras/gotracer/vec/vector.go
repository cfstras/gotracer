package vec

import (

)

type V2i struct {
	X,Y int
}

type C3d struct {
	R,G,B float64
}

type V3d struct {
	X,Y,Z float64
}

type V2d struct {
	X,Y float64
}

func I(x, y int) V2i {
	return V2i{x,y}
}

func D(x,y,z float64) V3d {
	return V3d{x,y,z}
}

func D2(x,y float64) V2d {
	return V2d{x,y}
}

func C(r,g,b float64) C3d {
	return C3d{r,g,b}
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

func (v *V3d) MulTo(m float64){
	v.X *= m
	v.Y *= m
	v.Z *= m
}

func (v V3d) Mul(m float64) V3d {
	return D(v.X*m, v.Y*m, v.Z*m)
}

func (v V3d) Add(u V3d) V3d {
	return D(v.X+u.X, v.Y+u.Y, v.Z+u.Z)
}

func (v V3d) Sub(u V3d) V3d {
	return D(v.X-u.X, v.Y-u.Y, v.Z-u.Z)
}

func (v V3d) Cross(u V3d) V3d {
	return D(
	u.X*v.Z - u.Z*v.Y,
	u.Z*v.X - u.X*v.Z,
	u.X*v.Y - u.Y*v.X)
}

func (v V3d) Dot(u V3d) float64 {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
}

