package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soheilkhaledabdi/FirstApi/api/helper"
)


type header struct {
	userId string

}

type User struct {

}
type persion struct{
	Name string `json:"name" binding:"required , alpha,min:4"`
	Family string `json:"family" binding:"required" ,max:255,min:10` 
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


func (u *User) HeaderBinder1(c *gin.Context) {
	userID := c.GetHeader("UserID")

	c.JSON(http.StatusOK , gin.H{
		"result" : "headerBinder1",
		"userID" : userID,
	})
}

func (u *User) HeaderBinder2(c *gin.Context) {
	h := header{}
	c.BindHeader(&h)
	c.JSON(http.StatusOK , gin.H{
		"result" : "headerBinder2",
		"header" : h,
	})
}

func (u *User) QueryBinder1(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	
	c.JSON(http.StatusOK ,helper.GenerateBaseResponse(gin.H{
		"result" : "headerBinder2",
		"id" : id,
		"name" : name,
	} , true , 200))
}

func (u *User) QueryBinder2(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	c.JSON(http.StatusOK , gin.H{
		"result" : "headerBinder2",
		"id" : id,
		"name" : name,
	})
}


func (u *User) BodyBinder(c *gin.Context) {
	 p := persion{}

	err := c.ShouldBindJSON(&p)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest , gin.H{
			"validtion" : err.Error(),
		})
	}

	c.JSON(http.StatusOK , gin.H{
		"result" : "headerBinder2",
		"persion" : p,
	})
}



func (u *User) FormBinder(c *gin.Context) {
	p := persion{}

	c.Bind(&p)

   c.JSON(http.StatusOK , gin.H{
	   "result" : "headerBinder2",
	   "persion" : p,
   })
}