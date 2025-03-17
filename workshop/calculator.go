package workshop

import (
	"bufio"   //bufio.NewReader(os.Stdin) ใช้สำหรับรับค่าจากคีย์บอร์ด
	"fmt"     //ใช้สำหรับพิมพ์ข้อความและรับค่า
	"os"      //ใช้ os.Exit(1) เพื่อปิดโปรแกรมหากเกิดข้อผิดพลาด
	"strconv" //ใช้ strconv.ParseFloat() แปลง string เป็น float64
	"strings" //ใช้ strings.TrimSpace() เพื่อลบช่องว่างและ \n ออกจาก input
)

var reader = bufio.NewReader(os.Stdin)

func getInput(promt string) float64 {
	fmt.Printf("%v", promt)                                        //แสดงข้อความ promt (เช่น "Enter first number: ")
	input, _ := reader.ReadString('\n')                            //รับค่าจากคีย์บอร์ดและเก็บไว้ในตัวแปร input
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64) //แปลงค่า input จาก string เป็น float64
	if err != nil {                                                //ถ้าเกิด error ในการแปลงค่าจะแสดงข้อความและจบโปรแกรม
		panic("Error: Input number only")
	}

	return value
}

func getOperator() string {
	fmt.Printf("Enter operator (+, -, *, /): ") //แสดงข้อความ "Enter operator (+, -, *, /): "
	operator, _ := reader.ReadString('\n')      //รับค่าจากคีย์บอร์ดและเก็บไว้ในตัวแปร operator
	operator = strings.TrimSpace(operator)      //ลบช่องว่างและ \n ออกจาก operator
	return operator
}

// ฟังก์ชันที่ใช้สำหรับการคำนวณ
func add(value1, value2 float64) float64 {
	return value1 + value2
}

func subtract(value1, value2 float64) float64 {
	return value1 - value2
}

func multiply(value1, value2 float64) float64 {
	return value1 * value2
}

func divide(value1, value2 float64) float64 {
	if value2 == 0 {
		fmt.Println("Error: Cannot divide by zero")
		os.Exit(1)
	}
	return value1 / value2
}

func Calculator() {
	// เรียกฟังก์ชัน getInput() สองครั้งเพื่อรับค่าจากคีย์บอร์ด
	value1 := getInput("Enter first number: ")
	value2 := getInput("Enter second number: ")

	// เรียกฟังก์ชัน getOperator() เพื่อรับค่า operator
	operator := getOperator()
	var result float64

	// ใช้ switch เพื่อตรวจสอบ operator และเรียกฟังก์ชันที่เกี่ยวข้อง
	switch operator {
	case "+":
		result = add(value1, value2)
	case "-":
		result = subtract(value1, value2)
	case "*":
		result = multiply(value1, value2)
	case "/":
		result = divide(value1, value2)
	default:
		fmt.Println("Invalid operator")
		os.Exit(1)
	}

	fmt.Printf("Result: %v %s %v = %v\n", value1, operator, value2, result)
}
