package trace

import (
	"vec"
	"math"
)

type Tri struct {
	V0, V1, V2 vec.V3d
}

type Ray struct {
	P, V vec.V3d
}

type Obj struct {
	Tris  []Tri
	Color vec.C3d
	Smooth bool
	Name string
}

func (r *Ray) Intersect(t *Tri) (hit bool, depth float64, intersect vec.V3d) {
	var v1, v2, norm vec.V3d
	var dot, dist float64
	hit = false
	depth = math.Inf(1)

	//vectors from triangle base to edges
	v1 = t.V1.Sub(t.V0)
	v2 = t.V2.Sub(t.V1)

	//triangle normal
	norm = v1.Cross(v2)

	//if dot is zero, line is parallel
	dot = norm.X*r.V.X + norm.Y*r.V.Y + norm.Z*r.V.Z

	if dot < 0 {
		//find intersect distance on ray
		dist = -(norm.X*(r.P.X-t.V0.X) + norm.Y*(r.P.Y-t.V0.Y) + norm.Z*(r.P.Z-t.V0.Z)) / dot
		dist = -( norm.Dot( r.P.Sub(t.V0) ) ) / dot

		//if negative, line started after the triangle
		if dist < 0 {
			return
		}

		intersect = r.P.Add(r.V.Mul(dist))

		if r.checkDir(t.V0, t.V1, intersect, norm) && r.checkDir(t.V1, t.V2, intersect, norm) && r.checkDir(t.V2, t.V0, intersect, norm) {
			hit = true
			depth = dist
		}
	}
	return
}

func (r *Ray) checkDir(v0, v1, v2, norm vec.V3d) bool {
	//answer
	if v1.Sub(v0).Cross(v2.Sub(v0)).Dot(norm) < 0 {
		return false
	}
	return true
}
