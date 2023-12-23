package middleware

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
    return cors.New(cors.Config{
        // Configure CORS settings here
        AllowAllOrigins: true,
        AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
    })
}

