/*Дано трехзначное число. Найдите сумму его цифр. 

Формат входных данных
На вход дается трехзначное число.

Формат выходных данных
Выведите одно целое число - сумму цифр введенного числа.*/

package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(a%10 + a/10%10 + a/100)
}
