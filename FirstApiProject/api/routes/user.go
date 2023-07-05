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
}