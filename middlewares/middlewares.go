package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	auth "github.com/jolycky/task-5-pbi-btpns-JoseBagusRamadhan/app/authentication"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}

		err := auth.ValidateToken(strings.Split(tokenString, "Bearer ")[1])
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
