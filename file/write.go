package file

import (
	"fmt"
	"os"
)

func Write() {
	//วิธีการเขียนไฟล์แบบใช้ os.WriteFile
	//ก็จะได้เป็นไฟล์ที่มีข้อมูลอยู่แล้ว
	writeTest := []byte("Hello, World! with Go Lang eieiza\nby ten eiei\nTest Success") //สร้างข้อมูลที่ต้องการเขียนลงไฟล์
	err := os.WriteFile("./file/Test.txt", writeTest, 0644)                //คำสั่งเขียนไฟล์ โดยกำหนด path,ข้อมูล และ permission
	// err คือตัวจัดการ error handling ถ้ามี error จะแสดงข้อความแจ้งเตือน
	if err != nil {
		fmt.Println("เกิดข้อผิดพลาด:", err) // แสดงข้อความแจ้งเตือน
		return
	}

	fmt.Println("เขียนไฟล์ Test สำเร็จ!") // หากไม่มี error ให้แสดงข้อความนี้

	//วิธีการเขียนไฟล์แบบใช้ os.Create
	//ก็จะได้เป็นไฟล์เปล่าขึ้นมา
	blankFile, err := os.Create("./file/blankfile.txt") //สร้างไฟล์เปล่า
	if err != nil {
		fmt.Println("เกิดข้อผิดพลาด:", err)
		return
	}

	fmt.Println("เขียนไฟล์เปล่าสำเร็จ!")

	defer blankFile.Close()

	writePet := []byte("Tigger, 2.5, idiot\nToufu, 1.5, cute\nchinchan, 0.7, smart")
	err = os.WriteFile("./file/petName.csv", writePet, 0644)
	
	if err != nil {
		fmt.Println("เกิดข้อผิดพลาด:", err)
		return
	}

	fmt.Println("เขียนไฟล์ petName สำเร็จ!")

}
