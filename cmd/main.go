package main

import (
	"Nerve/internal"
	"github.com/gin-gonic/gin"
)

//func CORSMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		// in deploy need be set global ip
//		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
//		c.Header("Access-Control-Allow-Credentials", "true")
//		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
//		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")
//
//		if c.Request.Method == "OPTIONS" {
//			c.AbortWithStatus(204)
//			return
//		}
//
//		c.Next()
//	}
//}

func main() {
	gin.SetMode("debug")
	router := gin.Default()
	//router.Use(CORSMiddleware())
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	internal.SetRouters(router)
	// go web.Broadcaster()
	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
