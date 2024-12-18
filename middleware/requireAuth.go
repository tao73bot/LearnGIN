package middleware

import (
	"Go-Gin/initializers"
	"Go-Gin/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(ctx *gin.Context) {
	// code for RequireAuth middleware
	fmt.Println("RequireAuth middleware")
	// get the cookie from the request
	tokenString, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
	}
	// decode the cookie
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return []byte(os.Getenv("JWT_SECRECT")), nil
	})
	if err != nil {
		log.Fatal(err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check if expired
		if exp, ok := claims["exp"].(float64); ok && float64(time.Now().Unix()) > exp {
			ctx.AbortWithError(http.StatusUnauthorized, err)
		}
		// find the user with token sub
		var user models.User
		result := initializers.Db.First(&user, claims["sub"])
		if result.Error != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
		}
		// attach to req
		ctx.Set("user", user)
		// call next
		ctx.Next()
	} else {
		fmt.Println(err)
	}
}
