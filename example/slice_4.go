package main

import "fmt"

func main() {
	s := make([]byte, 5)
	fmt.Println("s=", s, "len(s)=", len(s), "cap(s)=", cap(s))

	t := make([]byte, len(s), (cap(s)+1)*2) // cap(s)==0 の場合は +1
	for i := range s {
		t[i] = s[i]
	}
	//copy(t,s)
	s = t
	fmt.Println("t=", t, "len(t)=", len(t), "cap(t)=", cap(t))
	fmt.Println("s=", s, "len(s)=", len(s), "cap(s)=", cap(s))

	fmt.Println("---------------------------------------")
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	var s2 = make([]int, 6)
	n1 := copy(s2, a[0:])
	fmt.Println("n1=", n1, ",s2=", s2)

	fmt.Println("---------------------------------------")
	s3 := make([]int, 1)
	fmt.Println("s3=", s3, "len(s3)=", len(s3), "cap(s3)=", cap(s3))

	// s3==[]int{0}
	s3 = append(s3, 1, 2, 3)
	fmt.Println("s3=", s3, "len(s3)=", len(s3), "cap(s3)=", cap(s3))
	fmt.Println("---------------------------------------")
	s4_a := []string{"John", "Paul"}
	s4_b := []string{"George", "Ringo", "Pete"}
	s4_a = append(s4_a, s4_b...)
	fmt.Println("s4_a=", s4_a, "len(s4_a)=", len(s4_a), "cap(s4_a)=", cap(s4_a))
	//a == []string{"John", "Paul", "George", "Ringo", "Pete"}

	fmt.Println("---------------------------------------")

	l := make([]int, 10)
	s5 := Filter(l, abc)
	fmt.Println("s5=", s5, "len(s5)=", len(s5), "cap(s5)=", cap(s5))

}

func abc(c int) bool {
	if c == 3 && c == 5 {
		return true
	}
	return false
}

// Filterは関数f()を満たすsの要素だけを
//　保持する新しいスライスを返します。
func Filter(s []int, fn func(int) bool) []int {
	var p []int // == nil
	for _, i := range s {
		if fn(i) {
			p = append(p, i)
		}
	}
	return p
}
