package handlers

import (
	"log"
	"os"
	"path/filepath"

	"github.com/eXoterr/FManager/utils"
	"github.com/gin-gonic/gin"
)

func DirHandler(c *gin.Context) {
	path := utils.GetPath(c.Request)
	files, err := os.ReadDir(path)
	if err != nil {
		log.Println(err)
		return
	}

	listing := []string{}

	for _, file := range files {
		fullPath := filepath.Join(path, file.Name())
		sympath, _ := filepath.EvalSymlinks(fullPath)

		var currentFilePath *string = &sympath

		if len(sympath) == 0 {
			*currentFilePath = file.Name()
		}

		if utils.CheckIsDir(*currentFilePath) {
			listing = append(listing, file.Name()+"/")
		} else {
			listing = append(listing, file.Name())
		}

	}

	c.JSON(200, listing)
}
