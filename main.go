package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for n >= 10 {
		sum := 0
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		n = sum
	}
	fmt.Println(n)
}
