package topic

import (
	"github.com/google/uuid"
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

func (c *Controller) FindTopicById(ctx echo.Context) error {
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	topic, err := c.service.FindTopicById(&id)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithData(response.GetOneTopic(topic)))
}

func (c *Controller) FindTopicByNameWithAllPosts(ctx echo.Context) error {
	name := ctx.Param("name")

	topic, err := c.service.FindTopicByNameWithAllPosts(&name)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithData(response.GetTopicWithAllPosts(topic)))
}

func (c *Controller) UpdateTopic(ctx echo.Context) error {
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	var topic = new(request.Topic)
	if err := ctx.Bind(topic); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	if err := c.service.UpdateTopic(&id, topic.ToBusinessTopicSpec()); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) DeleteTopic(ctx echo.Context) error {
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	if err := c.service.DeleteTopic(&id); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}
