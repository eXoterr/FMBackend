package main

import (
	"log"
	"net/http"

	"github.com/eXoterr/FManager/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()
	server.POST("/pwd", cors(handlers.DirHandler))
	server.POST("/mkdir", cors(handlers.MKDirHandler))
	server.POST("/mv", cors(handlers.MoveFilesHandler))

	err := http.ListenAndServe("127.0.0.1:5001", server)
	if err != nil {
		log.Fatal(err)
	}
}

func cors(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		next(c)
	}
}
