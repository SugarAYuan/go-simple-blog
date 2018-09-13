package main

import "github.com/gin-gonic/gin"

func main () {
	r := gin.Default()
	r.GET("/" , func(context *gin.Context) {
		context.Writer.WriteString("Hello World.")

	})


	r.Run("127.0.0.1:8081")
}