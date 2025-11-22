package main

import (
	"bufio"
	"fmt"
	"os"
	"parkingApp/internal/handlers"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide input file")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	handler := handlers.NewHandler()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		output := handler.HandleCommand(line)
		if output != "" {
			fmt.Println(output)
		}
	}
}
