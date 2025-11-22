package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"parkingApp/internal/handlers"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		// CLI mode - original functionality
		runCLI()
	} else {
		// Web server mode
		runWebServer()
	}
}

func runCLI() {
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

func runWebServer() {
	handler := handlers.NewHandler()

	// Serve HTML file
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/index.html")
	})

	// API endpoint for commands
	http.HandleFunc("/execute", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "POST" {
			json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
			return
		}

		var request struct {
			Commands string `json:"commands"`
		}

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON"})
			return
		}

		// Process commands
		lines := strings.Split(request.Commands, "\n")
		var output strings.Builder

		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line != "" {
				result := handler.HandleCommand(line)
				if result != "" {
					output.WriteString(result)
					output.WriteString("\n")
				}
			}
		}

		json.NewEncoder(w).Encode(map[string]string{"output": output.String()})
	})

	fmt.Println("🚗 Parking App running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
