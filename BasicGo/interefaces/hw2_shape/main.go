package main

import "fmt"

type triangle struct{
	base float64
	height float64
}
type square struct{
	sideLength float64
}

type shape interface{
	getArea() float64
}

func main(){
	tr := triangle{
		base:6.0,
		height: 2.0,
	}
	sq := square{
		sideLength: 10.0,
	}

	printArea(tr)
	printArea(sq)
}

func printArea(s shape){
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}