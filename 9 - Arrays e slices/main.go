package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Arrays
	var a [5]int
	a[2] = 7
	println(a[2])

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	//os três pontos fixam o tamanho do array de acordo com a quantidade de elementos passados entre chaves
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// Slices
	s := []int{2, 3, 5, 7, 11, 13}
	s = append(s, 17)
	println(s[2])
	println(s[len(s)-1])

	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(array3))

	slice2 := array3[1:3]
	fmt.Println(slice2)

	array3[1] = 5000
	fmt.Println(slice2)
}
