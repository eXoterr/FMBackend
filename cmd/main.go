package main

import (
	"flag"
	"log"
	"net/http"

	_ "github.com/eXoterr/FManager/docs"
	"github.com/eXoterr/FManager/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	listen *string
)

func init() {
	listen = flag.String("listen", "0.0.0.0:5001", "Listen address (0.0.0.0:5001)")
}

func main() {
	flag.Parse()
	server := gin.New()
	server.POST("/pwd", cors(handlers.DirHandler))
	server.POST("/mkdir", cors(handlers.MKDirHandler))
	server.POST("/mv", cors(handlers.MoveFilesHandler))
	server.POST("/rm", cors(handlers.RemoveHandler))
	server.POST("/exists", cors(handlers.ExistsHandler))

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := http.ListenAndServe(*listen, server)
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
