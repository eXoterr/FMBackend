package handlers

import (
	"fmt"

	"github.com/eXoterr/FManager/utils"
	"github.com/gin-gonic/gin"
)

// Удаляет папку/файл по переданному пути godoc
// @Summary Удаляет папку/файл по переданному пути
// @Schemes
// @Param path body string false "Путь до директории"
// @Accept plain
// @Produce json
// @Success 200
// @Failure 500
// @Router /rm [post]
func RemoveHandler(c *gin.Context) {
	body := utils.GetPath(c.Request)

	fmt.Printf("removed %v", body)
	c.Status(200)
}
