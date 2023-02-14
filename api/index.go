package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	API_BASE_URI = os.Getenv("API_BASE_URI")
	API_KEY      = os.Getenv("API_KEY")
	API_JSON_KEY = "song.json"
	API_HTML_KEY = "index.html"
)

type CurrentListen struct {
	Timestamp time.Time `json:"Timestamp"`
	Content   string    `json:"Content"`
	Checksum  string    `json:"Checksum"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	logger := log.Default()

	logger.Println("---- Starting fetch")
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", API_BASE_URI, API_JSON_KEY), nil)
	if err != nil {
		logger.Printf("---- error creating request, %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("X-Api-Key", API_KEY)
	http.DefaultClient.Timeout = time.Second * 30

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Printf("---- error making request to api, %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()

	logger.Printf("--- res status: %d", res.StatusCode)

	b, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Printf("---- error unmarshalling body, %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}
