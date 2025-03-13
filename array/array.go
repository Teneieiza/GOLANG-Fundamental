package array

import "fmt"

func Array() {

	//การประกาศตัวแปรแบบ Array
	var PetName [5]string
	//การประกาศ array เราจำเป็นที่จะต้องประกาศข้อมูลให้มันด้วยผ่าน index ด้วย
	//ซึ่งถ้าเราไม่ทำการประกาศข้อมูลให้มัน จะเป็นค่าว่าง [		 ]
	//ตัวอย่างการประกาศ array
	PetName[0] = "Tigger"
	PetName[1] = "Toufu"
	PetName[2] = "Chinchan"
	PetName[3] = "LittleDragon"
	PetName[4] = "Milky"

	//ค่าที่จะออกมาจะมีค่าเนื่องจากเราได้ประกาศข้อมูลให้มันแล้ว
	//ตัวอย่างการเข้าถึงข้อมูลใน array
	fmt.Println(PetName)

	//การเข้าถึงข้อมูลใน array
	//โดยการใช้ index ของ array นั้นๆ ในการเข้าถึงข้อมูลใน array นั้นๆ จะเริ่มจาก 0
	fmt.Println(PetName[0])

	//ซึ่ง การประกาศ array string จะต่างกับ array ของ int หรือ float ตรงที่เราจะต้องกำหนดขนาดของ array ไว้ก่อน
	//โดย int หรือ float จะเป็น 0 แต่ string จะเป็นค่าว่าง
	var ageOfPet [5]int
	//ซึ่งถ้าเราไม่ทำการประกาศข้อมูลให้มัน จะเป็นค่า 0 แบบนี้ [0 0 0 0 0]
	//ตัวอย่างการประกาศ array
	ageOfPet[0] = 2
	ageOfPet[1] = 3
	ageOfPet[2] = 4
	ageOfPet[3] = 3
	ageOfPet[4] = 2

	//ค่าที่จะออกมาจะมีค่าเนื่องจากเราได้ประกาศข้อมูลให้มันแล้ว
	//ตัวอย่างการเข้าถึงข้อมูลใน array
	fmt.Println(ageOfPet)

	//การประกาศ array แบบ short variable declaration
	//ตัวอย่างการประกาศ array
	priceOfPet := [5]float64{1000, 2000, 3000, 4000, 5000}
	//ค่าที่จะออกมาจะมีค่าเนื่องจากเราได้ประกาศข้อมูลให้มันแล้ว

	//ตัวอย่างการเข้าถึงข้อมูลใน array
	fmt.Println(priceOfPet)
}
