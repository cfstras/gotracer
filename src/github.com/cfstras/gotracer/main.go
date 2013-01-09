package main

import(
	"fmt"
	"github.com/cfstras/gotracer/canvas"
	"github.com/cfstras/gotracer/vec"
	"os"
	"image/png"
)

func main() {
	fmt.Println("Tracing image.");
	
	test()
	
	fmt.Println("Exiting.");
}

func test() {
	canv := canvas.New(vec.I(20,30))
	save(canv, "test.png")
}

func save(canv canvas.Canvas, file string) {
	fo, err := os.Create(file);
	if err != nil {
		panic(err)
	}
	defer fo.Close()
	
	png.Encode(fo, canv)
}