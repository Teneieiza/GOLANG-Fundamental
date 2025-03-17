package structexam

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	phone        string
}

func Struct() {

	//การกำหนดข้อมูลให้กับ struct
	employee1 := employee{
		employeeID:   "001",
		employeeName: "Ten",
		phone:        "0812345678",
	}
	fmt.Println("This is employee1: ", employee1)

	//การกำหนดขนาดของ struct แบบ array
	//ต้องใช้ index ในการกำหนดข้อมูลให้กับ struct
	employeeArray := [3]employee{}
	employeeArray[0] = employee{
		employeeID:   "001",
		employeeName: "Ten",
		phone:        "0812345678",
	}
	employeeArray[1] = employee{
		employeeID:   "002",
		employeeName: "Eleven",
		phone:        "0812345679",
	}
	employeeArray[2] = employee{
		employeeID:   "003",
		employeeName: "Twelve",
		phone:        "0812345680",
	}
	fmt.Println("This is employeeArray: ", employeeArray)

	//การกำหนดขนาดของ struct แบบ slice
	//ต้องใช้ append เนื่องจาก slice ไม่มีการกำหนดขนาด ต่อให้เข้าถึง index ที่0 ก็จะ error
	employeeSlice := []employee{}
	employeeSlice = append(employeeSlice, employee{
		employeeID:   "001",
		employeeName: "Ten",
		phone:        "0812345678",
	})
	employeeSlice = append(employeeSlice, employee{
		employeeID:   "002",
		employeeName: "Eleven",
		phone:        "0812345679",
	})
	fmt.Println("This is employeeSlice: ", employeeSlice)

}
