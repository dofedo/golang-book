package main

import (
	"golang-book/newsfeeder/httpd/handler"
	"golang-book/newsfeeder/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func main() {

	feed := newsfeed.New()

	router := gin.Default()

	router.GET("/ping", handler.PingGet())
	router.GET("/newsfeed", handler.NewsfeedGet(feed))
	router.POST("/newsfeed", handler.NewsfeedPost(feed))

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
