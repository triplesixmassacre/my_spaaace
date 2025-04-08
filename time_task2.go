//https://stepik.org/lesson/359395/step/4?unit=343626
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func () {
	reader := bufio.NewReader(os.Stdin)
	timeStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка прочтения")
		return
	}
	timeStr = strings.TrimSpace(timeStr)

	timeTime, err := time.Parse(time.DateTime, timeStr)
	if err != nil {
		fmt.Println("2006-01-02 15:04:05", err)
		return
	}
	if timeTime.Hour() >= 13 {
		timeTime = timeTime.AddDate(0, 0, 1)
	}
	fmt.Println(timeTime.Format(time.DateTime))
}
