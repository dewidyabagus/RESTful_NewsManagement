package post

import (
	"RESTful/business"
	"RESTful/utils/validator"
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

func (s *service) FindAllPost() (*[]Post, error) {
	return s.repository.FindAllPost()
}
