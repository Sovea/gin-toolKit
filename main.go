package main

import (
	"toolKit/app/toolkit/service/service_xrate"
	"toolKit/app/toolkit/swagger_routers"
	"toolKit/routers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/unrolled/secure"
	"net/http"
)

//Cors cross-domain solution(Gin Middleware)
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,X-Nideshop-Token, x-nideshop-token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusOK)
		}
		context.Next()
	}
}

//TlsHandler is func to use TLS
func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:8779",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}

// @title toolKit
// @version 1.0
// @description This is a toolKit for server with Gin.
// @contact.name sovea
// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT
func main() {
	//RawInit
	Gin_Instance := routers.RawInit()
	//CORS
	Gin_Instance.Use(Cors())
	// Access rate limit
	Gin_Instance.Use(service_xrate.DefaultLimitMiddleware())
	//TLS
	Gin_Instance.Use(TlsHandler())
	//Include Some Router
	routers.IncludeWith(Gin_Instance, swagger_routers.Routers)
	// Static resource loading
	Gin_Instance.StaticFS("/public", http.Dir("./public"))
	Gin_Instance.StaticFile("/favicon.png", "../resources/index.png")
	//Start server
	Gin_Instance.RunTLS("127.0.0.1:8779", "./resources/ssl/toolkit.crt", "./resources/ssl/toolkit.key")
}
