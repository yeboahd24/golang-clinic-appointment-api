package middleware

import (
    "fmt"
    "net/http"
    "os"
    "strings"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")

        // Check if the token is present and well-formed
        if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
            return
        }

        tokenString = strings.TrimPrefix(tokenString, "Bearer ")

        // Initialize a new instance of `Claims`
        claims := &jwt.StandardClaims{}

        // Parse the JWT string and store the result in `claims`
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            // Ensure that the token method conform to "SigningMethodHMAC"
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }

            // Return the secret signing key
            return []byte(os.Getenv("JWT_SECRET")), nil
        })

        // If parsing is unsuccessful, respond with an error
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }

        // Check if the token is valid
        if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
            c.Set("userID", claims.Subject)
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }

        c.Next()
    }
}
