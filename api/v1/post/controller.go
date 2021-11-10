package post

import (
	"RESTful/api/common"
	"RESTful/api/v1/post/request"
	"RESTful/api/v1/post/response"
	"RESTful/business/post"

	echo "github.com/labstack/echo/v4"
)

type Controller struct {
	service post.Service
}

func NewController(service post.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) InsertPost(ctx echo.Context) error {
	var post = new(request.Post)

	if err := ctx.Bind(post); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	err := c.service.InsertPost(post.ToBusinessPostSpec())
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) FindAllPost(ctx echo.Context) error {
	posts, err := c.service.FindAllPost()
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithData(response.GetAllPostSummary(posts)))
}
