package main

import "fmt"

func main() {

	s := "abcd"
	t := "abcde"

	result := findTheDifference(s, t)
	fmt.Println(result)
}

//389. 找不同
func findTheDifference(s string, t string) byte {
	var result int32

	a := s + t
	for _, v := range a {
		result = result ^ v
	}

	return byte(result)
}
