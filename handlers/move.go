package handlers

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type FilesToMove struct {
	MoveTo   string         `json:"to"`
	FileList MovingFileList `json:"files"`
}

type MovingFileList []string

// Перемещает по указанному пути godoc
// @Summary Перемещает по указанному пути
// @Schemes
// @Param files body handlers.FilesToMove false "Список файлов"
// @Accept json
// @Produce plain
// @Success 200
// @Failure 500
// @Router /mv [post]
func MoveFilesHandler(c *gin.Context) {
	files := FilesToMove{}
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return
	}
	err = json.Unmarshal(body, &files)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return
	}

	for _, f := range files.FileList {
		err := os.Rename(f, files.MoveTo)
		if err != nil {
			log.Println(err)
			c.Status(500)
			return
		}
	}
}
