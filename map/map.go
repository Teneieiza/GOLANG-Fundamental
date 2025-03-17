package mapexam

import "fmt"

var product = make(map[string]float64)

func Map() {
	fmt.Println("Product deafult = ", product)

	//add
	product["milk"] = 20
	product["cookie"] = 10
	product["chocolate"] = 30
	product["rottenfood"] = 0
	fmt.Println("Product after Add = ", product)

	//delete
	delete(product, "rottenfood")
	fmt.Println("Product after Delete = ", product)

	//update
	product["milk"] = 25
	fmt.Println("Product after Update = ", product)

	//การเข้าถึงข้อมูลใน map
	value1 := product["milk"]
	fmt.Println("Product milk = ", value1)

	//การประกาศ map แบบ short variable declaration
	petName := map[string]string{
		"pigcat": "Tigger",
		"cat": "Toufu",
		"guiniepig": "Chinchan",
	}
	fmt.Println("Pet name: ", petName)
}