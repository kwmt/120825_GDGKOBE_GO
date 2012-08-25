package main

import "fmt"

func main(){

	//参考
	//http://go-book.appspot.com/more-functions.html#using-the-append-function-to-delete-an-element
	fmt.Println("---------------------------------------")
    slice := []int {1, 2, 3, 4, 5}


    fmt.Println("In the beginning...")
    fmt.Println("slice = ", slice, "len(slice)= ", len(slice), "cap(slice)= ", cap(slice))
	fmt.Printf("%p\n",slice)
    fmt.Println("Let's delete the first element")
    //slice = delete(2, slice)
	slice=append(slice[:2],slice[3:]...)
    fmt.Println("slice = ", slice, "len(slice)= ", len(slice), "cap(slice)= ", cap(slice))
	fmt.Printf("%p\n",slice)
    fmt.Println("Let's delete the last element")
    slice = delete(len(slice)-1, slice)
    fmt.Println("slice = ", slice, "len(slice)= ", len(slice), "cap(slice)= ", cap(slice))
	fmt.Printf("%p\n",slice)
    fmt.Println("Let's delete the 3rd element")
    slice = delete(2, slice)
    fmt.Println("slice = ", slice, "len(slice)= ", len(slice), "cap(slice)= ", cap(slice))
	fmt.Printf("%p\n",slice)

}

func delete(i int, slice []int) []int{
	slice = append(slice[:i], slice[i+1:]...)
	c := make([]int, len(slice))
	copy(c, slice)
    return c
	//return slice
}
