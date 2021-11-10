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
	status := ctx.QueryParam("status")

	posts, err := c.service.FindAllPost(&status)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithData(response.GetAllPostSummary(posts)))
}

func (c *Controller) FindPostBySlug(ctx echo.Context) error {
	slug := ctx.Param("slug")

	post, err := c.service.FindPostBySlug(&slug)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithData(response.GetOnePostDetail(post)))
}
