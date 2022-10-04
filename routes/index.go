package routes

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default()
func Run() {
	router.SetTrustedProxies(nil)
	getRoutes()
	router.Run(":3000")
}

func getRoutes() {
	v1 := router.Group("/v1")
	AddTodoRoutes(v1)
}