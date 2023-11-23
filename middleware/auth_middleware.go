package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/adityatresnobudi/library-api/shared"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if os.Getenv("ENV_MODE") == "testing" {
			c.Next()
			return
		}

		header := c.GetHeader("Authorization")
		splittedHeader := strings.Split(header, " ")
		if len(splittedHeader) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, shared.ErrInvalidAuthHeader.ToErrorDTO())
			return
		}

		token, err := shared.ValidateJWT(splittedHeader[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, shared.ErrInvalidToken.ToErrorDTO())
			return
		}

		claims, ok := token.Claims.(*shared.JWTClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, shared.ErrInvalidToken.ToErrorDTO())
			return
		}

		c.Set("id", claims.ID)

		c.Next()
	}
}
