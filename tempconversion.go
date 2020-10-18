package main

import "fmt"

type rectangle struct {
	length  float32
	breadth float32
}

func (rect *rectangle) area() float32 {
	var area float32
	if rect.length > 0 && rect.breadth > 0 {
		area = rect.length * rect.breadth
	} else {
		fmt.Println("Enter valid dimensions of rectangle")
		area = -1
	}
	return area

}

func mainppp() {
	var l float32
	var b float32
	fmt.Println("Enter dimensions of rectangle :")
	fmt.Scanln(&l, &b)
	r := rectangle{length: l, breadth: b}
	fmt.Println("The area of the recangle is ", r.area())

}
