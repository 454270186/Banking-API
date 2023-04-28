package main

import (
	"RESTful/domain"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleWare struct {
	repo domain.AuthRepo
}

func (am AuthMiddleWare) AuthMid() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		currentRouteParams := ctx.Params
		authHeader := ctx.Request.Header.Get("Authorization")

		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing token",
			})
			return
		}

		token := getTokenFromHeader(authHeader)
		isAuthorized := am.repo.IsAuthorized(token, currentRouteParams)
		if isAuthorized {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Unauthorized",
			})
		}
	}
}

func getTokenFromHeader(header string) string {
	token := strings.TrimPrefix(header, "Bearer ")
	log.Println(token)
	return token
}