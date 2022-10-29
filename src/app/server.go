package main

import (
	"net/http"
	"os"

	"money-management/src/config"
	"money-management/src/pkg/databases/postgre"
	"money-management/src/pkg/helpers"

	"github.com/gin-gonic/gin"
)

func main() {
	//init database
	helper.InitLogger()
	defer helper.Logger.Sync()

	db, err := postgre.InitConnection()
	if err != nil {
		helper.Logger.Error(err.Error())
		os.Exit(1)
	}
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