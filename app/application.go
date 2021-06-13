package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sebagalan/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

//StartApplication ...
func StartApplication() {
	MapUrls(router)

	logger.Info("Start - app")
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
