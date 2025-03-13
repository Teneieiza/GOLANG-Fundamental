package variable

import (
	"fmt"

)

func Variable() {
	//การประกาศตัวแปรจะประกาศด้วย var ตามด้วยชื่อตัวแปร และประเภทของตัวแปร และค่าของตัวแปร ตามลำดับ
	//ตัวอย่างการประกาศตัวแปร
	var name string = "Ten"
	fmt.Println(name)
	fmt.Println(name + "Nitipat") //หรือจะใช้งาน string ในรูปแบบนี้ก็ได้
	fmt.Printf("My nickname is %s and Firstname is Nitipat\n", name) //หรือจะใช้งาน string ในรูปแบบนี้ก็ได้


	//การประกาศตัวแปรแบบกำหนดค่าให้ตัวแปรโดยไม่ต้องระบุประเภทของตัวแปร
	//ตัวอย่างการประกาศตัวแปร
	var age = 25
	fmt.Printf("Age : %d Yearsold\n",age)

	//การประกาศตัวแปรแบบหลายตัวแปรในบรรทัดเดียวกัน
	//ตัวอย่างการประกาศตัวแปร
	var day, mounth, year int = 28, 04, 1999
	fmt.Printf("My Birthdate :%02d/%02d/%04d\n", day, mounth, year)
	

	//การประกาศตัวแปรแบบ short variable declaration
	//ตัวอย่างการประกาศตัวแปร
	gradeFloat := 2.75
	fmt.Printf("My grade : %f \n",gradeFloat)

	var numberOfFamily int = 5
	numberPet := 0.5
	//การเปลี่ยนประเภทของตัวแปร
	//ตัวอย่างการเปลี่ยนประเภทของตัวแปร
	//กรณีถ้าต้องการ + Int และ float เข้าด้วยกันจะไม่สามารถทำได้
	//ตัวอย่าง fmt.Println(numberOfFamily+numberPet)
	//ต้องแปลงประเภทของข้อมูลก่อน
	fmt.Println(float64(numberOfFamily) + numberPet);
}