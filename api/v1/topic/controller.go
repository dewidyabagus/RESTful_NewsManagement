package topic

import (
	echo "github.com/labstack/echo/v4"

	"RESTful/api/common"
	"RESTful/api/v1/topic/request"
	"RESTful/api/v1/topic/response"
	"RESTful/business/topic"
)

type Controller struct {
	service topic.Service
}

func NewController(service topic.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) InsertTopic(ctx echo.Context) error {
	var topic = new(request.Topic)

	if err := ctx.Bind(topic); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	if err := c.service.InsertTopic(topic.ToBusinessTopicSpec()); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) FindAllTopic(ctx echo.Context) error {
	topics, err := c.service.FindAllTopic()
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithData(response.GetAllTopic(topics)))
}
