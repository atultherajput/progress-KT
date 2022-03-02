package net

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func muxCustomMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func ginCustomMiddleware(c *gin.Context) {
	c.Request.Header.Add("X-Custom", "payload")
	c.Next()
}
