// На стандартный ввод подается строковое представление даты и времени в следующем формате:

// 1986-04-16T05:20:00+06:00
// Ваша задача конвертировать эту строку в Time, а затем вывести в формате UnixDate:

// Wed Apr 16 05:20:00 +0600 1986

package main

import (
	"fmt"
	"time"
)

func () {
	var n string
	fmt.Scan(&n)
	inputTime, err := time.Parse("2006-01-02T15:04:05Z07:00", n)
	if err != nil {
		fmt.Println("не парься")
		return
	}
	resultTime := inputTime.Format(time.UnixDate)
	fmt.Println(resultTime)
}
