package handler

import (
	"golang-book/newsfeeder/platform/newsfeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Getter interface created
// func NewsfeedGet(feed *newsfeed.Repo) gin.HandlerFunc {
// using interface
func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		feeds := feed.GetAll()
		c.JSON(http.StatusOK, feeds)
	}
}
