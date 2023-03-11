package app

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/staurran/messengerKR.git/internal/app/constProject"
	"github.com/staurran/messengerKR.git/internal/app/middlewares"
)

func (a *Application) StartServer() {
	log.Println("Server start up")
	log.Println("Server start up")
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://localhost:3000"}
	config.AllowHeaders = []string{"content-type", "Authorization"}
	config.AllowMethods = []string{"PUT", "PATCH", "GET", "POST", "DELETE"}
	r.Use(cors.New(config))

	public := r.Group("/api")
	public.POST("/register", a.Register)
	public.POST("/login", a.Login)

	r.GET("/messenger", a.GetAll)

	r.Use(middlewares.WithAuthCheck(constProject.Manager, constProject.Admin, constProject.User)).GET("/messenger/k", a.GetChats)
	r.Use(middlewares.WithAuthCheck(constProject.Manager, constProject.Admin, constProject.User)).GET("messenger/chat/:id_chat", a.GetChat)

	r.Use(middlewares.WithAuthCheck(constProject.Manager, constProject.Admin, constProject.User)).POST("/order", a.AddOrder)
	r.Use(middlewares.WithAuthCheck(constProject.Manager, constProject.Admin, constProject.User)).GET("/order", a.GetAllOrders)
	r.Use(middlewares.WithAuthCheck(constProject.Manager, constProject.Admin, constProject.User)).DELETE("/order/:id", a.DeleteOrder)

	r.Use(middlewares.WithAuthCheck(constProject.Manager, constProject.Admin, constProject.User)).GET("/order-status", a.GetStatus)

	r.Use(middlewares.WithAuthCheck(constProject.Manager, constProject.Admin)).GET("/user", a.CurrentUser)
	r.Use(middlewares.WithAuthCheck(constProject.Manager, constProject.Admin)).POST("/goods", a.PostProduct)
	r.Use(middlewares.WithAuthCheck(constProject.Manager, constProject.Admin)).PUT("/goods", a.ChangeProduct)
	r.Use(middlewares.WithAuthCheck(constProject.Manager, constProject.Admin)).DELETE("goods/:id", a.DeleteProduct)

	r.Use(middlewares.WithAuthCheck(constProject.Manager, constProject.Admin)).GET("goods/all-orders", a.GetOrders)

	r.Use(middlewares.WithAuthCheck(constProject.Manager, constProject.Admin)).PUT("/order/:id_order/:id_status", a.ChangeStatus)

	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Println("Run failed")
	}
	log.Println("Server down")
}

type AnswerJSON struct {
	Status      string `json:"status"`
	Description string `json:"description"`
}
