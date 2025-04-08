// Поэтапный поиск данных

// Данная задача в основном ориентирована на изучение типа bufio.Reader, поскольку этот тип позволяет считывать данные постепенно.

// В тестовом файле, который вы можете скачать из нашего репозитория на github.com, содержится длинный ряд чисел,
// разделенных символом ";". Требуется найти, на какой позиции находится число 0 и указать её в качестве ответа.
// Требуется вывести именно позицию числа, а не индекс (то-есть порядковый номер, нумерация с 1).

// Например:  12;234;6;0;78 :
// Правильный ответ будет порядковый номер числа: 4

package main

import (
	"fmt"
	"os"
	"strings"
)

func () {
	file, err := os.Open("task.data")
	if err != nil {
		fmt.Println("Can't open file", err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Can't get file info", err)
		return
	}
	fileSize := fileInfo.Size()

	content := make([]byte, fileSize)
	_, err = file.Read(content)
	if err != nil {
		fmt.Println("Can't read file", err)
		return
	}
	data := string(content)

	found := false
	position := 0

	numbers := strings.Split(data, ";")
	for _, num := range numbers {
		position++
		if num == "0" {
			found = true
			break
		}
	}

	if !found {
		fmt.Println("netu chisla")
	}
	fmt.Println(position)
}
