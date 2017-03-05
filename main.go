package main

import "fmt"

func main() {
	fmt.Println("my nam feature branch first commit")
	fmt.Println("Hey we are going to find max in the below array")
	a := []int{5, 1, 4, 8, 6, 3}
	for _, v := range a {
		fmt.Println(v)
	}

	max := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	fmt.Println("Max:", max)
}
