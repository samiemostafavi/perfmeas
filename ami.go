package advmobileinfo

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func Run() {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":50500", nil))
}

func getStatus() ([]byte, error) {
	// Execute the command and capture the output
	output, err := exec.Command("status", "-v", "mobile", "1").Output()
	if err != nil {
		log.Fatal(err)
	}

	// Parse the output and create a map to store key-value pairs
	result := make(map[string]string)

	// Split the output into blocks using two consecutive newline characters
	blocks := strings.Split(string(output), "\n\n")

	// Process each block
	for _, block := range blocks {
		// Split the block into lines
		lines := strings.Split(block, "\n")

		fmt.Println(lines)

		// Process each line
		for _, line := range lines {
			// Skip processing empty lines
			if line == "" {
				continue
			}

			// Split each line into key and value
			parts := strings.SplitN(line, ":", 2)
			// Skip lines where the delimiter ":" is not found
			if len(parts) < 2 {
				continue
			}

			fmt.Println(parts)

			// Trim leading/trailing spaces from key and value
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			// Add the key-value pair to the map
			result[key] = value
		}
	}

	// Convert the map to JSON
	jsonData, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	// Return the JSON data
	return jsonData, nil
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Define the JSON response
	response, err := getStatus()
	if err != nil {
		log.Fatal(err)
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(response)
}
