package main

import (
	"fmt"
)

func main() {

	/* デモ */
	//配列宣言
	var a [4]int

	// 0番目の要素に1を代入
	a[0] = 1

	//変数 i を宣言して a[0]を i に代入
	i := a[0]

	fmt.Printf("a[0]=%d\n", a[0])
	fmt.Printf("i=%d\n", i)

	/*----- 以下参考 -----*/
	if i == a[0] {
		fmt.Println("i == a[0]")
	} else {
		fmt.Println("i != a[0]")
	}

	/* デモ */
	for j := 0; j < len(a); j++ {
		if a[j] == 0 {
			fmt.Printf("a[%d] == %d\n", j, a[j])
		} else {
			fmt.Printf("a[%d] != %d\n", j, a[j])
		}
	}
	var p *int = &a[3]
	fmt.Printf("aのポインタ:%p\n", p)

	/* デモ */
	//b := [2]string{"Penn","Teller"}
	b := [...]string{"Penn", "Teller"}
	fmt.Println(len(b))

}
