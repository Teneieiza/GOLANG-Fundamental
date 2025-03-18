package main

import (
	"fmt"

	"github.com/Teneieiza/Golang-basic/file"
	// "github.com/Teneieiza/Golang-basic/defer"
	// "github.com/Teneieiza/Golang-basic/map"
	// "github.com/Teneieiza/Golang-basic/channel"
	// "github.com/Teneieiza/Golang-basic/interface"
	// "github.com/Teneieiza/Golang-basic/struct"
	// "github.com/Teneieiza/Golang-basic/workshop"
)

func main() {
	fmt.Println("Hello, World! with Go Lang eieiza")

	// workshop.Calculator()
	// structexam.Struct()  				// แสดงข้อมูลจาก struct.go
	// mapexam.Map()				 				// แสดงข้อมูลจาก map.go
	// interfaceexam.Interface() 	// แสดงข้อมูลจาก interface.go
	// channel.Routine()						// แสดงข้อมูลจาก goroutines.go
	// channel.Channel()						// แสดงข้อมูลจาก channel.go
	// deferexam.Defer()						// แสดงข้อมูลจาก defer.go
	file.Read()									// แสดงข้อมูลจาก read.go
	// file.Write()									// แสดงข้อมูลจาก write.go
}
