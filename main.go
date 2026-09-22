package main

import (
	"fmt"
	"os"
)

func main() {
	// Проверяем количество аргументов (их должно быть строго 3)
	if len(os.Args) != 3 {
		fmt.Println("Error: Usage: go run . <input_file> <output_file>")
		os.Exit(1)
	}

	inputFile, outputFile := os.Args[1], os.Args[2]

	// 2. Чтение файла
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// 3. Трансформация текста
	text := Transform(data)

	// 4. Запись в файл
	err = os.WriteFile(outputFile, text, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		os.Exit(1)
	}
}
