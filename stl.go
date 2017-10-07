package main

import (
	"strings"
	"os"
	"log"
	"fmt"
)

type stl struct{
	facets []facet
	maxX,maxY,maxZ float64
}

type facet struct {
	normal,vertex1,vertex2,vertex3 point
}

func (model *stl)parse(m string){
	m = strings.Replace(m,"\n\r","\n",-1)
	lines := strings.Split(m,"\n")
	for i := 0; i < len(lines); i++ {
		if strings.Contains(lines[i],"facet") {
			var buff facet
			buff.normal.parse(lines[i])
			vertex := 1
			for ;!strings.Contains(lines[i],"endfacet");i++ {
				if strings.Contains(lines[i],"vertex") {
					switch (vertex){
					case 1:
						buff.vertex1.parse(lines[i])
						break
					case 2:
						buff.vertex2.parse(lines[i])
						break
					case 3:
						buff.vertex3.parse(lines[i])
						break
					}
					vertex++
				}
			}
			model.facets=append(model.facets, buff)
		}
	}
}

func (model *stl)read(filename string){
	file, err := os.Open(filename) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	m := ""
	fi,_ := os.Stat(filename)
	data := make([]byte,fi.Size())

	_, err = file.Read(data)
	m+=string(data)
	if err != nil {
		log.Fatal(err)
	}
	model.parse(m)
}

func (model *stl)find_max(){
	var maxX,maxY,maxZ,minX,minY,minZ float64
	minX = model.facets[0].vertex1.X
	maxX = minX

	minY = model.facets[0].vertex1.Y
	maxY = minY

	minZ = model.facets[0].vertex1.Z
	maxZ = minZ

	for _,f := range model.facets {
		if f.vertex1.X > maxX {
			maxX = f.vertex1.X
		}
		if f.vertex1.X < minX {
			minX = f.vertex1.X
		}
		if f.vertex1.Y > maxY {
			maxY = f.vertex1.Y
		}
		if f.vertex1.Y < minY {
			minY = f.vertex1.Y
		}
		if f.vertex1.Z > maxZ {
			maxZ = f.vertex1.Z
		}
		if f.vertex1.X < minX {
			minZ = f.vertex1.Z
		}

		if f.vertex2.X > maxX {
			maxX = f.vertex2.X
		}
		if f.vertex2.X < minX {
			minX = f.vertex2.X
		}
		if f.vertex2.Y > maxY {
			maxY = f.vertex2.Y
		}
		if f.vertex2.Y < minY {
			minY = f.vertex2.Y
		}
		if f.vertex2.Z > maxZ {
			maxZ = f.vertex2.Z
		}
		if f.vertex2.X < minX {
			minZ = f.vertex2.Z
		}

		if f.vertex3.X > maxX {
			maxX = f.vertex3.X
		}
		if f.vertex3.X < minX {
			minX = f.vertex3.X
		}
		if f.vertex3.Y > maxY {
			maxY = f.vertex3.Y
		}
		if f.vertex3.Y < minY {
			minY = f.vertex3.Y
		}
		if f.vertex3.Z > maxZ {
			maxZ = f.vertex3.Z
		}
		if f.vertex3.X < minX {
			minZ = f.vertex3.Z
		}
	}

	for i := range model.facets{
		if minX < 0 {
			model.facets[i].vertex1.X-=minX
			model.facets[i].vertex2.X-=minX
			model.facets[i].vertex3.X-=minX
		}
		if minY < 0 {
			model.facets[i].vertex1.Y-=minY
			model.facets[i].vertex2.Y-=minY
			model.facets[i].vertex3.Y-=minY
		}
		if minZ < 0 {
			model.facets[i].vertex1.Z-=minZ
			model.facets[i].vertex2.Z-=minZ
			model.facets[i].vertex3.Z-=minZ
		}
	}
	if minX < 0 {
		maxX-=minX
	}
	if minY < 0 {
		maxY-=minY
	}
	if minZ < 0 {
		maxZ-=minZ
	}

	model.maxZ = maxZ
	model.maxX = maxX
	model.maxY = maxY
}

func (model *stl) to_side()(side){
	var ret side
	maxX := round(model.maxX)
	maxY := round(model.maxY)
	//maxZ := round(model.maxZ)

	ret.points=make([][]float64,maxX)
	for i := 0 ; i < maxX; i++{
		ret.points[i]=make([]float64,maxY)
	}
	for i := 0 ; i < maxX; i++{
		for j := 0; j < maxY; j++ {
			ret.points[i][j] = model.find_Z(float64(i),float64(j))
		}
	}
	ret.maxX = maxX
	ret.maxY = maxY
	return ret
}

func (model *stl)find_Z(x,y float64) float64{
	maxZ := 0.0
	for _,f := range model.facets{
		if isInTriangle(f.vertex1,f.vertex2,f.vertex3,x,y) {
			z := f.find_Z(x,y)
			if z > maxZ {
				maxZ = z
			}
		}
	}
	return maxZ
}

func (f *facet)find_Z(x,y float64) (z float64){
	x1 := f.vertex1.X
	y1 := f.vertex1.Y
	z1 := f.vertex1.Z
	x2 := f.vertex2.X
	y2 := f.vertex2.Y
	z2 := f.vertex2.Z
	x3 := f.vertex3.X
	y3 := f.vertex3.Y
	z3 := f.vertex3.Z

	zk := x1*y2-x1*y3-x2*y1+x2*y3+x3*y1-x3*y2
	fmt.Printf("zk: %f\n",zk)
	if zk == 0.0 {
		z = 100
	} else {
		xk := 0-y1*z2+y1*z3-y2*z1+y2*z3+y3*z1-y3*z2
		fmt.Printf("xk: %f\n",xk)
		yk :=   x1*z2-x1*z3-x2*z1+x2*z3+x3*z1-x3*z2
		fmt.Printf("yk: %f\n",yk)

		d  := x1*y2*z3-x1*y3*z2-x2*y1*z3+x2*y3*z1+x3*y1*z2-x3*y2*z1
		fmt.Printf("d: %f\n",d)
		z = (x*xk + y*yk + d) / zk
	}
	return z
}

