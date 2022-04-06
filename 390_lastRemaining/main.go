package main

func main() {

}

func lastRemaining(n int) int {

	result := 0

	forword := true
	toggle := false

	low := 1
	high := n
	slice :=[]bool
	m :=make(map[int]bool)
	for  i:=1; i<=n;i++{
		m[i]= true



		if forword {
			i = i+2
		}

	}

	return result
}
