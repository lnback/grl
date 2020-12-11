package main

import (
	"grl"
	"net/http"
)

func main() {
	r := grl.New()
	r.GET("/", func(c *grl.Context) {
		c.HTML(http.StatusOK, "<h1>Hello grl</h1>")
	})

	r.GET("/hello", func(c *grl.Context) {
		// expect /hello?name=grlktutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *grl.Context) {
		// expect /hello/grlktutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *grl.Context) {
		c.JSON(http.StatusOK, grl.H{"filepath": c.Param("filepath")})
	})

	r.POST("/login", func(context *grl.Context) {
		context.JSON(http.StatusOK,grl.H{
			"username":context.PostForm("username"),
			"password":context.PostForm("password"),
		})
	})
	r.Run(":8080")
}


