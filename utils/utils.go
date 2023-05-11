package utils

import (
	"io"
	"log"
	"net/http"
	"os"
)

func CheckIsDir(path string) bool {
	fileStat, err := os.Lstat(path)
	if err != nil {
		log.Println(err)
	}

	return fileStat.IsDir()
}

func GetPath(r *http.Request) string {
	buffer, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return ""
	}

	return string(buffer)
}
