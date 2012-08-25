package main

import (
	"fmt"
)

func main() {
	var a [4]int

/*
	var b [4]byte

	if &a == nil { 
		fmt.Println("a is nil")
	}
	fmt.Println("aの値:",a)
	
	for i:=0; i<len(a); i++ {
		fmt.Printf("aのアドレス:%p\n",&a[i])
	}
	
	for i:=0; i<len(b); i++ {
		fmt.Printf("bのアドレス:%p\n",&b[i])
	}
	
	c := 1
	//var c int =1 
	fmt.Printf("cの型:%T,%p\n",c, &c)
 */	
	a[0]=1
	i := a[0]
	fmt.Printf("a[0]のアドレス:%p\n", &a[0])
	fmt.Printf("a[1]のアドレス:%p\n", &a[1])
	fmt.Printf("a[2]のアドレス:%p\n", &a[2])
	fmt.Printf("a[3]のアドレス:%p\n", &a[3])
	fmt.Printf("i   のアドレス:%p\n", &i)

	i = 2
	fmt.Println("i=",i)
	fmt.Println("a[0]=",a[0])

/*
	s := make([]byte,5)
	if s == nil { 
		fmt.Println("s is nil")
	}
*/
	
}
