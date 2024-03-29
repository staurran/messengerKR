package middleware

import (
	"net/http"
	"strings"

	"golang.org/x/net/context"

	"github.com/staurran/messengerKR.git/internal/app/cookie"
	"github.com/staurran/messengerKR.git/internal/app/servicedefault"
	"github.com/staurran/messengerKR.git/service/proto/authProto"
	"github.com/staurran/messengerKR.git/utils/logger"
	"github.com/staurran/messengerKR.git/utils/writer"
)

func CorsMiddleware(allowedHosts []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})

}

var ContextUserKey = "userId"
var ProtectedPath = "/iuchat/"

func AuthMiddleware(authServ authProto.AuthClient, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, ProtectedPath) {
			next.ServeHTTP(w, r)
			return
		}
		token, err := cookie.GetValueCookie(r, servicedefault.SessionTokenCookieName)
		if err != nil {
			logger.Log(http.StatusUnauthorized, err.Error(), r.Method, r.URL.Path, true)
			writer.ErrorRespond(w, r, err, http.StatusUnauthorized)
			return
		}

		reqBody := authProto.UserToken{
			Token: token,
		}
		userResp, err := authServ.CheckSession(r.Context(), &reqBody)
		if err != nil {
			logger.Log(http.StatusUnauthorized, err.Error(), r.Method, r.URL.Path, true)
			writer.ErrorRespond(w, r, err, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ContextUserKey, userResp.UserId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
