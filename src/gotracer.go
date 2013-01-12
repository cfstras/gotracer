package main

import (
	"fmt"
	"trace"
	"vec"
	"image/png"
	"os"
	"time"
)

func main() {
	fmt.Println("Tracing image.")

	test()

	fmt.Println("Exiting.")
}

func test() {
	canv := trace.NewCanvas(vec.I(100, 100))
	scene := trace.NewScene(canv)

	/*tri := trace.Tri{vec.D(-5.0, 5.0, 5.0), vec.D(0.0, -5.0, 5.0), vec.D(5.0, 5.0, 5.0)}
	obj := trace.Obj{Color: vec.C(0.0, 1.0, 0.0)}
	obj.Tris = make([]trace.Tri, 0, 1)
	obj.Tris = append(obj.Tris, tri)
	scene.Objs = append(scene.Objs, obj)*/
	
	start := time.Now()
	trace.Parse("cubes.obj", scene);
	fmt.Println("Parsing took",time.Now().Sub(start))
	
	start = time.Now()
	scene.Trace()
	fmt.Println("Tracing took",time.Now().Sub(start))
	
	start = time.Now()
	save(canv, "test.png")
	fmt.Println("Exporting took",time.Now().Sub(start))
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

func Measure(f func(), name string) {
	start := time.Now()
	f()
	fmt.Println(name,"took",time.Now().Sub(start))
}