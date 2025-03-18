package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Read() {
	//คำสั่งเปิดไฟล์
	recieveFile, err := os.Open("./file/petName.csv")//เข้าถึง path ของไฟล์
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(recieveFile) //scan ข้อมูลในไฟล์
	count := 1
	for scanner.Scan() { //จากนั้นนำมา loop ในแต่ละ line
		line := scanner.Text()
		// fmt.Printf("Line %d: %s\n", count, line) //กรณีที่ต้องการแสดงข้อมูลทั้งหมดในแต่ละบรรทัด

		//กรณีที่ต้องการแสดงข้อมูลใน column ที่ต้องการ
		item := strings.Split(line, ",")
		fmt.Println(item[0])
		count++
	}
}
