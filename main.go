package main

import (
	"flag"
	"fmt"
)

func main(){
	fileName := flag.String("f","","File Name")
	flag.Parse()
	var model stl
	model.read(*fileName)
	model.find_max()
	//side := model.to_side()
	//for i := 0; i < side.maxX; i++{
	//	for j := 0; j < side.maxY; j++{
	//		fmt.Printf("%f\t",side.points[i][j])
	//	}
	//	fmt.Printf("\n")
	//}
	f := facet{point{0,0,0},point{1,3,101}, point{3,1,101},point{0,0,0}}
	fmt.Printf("%f\n",f.find_Z(1,1))

}

