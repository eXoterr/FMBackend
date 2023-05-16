package handlers

import (
	"log"
	"os"

	"github.com/eXoterr/FManager/utils"
	"github.com/gin-gonic/gin"
)

// Создает папку по указанному пути godoc
// @Summary Создает папку по указанному пути
// @Schemes
// @Param path body string false "Путь до директории"
// @Accept plain
// @Produce plain
// @Success 200
// @Failure 500
// @Router /mkdir [post]
func MKDirHandler(c *gin.Context) {
	path := utils.GetPath(c.Request)

	err := os.Mkdir(path, os.FileMode(int(0777)))

	if err != nil {
		log.Println(err)
		c.String(500, "", err)
	}
}
