package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func dirHandler(w http.ResponseWriter, r *http.Request) {
	path := getPath(r)
	files, err := os.ReadDir(path)
	if err != nil {
		log.Println(err)
		return
	}

	listing := []string{}

	for _, file := range files {
		if file.IsDir() {
			listing = append(listing, file.Name()+"/")
		} else {
			listing = append(listing, file.Name())
		}
	}

	j, err := json.Marshal(listing)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.Write(j)
}

func main() {
	http.HandleFunc("/pwd", dirHandler)
	http.HandleFunc("/mkdir", mkDirHandler)

	http.ListenAndServe("127.0.0.1:5001", nil)
}

func mkDirHandler(w http.ResponseWriter, r *http.Request) {
	path := getPath(r)

	err := os.Mkdir(path, os.FileMode(int(0777)))

	if err != nil {
		log.Println(err)
		return
	}
}

func getPath(r *http.Request) string {
	buffer, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return ""
	}

	return string(buffer)
}
