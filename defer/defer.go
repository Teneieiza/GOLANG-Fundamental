package deferexam

import "fmt"

func loop() {
	for i := 0; i < 5; i++ {
		fmt.Println("i not defer =",i)
	}
}

func deferLoop() {
	for i := range 5 {
		defer fmt.Println("i with defer =",i)
	}
}

func sum(a int, b int) int {
	result := a + b
	fmt.Println("result =", result)
	return result
}
//การใช้ defer เหมือนทำงานแบบ stack คือ ทำงานจากล่างขึ้นบน
// ซึ่งจะทำงานเมื่อ function ทำงานเสร็จ
// โดย code ไหนขึ้นก่อนและใช้ defer ก็จะอยู่ล่างสุดและค่อยๆทำงานขึ้นไป
func Defer() {

	fmt.Println("Wellcome to Defer")

	defer fmt.Println("End of Defer")

	loop()
	deferLoop()

	fmt.Println("result not defer")
	sum(2, 3)
	sum(2, 2)
	sum(2, 1)
	sum(1, 1)
	sum(1, 0)

	fmt.Println("result with defer")
	defer sum(2, 3)
	defer sum(2, 2)
	defer sum(2, 1)
	defer sum(1, 1)
	defer sum(1, 0)

	
}