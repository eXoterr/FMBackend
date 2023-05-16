package handlers

import (
	"log"
	"os"
	"path/filepath"

	"github.com/eXoterr/FManager/utils"
	"github.com/gin-gonic/gin"
)

// Список файлов/папок godoc
// @Summary Список файлов/папок из директории
// @Schemes
// @Param path body string false "Путь до директории"
// @Accept json
// @Produce json
// @Success 200 {object} handlers.FileList
// @Router /pwd [post]
func DirHandler(c *gin.Context) {
	path := utils.GetPath(c.Request)
	files, err := os.ReadDir(path)
	if err != nil {
		log.Println(err)
		return
	}

	listing := []File{}

	for _, file := range files {
		fullPath := filepath.Join(path, file.Name())
		sympath, _ := filepath.EvalSymlinks(fullPath)

		var currentFilePath *string = &sympath

		if len(sympath) == 0 {
			*currentFilePath = file.Name()
		}

		if utils.CheckIsDir(*currentFilePath) {
			listing = append(listing, File{
				Name:  file.Name(),
				IsDir: true,
			})
		} else {
			listing = append(listing, File{
				Name:  file.Name(),
				IsDir: false,
			})
		}

	}

	c.JSON(200, listing)
}
