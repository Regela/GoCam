package main

import "math"

func round(num float64) int {
	return int(num + math.Copysign(0.9,num))
}

// p1 - p3: вершины треугольника, ptest: проверяемая точка.
// VEC - структура, содержащая поля X, Y, написанная нами.
// Можно вполне использовать POINT из <windows.h>
// Возвращается TRUE, если принадлежит, иначе - FALSE.

func isInTriangle(p1,p2,p3 point,x,y float64) bool{
	a := (p1.X - x) * (p2.Y - p1.Y) - (p2.X - p1.X) * (p1.Y - y)
	b := (p2.X - x) * (p3.Y - p2.Y) - (p3.X - p2.X) * (p2.Y - y)
	c := (p3.X - x) * (p1.Y - p3.Y) - (p1.X - p3.X) * (p3.Y - y)
	return (a >= 0 && b >= 0 && c >= 0) || (a <= 0 && b <= 0 && c <= 0)
}