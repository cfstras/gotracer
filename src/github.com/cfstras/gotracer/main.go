package main

import (
	"fmt"
	"github.com/cfstras/gotracer/trace"
	"github.com/cfstras/gotracer/vec"
	"image/png"
	"os"
)

func main() {
	fmt.Println("Tracing image.")

	test()

	fmt.Println("Exiting.")
}

func test() {
	canv := trace.NewCanvas(vec.I(100, 100))
	scene := trace.NewScene(canv)

	tri := trace.Tri{vec.D(-5.0, 5.0, 5.0), vec.D(0.0, -5.0, 5.0), vec.D(5.0, 5.0, 5.0)}
	obj := trace.Obj{Color: vec.C(0.0, 1.0, 0.0)}
	obj.Tris = make([]trace.Tri, 0, 1)
	obj.Tris = append(obj.Tris, tri)
	scene.Objs = append(scene.Objs, obj)

	scene.Trace()
	save(canv, "test.png")
}

func save(canv *trace.Canvas, file string) {
	fo, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	png.Encode(fo, canv)
	fo.Sync()
}
