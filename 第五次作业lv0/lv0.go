package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var account map[string]User = make(map[string]User)

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password   string `form:"password" json:"password" bdinding:"required"`
	Age      int    `form:"age" json:"age"`
}

func main()  {
	r := gin.Default()
	r.POST("/login", func(context *gin.Context) {
		var user User
		err := context.Bind(&user)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		if v,ok := account[user.Username]; ok &&v.Password ==user.Password {
			context.JSON(http.StatusOK,gin.H{
				"username" : v.Username,
				"password" : v.Password,
				"age" : v.Age,
			})
		} else {
			context.JSON(http.StatusOK,gin.H{
				"message" : "账户或密码有错",
			})
		}
	})
	r.POST("/register", func(context *gin.Context) {
		var user User
		err := context.Bind(&user)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		username := user.Username
		if _,ok := account[username]; ok {
			message := "用户名" + username + "已存在"
			context.JSON(http.StatusOK,gin.H{
				"code" : 200,
				"message" : message,
			})
		} else {
			account[username] = user
			context.JSON(http.StatusOK,gin.H{
				"code" : 200,
				"message" : "注册成功",
			})
		}
	})
	r.Run()
}