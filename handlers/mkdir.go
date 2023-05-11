package handlers

import (
	"log"
	"os"

	"github.com/eXoterr/FManager/utils"
	"github.com/gin-gonic/gin"
)

func MKDirHandler(c *gin.Context) {
	path := utils.GetPath(c.Request)

	err := os.Mkdir(path, os.FileMode(int(0777)))

	if err != nil {
		log.Println(err)
		return
	}
}
