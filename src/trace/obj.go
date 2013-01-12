package trace

import (
	"os"
	"bufio"
	"fmt"
	"io"
	"vec"
	"strings"
	"strconv"
)

const debug = true

type parser struct {
	scene *Scene
	verts map[int]vec.V3d
	numVerts int
	currentObj *Obj
}

func Parse(file string, scene *Scene) {
	p := &parser{scene: scene, verts: make(map[int]vec.V3d)}
	
	fs, err := os.Open(file);
	if err != nil {
		fmt.Println("Could not open file",file,err)
		return
	}
	defer fs.Close()
	
	in := bufio.NewReader(fs)
	for num := 0;; {
		line, err := in.ReadString('\n')
		num++
		if(err == io.EOF) {
			break
		}
		if(err != nil) {
			panic(err)
		}
		line = strings.TrimSpace(line,)
		fmt.Printf("parsing '%s'\n",line)
		p.parseLine(line, num)
	}
	p.closeObj()
}

func (p *parser) parseLine(line string, num int) {
	parts := strings.Split(line, " ")
	if(len(parts)==0) {
		return;
	}
	tok := parts[0]
	switch {
	case tok == "mtllib":
		//let's just ignore this one
	case tok == "o":
		//a new object
		p.addObj(parts[1:])
	case tok == "v":
		p.addV(parts[1:])
	case tok == "usemtl":
		//ignore this for now
	case tok  == "s":
		p.setShading(parts[1:])
	case tok == "f":
		p.addFace(parts[1:])
	case len(tok)>0 && tok[0] == '#':
		//ignore comments
	default:
		fmt.Printf("Line %d: can't parse '%s'\n",num,line)
	}
}

func (p *parser) addObj(args []string) {
	p.closeObj()
	p.currentObj = &Obj{Tris: make([]Tri,0,3), Color: vec.Red()}
	if(len(args)>0) {
		p.currentObj.Name = fmt.Sprint(args)
	}
}

func (p *parser) closeObj() {
	if p.currentObj != nil {
		p.scene.Objs = append(p.scene.Objs,p.currentObj)
	}
	if debug {
		fmt.Println("Obj",p.currentObj)
	}
	p.currentObj = nil
}

func (p *parser) addV(args []string) {
	if(len(args) != 3) {
		fmt.Print("invalid vertex",args);
		return
	}
	x, _ := strconv.ParseFloat(args[0],64)
	y, _ := strconv.ParseFloat(args[1],64)
	z, _ := strconv.ParseFloat(args[2],64)
	//TODO check errors
	v := vec.D(x,y,z)
	p.numVerts++
	p.verts[p.numVerts] = v
	
	if debug {
		fmt.Println("vertex",p.numVerts,v)
	}
}

func (p *parser) addFace(args []string) {
	switch {
	case len(args) == 0:
		fmt.Println("no arguments for face given")
	case len(args) == 3:
		v1, _ := strconv.Atoi(args[0])
		v2, _ := strconv.Atoi(args[1])
		v3, _ := strconv.Atoi(args[2])
		//TODO check errors
		tri := Tri{p.verts[v1],p.verts[v2],p.verts[v3]}
		p.currentObj.Tris = append(p.currentObj.Tris,tri)
	case len(args) == 4:
		v1, _ := strconv.Atoi(args[0])
		v2, _ := strconv.Atoi(args[1])
		v3, _ := strconv.Atoi(args[2])
		v4, _ := strconv.Atoi(args[3])
		tri1 := Tri{p.verts[v1],p.verts[v2],p.verts[v3]}
		tri2 := Tri{p.verts[v1],p.verts[v3],p.verts[v4]}
		p.currentObj.Tris = append(p.currentObj.Tris,tri1,tri2)
	default:
		fmt.Println("no idea what to do with vertex",args)	
	}
}

func (p *parser) setShading(args []string) {
	if(len(args) == 0) {
		fmt.Println("no arguments for shading option given")
		return
	}
	switch {
	case args[0] == "off":
		p.currentObj.Smooth = false
	case args[0] == "1":
		p.currentObj.Smooth = true
	default:
		fmt.Println("unknown shading option",args)
	}
}
