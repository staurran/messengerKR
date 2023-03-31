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
	r.Use(middlewares.WithAuthCheck(constProject.Manager, constProject.Admin, constProject.User)).GET("/user", a.CurrentUser)

	protected := r.Group("/message").Use(middlewares.WithAuthCheck(constProject.Manager, constProject.Admin, constProject.User))

	protected.GET("/chats/:id_user/:k", a.GetChats)
	//protected.GET("messenger/chat/:id_chat", a.GetChat)
	protected.POST("/messenger/chat", a.CreateChat)
	protected.DELETE("/messenger/chat/:id_chat", a.DeleteChat)
	//protected.PUT("messenger/chat/avatar", a.ChangePhoto)
	//protected.PUT("messenger/chat/description", a.ChangeDescription)

	//protected.GET("messenger/chat/users/:id_chat", a.GetChatUsers)
	//protected.POST("messenger/chat/users/:id_chat", a.InviteUser)
	//protected.DELETE("messenger/chat/users/:id_chat", a.DeleteChatUser)
	//protected.PUT("messenger/chat/users/:id_chat", a.ChangeChatRole)

	//protected.POST("/messenger/message/:id_chat", a.CreateMessage)
	//protected.Get("messenger/message/view/:id_message", a.PostShowStatus)
	//protected.POST("/messenger/message/reaction", a.CreateReaction)
	//protected.DELETE("/messenger/message/reaction", a.DeleteReaction)

	//protected.GET("/messenger/contacts", a.GetContacts)
	//protected.POST("/messenger/contacts", a.CreateContact)
	//protected.DELETE("/messenger/contacts", a.DeleteContact)

	//protected.GET("/messenger/profile/:username", a.GetProfile)
	//protected.PUT("/messenger/profile", a.ChangeProfile)

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
