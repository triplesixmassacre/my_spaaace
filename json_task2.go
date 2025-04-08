//https://stepik.org/lesson/353243/step/9?unit=337227

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type OKVED struct {
	GlobalID int `json:"global_id"`
}

func () {
	file, err := os.Open("data-20190514T0100.json")
	if err != nil {
		log.Fatal("Не открыл файл", err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Не прочитал файл", err)
	}
	defer file.Close()

	var okveds []OKVED
	err = json.Unmarshal(data, &okveds)
	if err != nil {
		log.Fatal("Чёт не анмаршал")
	}

	sum := 0
	for _, okved := range okveds {
		sum += okved.GlobalID
	}

	fmt.Println(sum)

}
