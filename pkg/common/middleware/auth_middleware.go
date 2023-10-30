package middleware

import (
	"github.com/leeeo2/backend/pkg/common/jwt"
	"github.com/leeeo2/backend/pkg/dao"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// obtain token string
		tokenStr := ctx.GetHeader("Authorization")

		// validate
		if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "permission denied",
			})
			ctx.Abort()
			return
		}

		tokenStr = tokenStr[7:]

		token, claims, err := jwt.ParseToken(tokenStr)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "permission denied",
			})
			ctx.Abort()
			return
		}

		// pass
		userId := claims.UserId
		user, err := dao.DescribeUserById(userId)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "permission denied",
			})
			ctx.Abort()
			return
		}

		// write user to context
		ctx.Set("user", user)

		ctx.Next()

	}
}
