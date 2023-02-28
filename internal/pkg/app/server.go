package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"lab3/internal/app/middlewares"
	"lab3/internal/app/role"
	"log"
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

	r.GET("/goods", a.GetAll)
	r.GET("/goods/:id", a.GetProduct)

	r.GET("/ping/:name", a.Ping)

	r.Use(middlewares.WithAuthCheck(role.Manager, role.Admin, role.User)).POST("/order", a.AddOrder)
	r.Use(middlewares.WithAuthCheck(role.Manager, role.Admin, role.User)).GET("/order", a.GetAllOrders)
	r.Use(middlewares.WithAuthCheck(role.Manager, role.Admin, role.User)).DELETE("/order/:id", a.DeleteOrder)

	r.Use(middlewares.WithAuthCheck(role.Manager, role.Admin, role.User)).GET("/order-status", a.GetStatus)

	r.Use(middlewares.WithAuthCheck(role.Manager, role.Admin)).GET("/user", a.CurrentUser)
	r.Use(middlewares.WithAuthCheck(role.Manager, role.Admin)).POST("/goods", a.PostProduct)
	r.Use(middlewares.WithAuthCheck(role.Manager, role.Admin)).PUT("/goods", a.ChangeProduct)
	r.Use(middlewares.WithAuthCheck(role.Manager, role.Admin)).DELETE("goods/:id", a.DeleteProduct)

	r.Use(middlewares.WithAuthCheck(role.Manager, role.Admin)).GET("goods/all-orders", a.GetOrders)

	r.Use(middlewares.WithAuthCheck(role.Manager, role.Admin)).PUT("/order/:id_order/:id_status", a.ChangeStatus)

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
