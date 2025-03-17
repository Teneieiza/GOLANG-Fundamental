package channel

import (
	"fmt"
	"time"
)

func processel(c chan string, data string) {
	c <- data
}

func Channel() {
	//กรณีที่เราต้องการระบุพิ้นที่ของ channel ที่ชัดเจน เช่น จองพื้นที่ไว้ 3 ที่ เราก็จะใส่เลขไปเลย
	//แต่กรณีที่เราต้องการเพิ่ม channel ได้เรื่อยๆ โดยไม่จำกัดจำนวน ก็ให้ใส่เป็น make(chan string)
	// และใช้ go เพิ่มข้อมูลให้กับ channel ได้เรื่อยๆ
	ch := make(chan string, 3) //ถ้าจะต้องการให้ไม่ระบุ capacity ก็ให้ใส่ make(chan string) เฉยๆนะจ้ะ
	processel(ch, "Data 1")
	processel(ch, "Data 2")
	processel(ch, "Data 3")
	//go processel(ch, "Data 4")
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch)
	time.Sleep(2 * time.Second)

	
}
