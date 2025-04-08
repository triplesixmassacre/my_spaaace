// Давайте используем ваши знания структур, методов и интерфейсов на практике и реализуем объект,
// удовлетворяющий интерфейсу fmt.Stringer. Назовем его "Батарейка".

// Во-первых, вы должны объявить новый тип, удовлетворяющий интерфейсу fmt.Stringer.

// Ваш тип должен предусматривать, что на печати он будет выглядеть так: [      XXXX]: где пробелы - "опустошенная" емкость батареи,
// а X - "заряженная".

// Во-вторых, на стандартный ввод вы получаете строку, состоящую ровно из 10 цифр: 0 или 1 (порядок 0/1 случайный).
// Ваша задача считать эту строку любым возможным способом и создать на основе этой строки объект объявленного вами на первом этапе типа: надеюсь, вы понимаете, что строка символизирует емкость батарейки: 0 - это "опустошенная" часть, а 1 - "заряженная".

// В-третьих, созданный вами объект должен называться batteryForTest (использование этого имени обязательно).

// В вашем распоряжении фактически весь файл, НО завершающая фигурная скобка функции main() вам не видна, но она присутствует.
// Перед этой скобкой присутствует функция (которая вам тоже не видна), принимающая в качестве аргумента
// объект типа fmt.Stringer - batteryForTest, и направляющая его на стандартный вывод, поэтому вам не требуется
// выводить что-то на печать самостоятельно.

// Удачи!

// Sample Input:

// 1000010011

// Sample Output:

// [      XXXX]

package main

import (
	"fmt" // пакет используется для проверки ответа, не удаляйте его
	"strings"
)

type Battery string

func (b Battery) String() string {
	X := strings.Count(string(b), "1")
	space := 10 - X

	spaceStr := strings.Repeat(" ", space)
	XStr := strings.Repeat("X", X)
	return fmt.Sprintf("[%s%s]", spaceStr, XStr)
}

func () {

	var input string
	fmt.Scan(&input)

	if len(input) != 10 {
		fmt.Println("Error: look at your input(count)")
		return
	}

	for _, zn := range input {
		if zn != '0' && zn != '1' {
			fmt.Println("Error: there's imposter")
			return
		}
	}

	// batteryForTest := Battery(input) - какая-то функция для проверки
	fmt.Println(Battery(input))

}

// my code's output:

// PS D:\golang> go run interface_begin.go
// 10211110000
// Error: look at your input(count)

// PS D:\golang> go run interface_begin.go
// 1020301020
// Error: there's imposter

// PS D:\golang> go run interface_begin.go
// 1010010010
// [      XXXX]
