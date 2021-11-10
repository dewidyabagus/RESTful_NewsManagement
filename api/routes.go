package api

import (
	echo "github.com/labstack/echo/v4"

	"RESTful/api/v1/post"
	"RESTful/api/v1/topic"
)

func RegisterRouters(e *echo.Echo, topic *topic.Controller, post *post.Controller) {
	if topic == nil || post == nil {
		panic("route parameter initialization failed")
	}

	topicGroup := e.Group("/v1/topics")
	topicGroup.POST("", topic.InsertTopic)
	topicGroup.GET("", topic.FindAllTopic)
	topicGroup.GET("/:id", topic.FindTopicById)
	topicGroup.PUT("/:id", topic.UpdateTopic)
	topicGroup.DELETE("/:id", topic.DeleteTopic)

	postGroup := e.Group("/v1/posts")
	postGroup.POST("", post.InsertPost)
	postGroup.GET("", post.FindAllPost)
}
