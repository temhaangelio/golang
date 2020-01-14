package main

import (
	"fmt"
)

func main() {

	/*
		GO'yu yazan abilerimiz diğer döngülerin gereksizliğini sonunda
		fark etmişler ve sadece for döngüsünü bırakmışlar.
	*/

	//neredeyse tüm dillerde olan standart for döngü yapısı
	var count int = 10
	for i := 0; i < count; i++ {
		fmt.Println("Hello", i)
	}

	//for yapısı biraz daha kısaltabiliyoruz
	x := 0
	for x < count {
		fmt.Println("Mello", x)
		x += 1
	}

	//break ifadesi ile döngüsü gerektiğinde kırabiliyoruz
	y := 0
	for y < count {
		fmt.Println("Cello", y)
		if y == 5 {
			break
		}
		y += 1
	}
}
