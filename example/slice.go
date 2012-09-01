package slice

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	fmt.Println(string(b))
	b = digitRegexp.Find(b)
	fmt.Println(string(b))
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
