package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

//StartApplication ...
func StartApplication() {
	MapUrls(router)
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
