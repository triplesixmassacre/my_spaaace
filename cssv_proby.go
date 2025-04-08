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

	// Get the file size
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Can't get file info", err)
		return
	}
	fileSize := fileInfo.Size()

	// Read the entire file into a byte slice
	content := make([]byte, fileSize)
	_, err = file.Read(content)
	if err != nil {
		fmt.Println("Can't read file", err)
		return
	}

	// Convert the byte slice to a string
	data := string(content)

	// Split the string by the delimiter ";"
	numbers := strings.Split(data, ";")

	position := 0
	found := false

	for _, numStr := range numbers {
		position++
		if numStr == "0" {
			fmt.Println(position)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("0 not found")
	}
}
