package main

import (
	"github.com/gin-gonic/gin"
	"goAuth/controller"
	"goAuth/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.Connect()
	initializers.SyncDb()
}

func main() {
	r := gin.Default()
	r.POST("/sing-up", controller.Singup)
	r.GET("/login", controller.Login)
	r.Run() 
