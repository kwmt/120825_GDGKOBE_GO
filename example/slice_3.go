package main

import "fmt"

func main(){
	s := make([]byte,5,5)
	s=s[2:4]
	fmt.Println("len(s)=",len(s),"cap(s)=",cap(s))
	fmt.Println("---------------------------------------")

	d := []byte{'r','o','a','d'}
	fmt.Println("d=",string(d))
	e := d[2:]
	fmt.Println("e=",string(e))
	e[1] = 'm'
	fmt.Println("d=",string(d))

	fmt.Println("---------------------------------------")

	s = s[:cap(s)]
	fmt.Println("len(s)=",len(s),"cap(s)=",cap(s))
}