package handlers

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type FilesToMove struct {
	MoveTo   string `json:"to"`
	FileList []file `json:"files"`
}

type file struct {
	Name string `json:"name"`
	// Type string `json:"type"`
	Path string `json:"path"`
}

func MoveFilesHandler(c *gin.Context) {
	files := FilesToMove{}
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(body, &files)
	if err != nil {
		log.Println(err)
		return
	}

	for _, f := range files.FileList {
		err := os.Rename(f.Path, files.MoveTo)
		if err != nil {
			log.Println(err)
			return
		}
	}
}