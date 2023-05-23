package app

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/staurran/messengerKR.git/internal/app/middlewares"
	authDel "github.com/staurran/messengerKR.git/internal/pkg/auth/delivery"
	"github.com/staurran/messengerKR.git/service/proto/authProto"

	chatDel "github.com/staurran/messengerKR.git/internal/pkg/chat/delivery"
	ChatRepository "github.com/staurran/messengerKR.git/internal/pkg/chat/repository"
	chatUC "github.com/staurran/messengerKR.git/internal/pkg/chat/usecase"

	userDel "github.com/staurran/messengerKR.git/internal/pkg/user/delivery"
	UserRepository "github.com/staurran/messengerKR.git/internal/pkg/user/repository"
	userUC "github.com/staurran/messengerKR.git/internal/pkg/user/usecase"
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
		return middleware.AuthMiddleware(authServ, h)
	})

	a.Router.Use(func(h http.Handler) http.Handler {
		return middleware.CorsMiddleware(frontendHosts, h)
	})

	a.Router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Printf("OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		return
	})

	chatRepo := ChatRepository.NewChatRepo(db)
	ucChat := chatUC.NewChatUseCase(chatRepo)
	chatDel.RegisterHTTPEndpoints(a.Router, ucChat)

	userRepo := UserRepository.NewUserRepo(db)
	ucUser := userUC.NewUserUseCase(userRepo)
	userDel.RegisterHTTPEndpoints(a.Router, ucUser)

	authDel.RegisterHTTPEndpoints(a.Router, authServ)

}
