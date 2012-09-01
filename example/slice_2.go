package main

import "fmt"

func main() {
	letters := []string{"a", "b", "c", "d"}
	fmt.Println(letters)
	fmt.Println("---------------------------------------")

	var s []byte
	s = make([]byte, 5, 5)
	fmt.Printf("sの型=%T\n", s)
	fmt.Printf("s=%v\n", s)
	fmt.Println("---------------------------------------")

	s2 := make([]byte, 5)
	fmt.Printf("s2の型=%T\n", s2)
	fmt.Printf("s2=%v\n", s2)
	fmt.Printf("len(s2)=%d, cap(s2)=%d\n", len(s2), cap(s2))
	fmt.Println("---------------------------------------")

	var s3 []byte
	if s3 == nil {
		fmt.Printf("s3 is nil\n")
	}
	if len(s3) == 0 && cap(s3) == 0 {
		fmt.Printf("len and cap of nil slice retrun 0\n")
	}
	fmt.Println("---------------------------------------")

	s4 := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	for i := 1; i < 4; i++ {
		fmt.Printf("s4[%d]のアドレス：%p\n", i, &s4[i])
	}
	s4_re := s4[1:4]

	fmt.Println(string(s4_re))                                        //内容を確認
	fmt.Printf("%T\n", s4_re)                                         //型を確認
	fmt.Println("len(s4_re)=", len(s4_re), "cap(s4_re)=", cap(s4_re)) //len(s4_re)==3, cap(s4_re)==5
	for i := 0; i < len(s4_re); i++ {
		fmt.Printf("s4_re[%d]のアドレス：%p\n", i, &s4_re[i])
	}

	fmt.Println("---------------------------------------")
	fmt.Println(string(s4[:]))
	fmt.Println(string(s4))

	fmt.Println(string(s4[:2]))
	fmt.Println(string(s4[2:]))

	fmt.Println("---------------------------------------")
	s5 := make([]int, 3, 6)
	fmt.Println(s5)
	s5 = s5[:cap(s)]
	fmt.Println(s5)

}
