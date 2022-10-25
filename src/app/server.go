package main

import (
	"net/http"

	"money-management/src/config"
	"money-management/src/pkg/databases/postgre"

	"github.com/gin-gonic/gin"
)

func main() {
	//init database
	db, _ := postgre.InitConnection()
	defer postgre.CloseConnection(db)

	//init GIN
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "Running...",
		})
	})

	r.Run(":" + config.Get().AppPort)
}