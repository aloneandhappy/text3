package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main()  {
	r := gin.Default()
	r.POST("/Xuehao", func(c *gin.Context) {
		xuehao :=c.PostForm("Xuehao")
		resp,_:=http.Get("http://jwzx.cqupt.edu.cn/data/json_StudentSearch.php?searchKey="+xuehao)
		defer resp.Body.Close()
		body,_ :=ioutil.ReadAll(resp.Body)
		c.String(http.StatusOK,string(body))
	})
	r.Run()
}
