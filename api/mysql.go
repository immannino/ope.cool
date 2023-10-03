package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/immannino/ope.cool/pkg/orm"
)

var (
	db      *sql.DB
	querier *orm.Queries
)

func init() {
	// setup DB stuff
	log.Println("connecting to DB")
	db, err := sql.Open("mysql", os.Getenv("DATABASE_URI"))
	if err != nil {
		fmt.Println("Error connecting to DB", err)
		panic(err)
	}

	log.Println("connecting to Querier")
	querier = orm.New(db)
}

func Mysql(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var songs []orm.Listen

	if r.URL.Query().Get("many") != "" {
		log.Println("Fetching many")
		resp, err := querier.GetLastNListens(ctx, 10)
		if err != nil {
			respondError(w, err)
			return
		}
		songs = resp
	} else if r.URL.Query().Get("top") != "" {
		top, err := querier.GetTopNListens(ctx, 10)
		if err != nil {
			respondError(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(top)
		return
	} else {
		log.Println("Fetching single")
		resp, err := querier.GetLatestListen(ctx)
		if err != nil {
			respondError(w, err)
			return
		}
		songs = append(songs, resp)
	}

	b, err := json.Marshal(songs)
	if err != nil {
		respondError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}

func respondError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(err.Error()))
}
