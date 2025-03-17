package interfaceexam

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64  //function ที่ใช้หาพื้นที่
	perim() float64 //function ที่ใช้หาเส้นรอบรูป
}

type rect struct {
	width, height float64 //กำหนดความกว้างและความสูงของรูปสี่เหลี่ยม
}

type circle struct {
	radius float64 //กำหนดรัศมีของวงกลม
}

// function สำหรับหาพื้นที่ของรูปสี่เหลี่ยม
// โดย (r rect) คือการระบุว่า function นี้เป็นของ struct rect
func (r rect) area() float64 {
	return r.width * r.height
}

// function สำหรับหาเส้นรอบรูปของรูปสี่เหลี่ยม
// โดย (r rect) คือการระบุว่า function นี้เป็นของ struct rect
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// function สำหรับหาพื้นที่ของวงกลม
// โดย (c circle) คือการระบุว่า function นี้เป็นของ struct circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// function สำหรับหาเส้นรอบรูปของวงกลม
// โดย (c circle) คือการระบุว่า function นี้เป็นของ struct circle
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// function สำหรับเรียกใช้งาน function ของ struct ที่เป็น interface
// โดย (g geometry) คือการระบุว่า function นี้รับค่าเป็น interface geometry
func measure(g geometry) {
	switch shape := g.(type) {
	case rect:
		fmt.Println("Rectangle")
		fmt.Printf("Width: %.2f, Height: %.2f\n", shape.width, shape.height)
	case circle:
		fmt.Println("Circle")
		fmt.Printf("Radius: %.2f\n", shape.radius)
	}

	fmt.Println("Area =", g.area())
	fmt.Println("Perimeter =", g.perim())
	fmt.Println("-----------------------------")
}

//function สำหรับเรียกใช้งาน
func Interface() {
	r := rect{width: 3, height: 4} //กำหนดความกว้างและความสูงของรูปสี่เหลี่ยม
	c := circle{radius: 5}         //กำหนดรัศมีของวงกลม

	measure(r)
	measure(c)
}
