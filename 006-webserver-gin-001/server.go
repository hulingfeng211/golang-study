package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID   string
	Name string
}

func main() {

	users := []User{{
		ID: "1", Name: "name1",
	}, {
		ID: "2", Name: "name2",
	}}
	r := gin.Default()
	r.GET("/users", func(context *gin.Context) {
		//todo
		//context.DefaultQuery()
	})
	//create user
	r.POST("/users", func(context *gin.Context) {
		//todo
	})
	//update user
	r.PUT("/users/:id", func(context *gin.Context) {
		//todo
		id := context.Param("id")
		context.String(200, id)
	})
	//update user
	r.DELETE("/users/:id", func(context *gin.Context) {
		//todo
		id := context.Param("id")
		context.String(200, id)
	})

	r.GET("/", func(c *gin.Context) {
		//c.JSON(200,gin.H{
		//	"Blog":"www.flysnow.org",
		//	"wechat":"flysnow_org",
		//})
		//c.Request
		c.JSON(200, users)
	})

	r.Run(":8080")

}
