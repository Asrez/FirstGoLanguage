package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/soheilkhaledabdi/FirstApi/api/handlers"
)

func UserRouteHandler (r *gin.RouterGroup){
	user := handlers.NewUserHandler()

	r.GET("/", user.User)
	r.GET("/users", user.UserList)
	r.POST("/:id", user.GetUserWithID)
	// binder 	
	r.POST("/binder/header1" , user.HeaderBinder1)
	r.POST("/binder/header2" , user.HeaderBinder2)

	r.POST("/binder/query1" , user.QueryBinder1)
	r.POST("/binder/query2" , user.QueryBinder2)

	r.POST("/binder/body" , user.BodyBinder)
	r.POST("/binder/form" , user.FormBinder)

	// end binder 
}