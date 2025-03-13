package slice

import "fmt"

func Slice() {

	//การประกาศตัวแปรแบบ Slice
	//การประกาศ Slice จะไม่ต้องกำหนดขนาดของ Slice ไว้ก่อน
	var petName []string
	petName = []string{"Tigger", "Toufu"}
	fmt.Printf("Pet name default: %s\n", petName)

	petName = append(petName, "Chinchan", "LittleDragon", "Milky")
	fmt.Printf("Pet name after append: %s\n", petName)

	//ลองใช้เป็นตัวเลขดูบ้าง
	var priceOfFoodPet []float64
	priceOfFoodPet = []float64{1000, 2000, 3000, 4000, 5000}
	fmt.Printf("price of Foodpet default: %f\n", priceOfFoodPet)

	priceOfFoodPet = append(priceOfFoodPet, 6000, 7000, 8000, 9000, 10000)
	fmt.Printf("price of Foodpet after append: %f\n", priceOfFoodPet)

	//การประกาศ Slice แบบ short variable declaration
	//การประกาศ Slice แบบ short variable declaration จะต้องกำหนดข้อมูลให้มันเสมอ
	// weightOfPet := []float64{2.3, 0.2, 0.4, 0.3, 0.2} ซึ่งเราสามารถประกาศแบบนี้ได้
	// แต่ถ้าเราประกาศแบบนี้ weightOfPet := []float64{} จะเกิด error เพราะเราไม่ได้กำหนดข้อมูลให้มัน
	weightOfPet := []float64{2.3, 0.2, 0.4, 0.3, 0.2}
	fmt.Printf("Weight of all Pet: %f\n", weightOfPet)

	//การเข้าถึงข้อมูลใน Slice แบบเลิอกเฉพาะ index ที่เราต้องการ โดย 5 คือเริ่มจาก index ที่ 5 และ 9 คือจนถึง index ที่ 9
	//โดยที่ index 9 จะไม่เอาเข้าไปด้วย
	//ตัวอย่างการเข้าถึงข้อมูลใน Slice
	slicePrice := priceOfFoodPet[5:9]
	fmt.Println("Pick price from index 5 to index 9", slicePrice)

	//โดยถ้าเราจะเลือกข้อมูลใหม่ให้ทำตามนี้
	//ซึ่งกรณีนี้จะเลือกข้อมูล 5 ตำแหน่งแรก
	slicePrice = priceOfFoodPet[:5]
	fmt.Println("Pick 5 first index from Price of food", slicePrice)

	// กรณีนี้จะเลือกข้อมูล 5 ตำแหน่งหลัง
	slicePrice = priceOfFoodPet[5:]
	fmt.Println("Pick 5 back index from Price of food", slicePrice)

	// กรณีที่เราแก้ไขข้อมูลจะส่งผลกระทบกับข้อมูลต้นฉบับด้วย
	priceOfFoodPet[0] = 6000
	fmt.Println("priceOfFoodPet after edit index[0]", priceOfFoodPet)

	//การดึงข้อมูลชิ้นเดียวจาก Slice
	//ตัวอย่างการดึงข้อมูลชิ้นแรกจาก Slice
	firstPrice := priceOfFoodPet[0]
	//และการดึงข้อมูลชิ้นสุดท้ายจาก Slice
	lastPrice := priceOfFoodPet[len(priceOfFoodPet)-1]

	fmt.Println("Pick first index: ", firstPrice)
	fmt.Println("Pick last index: ", lastPrice)

	//การสร้าง Slice โดยใช้ make จะต้องกำหนดขนาดของ Slice ไว้ก่อน
	// ในที่นี้เราจะสร้าง Slice ที่มีขนาด 5 โดยให้ค่าของ Slice เป็นค่าว่าง
	// และเราสามารถเพิ่มข้อมูลเข้าไปได้โดยการใช้ append
	// โดยการใช้ make จะทำให้เราสามารถเพิ่มข้อมูลเข้าไปได้
	// แต่ถ้าเรากำหนดขนาดของ Slice ไว้เท่ากับ 0 จะไม่สามารถเพิ่มข้อมูลเข้าไปได้
	// ตัวอย่างการสร้าง Slice โดยใช้ make
	// foodPrice := make([]float64, 5)
	// fmt.Println("Create slice by make before add info:",foodPrice)
	// // การเพิ่มข้อมูลเข้าไปใน Slice
	// foodPrice = append(foodPrice, 1000, 2000, 3000, 4000, 5000)
	// fmt.Println("Create slice by make after add info:",foodPrice)

	//การลบข้อมูลใน Slice
	// การลบข้อมูลใน Slice จะต้องใช้ append โดยการใช้ append และใช้ Slice ในการเลือกข้อมูลที่เราต้องการลบออก
	// ตัวอย่างการลบข้อมูลใน Slice
	foodPrice := []float64{1000, 2000, 3000, 4000, 5000}
	foodPrice = append(foodPrice[:2], foodPrice[4:]...) //โดยกรณีนี้คือการลบจ้อมูล index ที่ 2 และ 3 ออก
	fmt.Println("After Delete:",foodPrice)

}