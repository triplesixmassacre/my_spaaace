//https://stepik.org/lesson/359395/step/8?unit=343626
package main

import (
	"fmt"
	"time"
)

func () {
	var min, sec int
	_, err := fmt.Scanf("%d мин. %d сек.", &min, &sec)
	if err != nil {
		fmt.Println("Ошибка при чтении данных")
		return
	}

	duration := time.Minute*time.Duration(min) + time.Second*time.Duration(sec)

	unixTime := int64(1589570165)

	baseTime := time.Unix(unixTime, 0).UTC()

	resultTime := baseTime.Add(duration)

	fmt.Println(resultTime.Format(time.UnixDate))
}
