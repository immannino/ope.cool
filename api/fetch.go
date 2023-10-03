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
	API_BASE_URI  = os.Getenv("API_BASE_URI")
	API_KEY       = os.Getenv("API_KEY")
	API_SONG_KEY  = "song.json"
	API_SONGS_KEY = "songs.json"
	API_HTML_KEY  = "index.html"

	logger = log.Default()
)

type CurrentListen struct {
	Timestamp time.Time `json:"Timestamp"`
	Content   string    `json:"Content"`
	Checksum  string    `json:"Checksum"`
}

type FetchResponse struct {
}

func Fetch(w http.ResponseWriter, r *http.Request) {
	path := fmt.Sprintf("%s/%s", API_BASE_URI, API_SONG_KEY)

	if r.URL.Query().Get("many") != "" {
		path = fmt.Sprintf("%s/%s", API_BASE_URI, API_SONGS_KEY)
	}

	b, err := fetch(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}

func fetch(path string) ([]byte, error) {
	logger.Println("---- Starting fetch")
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		logger.Printf("---- error creating request, %v", err)
		return nil, err
	}

	req.Header.Set("X-Api-Key", API_KEY)
	http.DefaultClient.Timeout = time.Second * 30

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Printf("---- error making request to api, %v", err)
		return nil, err
	}

	defer res.Body.Close()

	logger.Printf("--- res status: %d", res.StatusCode)

	b, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Printf("---- error unmarshalling body, %v", err)
		return nil, err
	}

	return b, nil
}
