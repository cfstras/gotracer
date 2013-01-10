package trace

import (
	"github.com/cfstras/gotracer/vec"
	"math"
)

const recursion = 5;

type Scene struct {
	Canvas *Canvas
	Res vec.V2i
	Objs []Obj
	
	ViewPlane vec.V2d
	CameraPos, ViewDir, RightDir, UpDir vec.V3d
	Fov float64
	
	Bg vec.C3d
}

func NewScene(canvas *Canvas) Scene {
	s := Scene{Canvas: canvas}
	s.Res = canvas.Size
	s.Objs = make([]Obj,2)
	s.ViewDir = vec.D(0.0, 0.0, 1.0)
	s.RightDir = vec.D(1.0, 0.0, 0.0)
	s.UpDir = vec.D(0.0, 1.0, 0.0)
	s.Fov = 90;
	
	height := 2.0 * math.Tan(s.Fov/2.0)
	s.ViewPlane = vec.D2( // viewplane at distance 1.0
		(height * float64(s.Res.X)) / float64(s.Res.Y),
		height)
	
	s.Bg = vec.C(0.0,0.0,0.5)
	
	return s
}

func (s *Scene) Trace() {
	for y := 0; y < s.Canvas.Size.Y; y++ {
		for x := 0; x < s.Canvas.Size.X; x++ {
			//create ray
			ray := s.Ray(vec.I(x,y))
			s.Canvas.Pixels[x + y * s.Res.X] = s.traceRay(ray, recursion)
		}
	}
}

func (s *Scene) Ray(pix vec.V2i) Ray {
	offsetX := s.ViewPlane.X * (float64(pix.X) / float64(s.Res.X) - 0.5 + 0.5 / float64(s.Res.X))
	offsetY := -s.ViewPlane.Y * (float64(pix.Y) / float64(s.Res.Y) - 0.5 + 0.5 / float64(s.Res.Y))
	
	planeOffX := s.RightDir.Mul(offsetX)
	planeOffY := s.UpDir.Mul(offsetY)
	
	dir := s.ViewDir.Add(planeOffX).Add(planeOffY);
	return Ray{s.CameraPos, dir}
}

func (s *Scene) traceRay (ray Ray, recurse int) vec.C3d {
	depth := math.Inf(1)
	var hitObj Obj;
	//var hitTri Tri;
	for _, obj := range s.Objs {
		for _, tri := range obj.Tris {
			if hit, length, _ := ray.Intersect(tri); hit && length <= depth {
				depth = length
				//TODO reflect
				hitObj = obj
			}
		}
	}
	
	if depth != math.Inf(1) {
		return hitObj.Color
	}
	return s.Bg
}
