package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/appleboy/gin-jwt.v2"

	"massliking/backend/models"
)

var JWT *jwt.GinJWTMiddleware

func Init(jwtRealm string, jwtSecret string, jwtTTL int) {
	JWT = &jwt.GinJWTMiddleware{
		Realm:      jwtRealm,
		Key:        []byte(jwtSecret),
		Timeout:    time.Hour * time.Duration(jwtTTL),
		MaxRefresh: time.Hour * time.Duration(jwtTTL),
		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
			user, err := models.GetUser(userId)
			if err != nil {
				return userId, false
			}

			verified := user.Verify(password)

			return userId, verified
		},
		Authorizator: func(userId string, c *gin.Context) bool {
			_, err := models.GetUser(userId)
			if err != nil {
				return false
			} else {
				return true
			}
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",
	}
}
