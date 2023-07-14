package advmobileinfo

import (
	"encoding/json"
	"log"
	"net/http"
)

func Run() {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":50500", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Define the JSON response
	response := map[string]interface{}{
		"message": "Hello, world!",
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Convert the response to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.Write(jsonResponse)
}
