// https://stepik.org/lesson/353243/step/6?unit=337227
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// надо запомнить, что можно не объвлять полную структуру, если требуется работать только с одним полем этой структуры
type Group struct {
	Students []struct {
		Rating []int
	}
}
type averageRes struct {
	Average float64
}

func () {
	file, err := os.Open("json_zadacha.json") // открытие файла с проверкой на ошибку
	if err != nil {
		fmt.Println("Не открыл", err)
	}
	data, err := io.ReadAll(file) // чтение файла с проверкой
	if err != nil {
		log.Fatal("Ошибка чтения", err)
		return
	}

	// создаём переменную, в которую декодируем данные и делаем из json срез []byte
	var group Group
	err = json.Unmarshal(data, &group)
	if err != nil {
		log.Fatal("Чёт не анмаршал", err)
		return
	}

	allScores := 0
	allStudents := len(group.Students)

	// пробегаемся по студентам и добавляем количество оценок
	for _, student := range group.Students {
		allScores += len(student.Rating)
	}

	// вычисляем среднее значение оценок
	var Average float64
	if allStudents > 0 {
		Average = float64(allScores) / float64(allStudents)
	} else {
		Average = 0
	}

	result := averageRes{
		Average: Average,
	}

	output, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatal("аутпут пришёл капут", err)
	}
	fmt.Println(string(output))
}
