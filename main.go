package main

import "fmt"

func main() {
	foo()
	a := []int{5, 1, 4, 8, 6, 3}
	max := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	fmt.Println("Max:", max)
}
