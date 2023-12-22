package main

import (
	"fmt"
	"sort"
)

func main() {
	var num []int //num is slic of integers
	fmt.Printf("%#v", num)
	var num2 = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("\n%#v -len(num): %d\n", num2, len(num2))
	num = num2[2:3]
	num = append(num, 3)

	fmt.Printf("%d\n", num)
	fmt.Printf("%d\n", num2) //num2 value changed as well

	fmt.Printf("%d-%d\n", len(num2), cap(num2))
	fmt.Printf("%d-%d\n", len(num), cap(num))

	// if num == num2 { //You cant compare slice to another slice
	if num == nil { //slice can be compared only to nil
		fmt.Print("equal")
	}
	// var number  []int
	// for i := 0; i< 1_0; i++{
	// 	//you dont let the slice panic with out-of-range error
	// 	number = appendInt(number,i)
	// }

	//excersise:
	fmt.Println(concat([]string{"a", "b"}, []string{"c", "d", "e"})) //{a, b, c, d, e}

	value := []float64{2, 11, 3, 4}
	fmt.Printf("%v \n %v", value, median(value))
}

func median(values []float64) float64 {
	v := make([]float64, len(values))
	copy(v,values)
	sort.Float64s(v) //changes the values in memory, you lose you original slice 
	i := len(v) / 2
	var median float64
	if i%2 == 1 {
		median = v[i]
	}
	median = (v[i-1] + v[i])/2
	return median
}

func concat(str1 []string, str2 []string) []string {
	//Restriction: no for-loop
	// str := append(str1, str2...) //naive approach

	str := make([]string, len(str1)+len(str2))
	copy(str, str1)
	copy(str[len(str1):], str2)
	return str
}

func appendInt(s []int, v int) []int {
	i := len(s)
	if len(s) < cap(s) {
		// we have enough space
		fmt.Printf("no allocation %d-%d\n", len(s), cap(s))
		s = s[:len(s)+1]
		fmt.Printf("s1: without shallow copy %d \n", s)
	} else {
		//when there is no enough space
		fmt.Printf("re-allocate %d-%d\n", len(s), cap(s))
		s2 := make([]int, 2*len(s)+1)
		copy(s2, s)
		fmt.Printf("s2: after shallow copy %d \n", s2)
		s = s2[:len(s)+1]
		fmt.Printf("s1: after shallow copy %d \n", s)
	}
	s[i] = v
	return s
}
