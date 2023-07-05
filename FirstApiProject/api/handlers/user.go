package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


type User struct {

}
type persion struct{
	Name string `json:"name"`
	Family string `json:"family"`
}

func NewUserHandler() *User{
	return &User{}
}

func (u *User) User(c *gin.Context) {
	c.JSON(http.StatusOK,"user")
	return
}

func (u *User) UserList(c *gin.Context) {
	soheil_persion := persion{Name : "soheil" , Family : "soheil"}
	js , _ := json.Marshal(soheil_persion)
	c.JSON(http.StatusOK,string(js))
	return
}

func (u *User) GetUserWithID(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK,fmt.Sprintf("this user is %s" , id))
	return
}