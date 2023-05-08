package app

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/staurran/messengerKR.git/internal/app/middlewares"
	authDel "github.com/staurran/messengerKR.git/internal/pkg/auth/delivery"
	ChatServerPackage "github.com/staurran/messengerKR.git/internal/pkg/chat2/pkg/server"
	photoDel "github.com/staurran/messengerKR.git/internal/pkg/photo/delivery"
	PhotoRepository "github.com/staurran/messengerKR.git/internal/pkg/photo/repository"
	photoUC "github.com/staurran/messengerKR.git/internal/pkg/photo/usecase"
	"github.com/staurran/messengerKR.git/service/proto/authProto"
)

var frontendHosts = []string{
	"http://localhost:8080",
	"http://localhost:3000",
	"http://5.159.100.59:3000",
	"http://5.159.100.59:8080",
	"http://192.168.0.2:3000",
	"http://192.168.0.2:8080",
	"http://5.159.100.59:8080",
	"http://192.168.0.45:3000",
	"http://95.163.180.8:3000",
	"http://meetme-app.ru:3000",
	"http://meetme-app.ru:80",
	"http://meetme-app.ru",
	"http://localhost",
	"http://localhost:8080",
	"http://localhost:80",
}

func (a *Application) InitRoutes(db *gorm.DB, authServ authProto.AuthClient) {

	a.Router.Use(func(h http.Handler) http.Handler {
		return middleware.CorsMiddleware(frontendHosts, h)
	})

	a.Router.Use(func(h http.Handler) http.Handler {
		return middleware.AuthMiddleware(authServ, h)
	})

	photoRepo := PhotoRepository.NewPhotoRepo(db)
	ucPhoto := photoUC.NewPhotoUseCase(photoRepo)
	photoDel.RegisterHTTPEndpoints(a.Router, ucPhoto)

	authDel.RegisterHTTPEndpoints(a.Router, authServ)

	chatServerOptions := ChatServerPackage.ServerOptions{
		Addr:       "localhost",
		Port:       3030,
		PathPrefix: "/iuchat/chats",
	}
	chatRouter := ChatServerPackage.InitServer(chatServerOptions)

	a.Router.PathPrefix(chatServerOptions.PathPrefix).Handler(chatRouter)
}
