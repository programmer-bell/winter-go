package server

import(

	"github.com/gin-gonic/gin"
	"net/http"
)


func NewRouter() *gin.Engine{
	r := gin.Default()

	r.GET("/health",func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
			"status": "healthy",
			"content-type":"jsonplaceholder",
		})
	})

	r.GET("/post",func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
			"status": "get-post",
			"content-type":"jsonplaceholder",
		})
	})
	return  r


}
