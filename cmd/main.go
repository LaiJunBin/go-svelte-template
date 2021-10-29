package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    router := gin.New()
	router.LoadHTMLGlob("front-end/public/*.html")
    router.Static("/build", "./front-end/public/build")

    router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
    })

    s := &http.Server{
        Addr:           ":8000",
        Handler:        router,
    }

    s.ListenAndServe()
}