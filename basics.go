package main

import (
	"regexp"
	"fmt"
)

type point struct {
	X,Y,Z float64
}

func (p *point)parse(m string){

	re2 := regexp.MustCompile("(?P<v>[0-9-.]+)")
	result := re2.FindAllString(m, -1)
	if len(result) == 3 {
		fmt.Sscanf(result[0]+" "+result[1]+" "+result[2], "%f %f %f", &p.X, &p.Y, &p.Z)
	}
}
