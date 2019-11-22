package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var name string

func AuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		if cookie, err := context.Request.Cookie("Username"); err ==nil {
			value := cookie.Value
			if name == value {
				context.Next()
				return
			}
		}
		context.JSON(http.StatusOK,gin.H{
			"code":200,
			"message": "hello guest",
		})
		context.Abort()
		return
	}
}

func main()  {
	r :=gin.Default()
	r.GET("/login", func(context *gin.Context) {
		name =context.Query("Username")
		cookie := &http.Cookie{
			Name:     "Usernsme",
			Value:    name,
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(context.Writer,cookie)
		context.String(http.StatusOK,"Login success !")
	})
	r.GET("/home", AuthMiddleWare(), func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{
			"code":200,
			"message": "hello " + name,
		})
	})
	r.Run()
}