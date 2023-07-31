package pfm

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
)

type Command struct {
	Cmd string `json:"cmd"`
}

func runCommandInBackground(command string) {
	log.Printf("Running command \"%s\"\n", command)
	cmd := exec.Command("bash", "-c", command+" > /proc/1/fd/1")
	err := cmd.Start()
	if err != nil {
		log.Printf("Error running command: %s", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var command Command
	if err := decoder.Decode(&command); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	go runCommandInBackground(command.Cmd)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Command received and started in the background.\n"))
}

func Run() {
	// Print "Server is starting..." message
	port := "50505"
	log.Printf("Server is starting on port %s...\n", port)

	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
