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
	var a int
	fmt.Scan(&a)
	fmt.Println(a%10 + a/10%10 + a/100)
}
