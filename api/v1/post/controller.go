package post

import (
	"RESTful/api/common"
	"RESTful/api/v1/post/request"
	"RESTful/api/v1/post/response"
	"RESTful/business/post"

	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"
)

type Controller struct {
	service post.Service
}

func NewController(service post.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) InsertPost(ctx echo.Context) error {
	var postNews = new(request.Post)

	if err := ctx.Bind(postNews); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	err := c.service.InsertPost(postNews.ToBusinessPostSpec())
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

	postNews, err := c.service.FindPostBySlug(&slug)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithData(response.GetOnePostDetail(postNews)))
}

func (c *Controller) FindPostById(ctx echo.Context) error {
	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	postNews, err := c.service.FindPostById(&id)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithData(response.GetOnePostDetail(postNews)))
}

func (c *Controller) PublishPost(ctx echo.Context) error {
	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	if err := c.service.PublishPost(&id); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) UpdatePost(ctx echo.Context) error {
	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	var postNews = new(request.Post)
	if err := ctx.Bind(postNews); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	if err := c.service.UpdatePost(&id, postNews.ToBusinessPostSpec()); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) DeletePost(ctx echo.Context) error {
	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	if err := c.service.DeletePost(&id); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}
