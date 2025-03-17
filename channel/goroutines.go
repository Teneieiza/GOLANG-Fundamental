package channel

import (
	"fmt"
	"time"
	// "time"
)

func f(from string) {
	for i := 0; i < 100; i++ {
		fmt.Println(from, ":", i)
	}
}

// ในหัวข้อนี้จะหมายถึง routine ของ go
// โดยการใช้ go นั้นจะทำให้การทำงานของ function นั้นๆ ทำงานพร้อมกัน
// โดยไม่ต้องรอให้ function ก่อนหน้าทำงานเสร็จก่อน

// เช่น array A มีข้อมูลจำนวณ 1000 ข้อมูล และ array B มีข้อมูลจำนวณ 1000 ข้อมูล

// ถ้าเราใช้ go ในการทำงาน function ของ array A และ array B จะทำงานพร้อมกัน
// และเสร็จพร้อมกันเช่นกัน

// แต่ถ้าเราไม่ใช้ go ในการทำงาน function ของ array A และ array B จะทำงานตามลำดับ
// ซึ่ง array A จะทำงานเสร็จก่อน array B จะทำงานต่อ
func Routine() {

	//จะสังเกตผ่าน console ได้ว่า การทำงานของ function จะทำงานgเป็นลำดับขั้น
	func(msg string) {
		fmt.Println(msg)
	}("This is a direct not use go")

	f("direct")

	f("direct2")

	f("direct3")

	time.Sleep(2 * time.Second)

	//จะสังเกตผ่าน console ได้ว่า การทำงานของ function จะทำงานพร้อมกัน
	func(msg string) {
		fmt.Println(msg)
	}("This is a goroutine use go")

	go f("goroutine")

	go f("goroutine2")

	go f("goroutine3")

	go f("goroutine4")

	go func(msg string) {
		fmt.Println(msg)
	}("Loading done")

	fmt.Scanln()
	fmt.Println("done")
}
