package handler

import (
	"golang-book/newsfeeder/platform/newsfeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

type newsfeedPostRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func NewsfeedPost(feed *newsfeed.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := newsfeedPostRequest{}
		c.Bind(&requestBody)

		item := newsfeed.Item{
			Title: requestBody.Title,
			Body:  requestBody.Body,
		}

		feed.Add(item)

		c.Status(http.StatusOK)
	}
}
