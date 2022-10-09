package main

import (
	"KokoChatting/router"
	"fmt"
)

func main(){
	engine := router.Routers()
	if err := engine.Run(":8080");err != nil{
		// TODO:暂时打印,后面改成日志
		fmt.Println(err)
	}
}