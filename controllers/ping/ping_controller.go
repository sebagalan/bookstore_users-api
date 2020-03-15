package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/sebagalan/bookstore_users-api/types"
)

//Ping ...
func Ping(c *types.ContextRequest) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
