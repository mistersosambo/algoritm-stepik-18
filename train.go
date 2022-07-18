package main

import "fmt"

func main() {
	var a, b, z int
	r := -1
	fmt.Scan(&z)
	var m = []int{1, 2}
	a, b = 0, 1
	for j := 0; j < 100; j++ {
		a, b = b, a+b
		m = append(m, b)
	}
	for i := range m {
		if m[i] == z {
			r = i
		}
	}
	fmt.Println(r)
}
