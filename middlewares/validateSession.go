package middleware

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ValidateSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("emailVerified") == nil {
			c.Redirect(http.StatusFound, "/")
			c.Abort()
			return
		}

		if !session.Get("emailVerified").(bool) {
			c.Redirect(http.StatusFound, "/")
			c.Abort()
			return
		}

		c.Next()

	}
}
