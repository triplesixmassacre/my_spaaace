// https://stepik.org/lesson/359395/step/7?unit=343626
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"time"
)

const dateFormat = "02.01.2006 15:04:05"

func () {
	reader := bufio.NewReader(os.Stdin)
	timeStr, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println("Ошибка прочтения")
		return
	}
	timeStr = strings.TrimSpace(timeStr)
	timeParts := strings.Split(timeStr, ",")
	if len(timeParts) != 2 {
		fmt.Println("Никаретте")
		return
	}

	firstDate := strings.TrimSpace(timeParts[0])
	secondDate := strings.TrimSpace(timeParts[1])

	Date1, err := time.Parse(dateFormat, firstDate)
	if err != nil {
		fmt.Println("Ошибка парсинга")
		return
	}
	Date2, err := time.Parse(dateFormat, secondDate)
	if err != nil {
		fmt.Println("Ошибка парсинга")
		return
	}

	sub := (Date1.Sub(Date2))
	result := time.Duration(math.Abs(float64(sub)))
	fmt.Println(result)
}
