package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
func Authentication(context *gin.Context) {
    token := context.Request.Header.Get("Authorization")
    if token == "" {
        context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token is missing"})
        return
    }

    if len(token) > 7 && token[:7] == "Bearer " {
        token = token[7:]
    } else {
        context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token format"})
        return
    }

    studentId, err := verifyToken(token)
    if err != nil {
        context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
        return
    }
    context.Set("studentId", studentId)
    context.Next()
}
