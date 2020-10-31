package infrastructure

import (
	"github.com/gin-gonic/gin"

	"api/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := controllers.NewUserController(NewSqlHandler())

	router.POST("api/v1/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("api/v1/users", func(c *gin.Context) { userController.Index(c) })
	router.GET("api/v1/users/:id", func(c *gin.Context) { userController.Show(c) })

	Router = router
}
