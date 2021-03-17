package middleware

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//tokenString:=ctx.GetHeader("Authorization")
		//if tokenString==""||!strings
	}
}
