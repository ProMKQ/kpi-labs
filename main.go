package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
}

func timeHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("Received request:", request.Method, request.URL.Path)
	if request.Method != http.MethodGet {
		http.Error(writer, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}
	response := TimeResponse{Time: time.Now().Format(time.RFC3339)}
	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		log.Println("Error: Could not encode response:", err)
	}
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "/time", http.StatusFound)
}

func main() {
	f, err := os.OpenFile("F:\\test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/time", timeHandler)

	port := ":8795"
	log.Println("Starting server on http://localhost" + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Server failed:", err)
	}
}
