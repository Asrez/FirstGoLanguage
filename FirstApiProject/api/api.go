package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	// "github.com/soheilkhaledabdi/FirstApi/api/middlewares"
	routes "github.com/soheilkhaledabdi/FirstApi/api/routes"
	validtion "github.com/soheilkhaledabdi/FirstApi/api/validation"
)



func InitServer(){
	r := gin.New()
	r.Use(gin.Logger()  ,gin.Recovery(), /*middlewares.UserMiddleware()*/)
	RegisterValidtionCustom()
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

func RegisterValidtionCustom(){
		val, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			val.RegisterValidation("mobile", validtion.IranianMobileNumberValidator, true)

		}
}