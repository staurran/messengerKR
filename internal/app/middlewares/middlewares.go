package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/staurran/messengerKR.git/internal/app/constProject"
	"github.com/staurran/messengerKR.git/internal/app/utils/token"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

func WithAuthCheck(assignedRoles ...constProject.Role) func(ctx *gin.Context) {
	return func(gCtx *gin.Context) {
		err := token.TokenValid(gCtx)
		if err != nil {
			gCtx.String(http.StatusUnauthorized, "Unauthorized")
			gCtx.Abort()
			return
		}
		role_user, err := token.ExtractTokenRole(gCtx)

		if err != nil {
			gCtx.AbortWithStatus(http.StatusForbidden)
		}
		for _, oneOfAssignedRole := range assignedRoles {
			if role_user == oneOfAssignedRole {
				gCtx.Next()
				return
			}
		}
		gCtx.AbortWithStatus(http.StatusForbidden)
		return
	}

}
