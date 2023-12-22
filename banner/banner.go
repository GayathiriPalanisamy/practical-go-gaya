package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	banner("Go", 6)
	fmt.Println(len("G☺"))
	banner("G☺", 6)
	s := "G☺"
	for i, r := range s{
		fmt.Println(i, r)
		if i == 0{
			fmt.Printf("%c of type %T\n",r,r)
			//rune (int32)
		}
	}
	b := s[0]
	fmt.Printf("%c of type %T\n",b,b)
	//byte (uint8)

	x, y := 1, "1"
	fmt.Printf("x=%v, y=%v\n", x, y)
	fmt.Printf("x=%#v, y=%#v\n", x, y) //USE to debug/log
	fmt.Printf("%20s!\n\n", s)
	fmt.Printf("gogo: %v\n" , isPalindrome("gogo"))
	fmt.Printf("g: %v\n" , isPalindrome("g"))
	fmt.Printf("gog: %v\n" , isPalindrome("gog"))
	fmt.Printf("g☺g: %v\n" , isPalindrome("g☺g"))
	
}

func isPalindrome(s string) bool{
	rs := (s) //to get slice of runes out of s
	fmt.Printf("count: %#v, value: %v \t", rs, s)
	for i := 0 ;i < len(rs)/2; i ++{
		if rs[i] != rs[len(rs)-i-1]{
			return false
		}
	}	
	return true
}
func banner(text string, width int){
	padding := (width - utf8.RuneCountInString(text)) / 2
	// padding := (width - len(text)) / 2 BUG: len is in bytes
	for i :=0; i< padding; i++{
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i :=0; i < width; i++{
		fmt.Print("-")
	}
} 