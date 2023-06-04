package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, n := range nums {
		if n%2 == 0 {
			fmt.Println(n, "es par.")
		} else {
			fmt.Println(n, "es impar.")
		}
	}
}
