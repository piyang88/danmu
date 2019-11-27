package main

import "github.com/justinas/alice"

func main(){
	r:=gin.Default()
	r.GET("/hello",func(ctx *gin.Context){
			ctx.JSON(200,gin.H{
				"message":"hello,world",
			})
	})
    r.Run(":9876")
}

