package types

import "github.com/gin-gonic/gin"

//RouterEngine is an alias for gin.Engine
type RouterEngine = gin.Engine

//ContextRequest is an alias for gin.Context
type ContextRequest = gin.Context

//RestError ...
type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}
