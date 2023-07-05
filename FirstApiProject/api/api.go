package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	routes "github.com/soheilkhaledabdi/FirstApi/api/routes"
)



func InitServer(){
	r := gin.New()
	r.Use(gin.Logger()  ,gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/hello" , func (c *gin.Context){
			c.JSON(http.StatusOK, "Hello world!")
			return
		})
		user := v1.Group("/user")

		routes.UserRouteHandler(user)
	}

	r.Run(":5005")
}