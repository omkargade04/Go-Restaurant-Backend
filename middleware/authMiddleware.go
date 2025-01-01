package middleware

import (
	helper "go-restro-backend/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "No Authorization header provided"})
			c.Abort()
			return
		}

		claims, err := helper.ValdateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid", claims.Uid)

		c.Next()
	}
}
