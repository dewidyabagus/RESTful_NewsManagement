package post

import (
	"RESTful/business"
	"RESTful/utils/validator"
	"strings"
	"time"
)

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) InsertPost(post *PostSpec) error {
	if err := validator.GetValidator().Struct(post); err != nil {
		return business.ErrDataNotSpec
	}

	_, err := s.repository.FindPostBySlug(&post.Slug)
	if err == nil {
		return business.ErrDataConflict

	} else if err != business.ErrDataNotFound {
		return err

	}

	return s.repository.InsertPost(post.toInsertPost())
}

func (s *service) FindAllPost(status *string) (*[]Post, error) {
	*status = strings.ToLower(*status)

	if *status != "draft" && *status != "deleted" && *status != "publish" && *status != "" {
		return nil, business.ErrBadRequest
	}

	return s.repository.FindAllPost(status)
}

func (s *service) FindPostBySlug(slug *string) (*Post, error) {
	if strings.TrimSpace(*slug) == "" {
		return nil, business.ErrBadRequest
	}

	return s.repository.FindPostBySlug(slug)
}

func (s *service) FindPostById(id *string) (*Post, error) {
	return s.repository.FindPostById(id)
}

func (s *service) FindPostByTopicId(topicId *string) (*[]Post, error) {
	return s.repository.FindPostByTopicId(topicId)
}

func (s *service) PublishPost(id *string) error {
	postNews, err := s.repository.FindPostById(id)

	if err != nil {
		return err

	} else if postNews.Published {
		return business.ErrHasBeenPublished

	}

	return s.repository.PublishPost(id, time.Now())
}

func (s *service) UpdatePost(id *string, post *PostSpec) error {
	if err := validator.GetValidator().Struct(post); err != nil {
		return business.ErrDataNotSpec
	}

	if _, err := s.repository.FindPostById(id); err != nil {
		return err
	}

	return s.repository.UpdatePost(id, post.toUpdatePost())
}

func (s *service) DeletePost(id *string) error {
	if _, err := s.repository.FindPostById(id); err != nil {
		return err
	}

	return s.repository.DeletePost(id, time.Now())
}
