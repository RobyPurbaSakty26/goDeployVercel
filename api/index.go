package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// . "github.com/tbxark/g4vercel"
)

// func Handlmaer(w http.ResponseWriter, r *http.Request) {
// 	server := New()

// 	server.GET("/", func(context *Context) {
// 		context.JSON(200, H{
// 			"message": "hello go from vercel !!!!",
// 		})
// 	})
// 	server.GET("/hello", func(context *Context) {
// 		name := context.Query("name")
// 		if name == "" {
// 			context.JSON(400, H{
// 				"message": "name not found",
// 			})
// 		} else {
// 			context.JSON(200, H{
// 				"data": fmt.Sprintf("Hello %s!", name),
// 			})
// 		}
// 	})
// 	server.GET("/user/:id", func(context *Context) {
// 		context.JSON(400, H{
// 			"data": H{
// 				"id": context.Param("id"),
// 			},
// 		})
// 	})
// 	server.GET("/long/long/long/path/*test", func(context *Context) {
// 		context.JSON(200, H{
// 			"data": H{
// 				"url": context.Path,
// 			},
// 		})
// 	})
// 	server.Handle(w, r)
// }

func Handler(req *http.Request) {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.ServeHTTP(http.ResponseWriter, req)

}
