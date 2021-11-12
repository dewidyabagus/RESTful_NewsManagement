package post

import (
	"strings"
	"time"

	"RESTful/business"
	cachePost "RESTful/business/cache/post"
	"RESTful/utils/validator"
)

type service struct {
	repository Repository
	cache      cachePost.Service
}

func NewService(repository Repository, cache cachePost.Service) Service {
	return &service{repository, cache}
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

	cachePostBySlug, err := s.cache.GetPostBySlug(slug)
	if err != nil {
		return nil, err
	}

	if cachePostBySlug != nil {
		return postCacheToPost(cachePostBySlug), nil
	}

	newsPost, err := s.repository.FindPostBySlug(slug)
	if err != nil {
		return nil, err
	}

	if err := s.cache.SetNewPost(toCachePost(newsPost)); err != nil {
		return nil, err
	}

	return newsPost, nil
}

func (s *service) FindPostById(id *string) (*Post, error) {
	cacheNewsPost, err := s.cache.GetPostById(id)
	if err != nil {
		return nil, err
	}

	if cacheNewsPost != nil {
		return postCacheToPost(cacheNewsPost), err
	}

	newsPost, err := s.repository.FindPostById(id)
	if err != nil {
		return nil, err
	}

	if err := s.cache.SetNewPost(toCachePost(newsPost)); err != nil {
		return nil, err
	}

	return newsPost, nil
}

func (s *service) FindPostByTopicId(topicId *string) (*[]Post, error) {
	return s.repository.FindPostByTopicId(topicId)
}

func (s *service) PublishPost(id *string) error {
	postNews, err := s.FindPostById(id)

	if err != nil {
		return err

	} else if postNews.Published {
		return business.ErrHasBeenPublished

	}

	if err := s.cache.DeletePost(id); err != nil {
		return err
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

	if err := s.cache.DeletePost(id); err != nil {
		return err
	}

	return s.repository.UpdatePost(id, post.toUpdatePost())
}

func (s *service) DeletePost(id *string) error {
	if _, err := s.repository.FindPostById(id); err != nil {
		return err
	}

	if err := s.cache.DeletePost(id); err != nil {
		return err
	}

	return s.repository.DeletePost(id, time.Now())
}
