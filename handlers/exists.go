package handlers

import (
	"log"
	"os"

	"github.com/eXoterr/FManager/utils"
	"github.com/gin-gonic/gin"
)

// Проверяет есть ли путь(файл/папка) godoc
// @Summary Проверяет есть ли путь(файл/папка)
// @Schemes
// @Param path body string false "Путь до директории"
// @Accept plain
// @Produce json
// @Success 200
// @Failure 404
// @Router /exists [post]
func ExistsHandler(c *gin.Context) {
	body := utils.GetPath(c.Request)

	_, err := os.Stat(body)

	if err != nil {
		log.Println(err)
		c.Status(404)
		return
	}

	c.Status(200)
}
