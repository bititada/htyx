package middleware

import (
	"htyx/handler"
	"htyx/lib/errno"
	"htyx/lib/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if pl, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		} else {
			c.Set("uid", pl.Uid)
			c.Next()
		}
	}
}
