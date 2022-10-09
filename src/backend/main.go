package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){
	engine := gin.Default()
	if err := engine.Run(":8080");err != nil{
		// TODO:暂时打印吧,后面改成日志
		fmt.Println(err)
	}
}