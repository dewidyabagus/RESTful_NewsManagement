package topic

import (
	"strings"
	"time"

	"RESTful/business"
	topicCache "RESTful/business/cache/topic"
	"RESTful/business/post"
	"RESTful/utils/validator"
)

type service struct {
	repository Repository
	post       post.Service
	cache      topicCache.Service
}

func NewService(repository Repository, post post.Service, topic topicCache.Service) Service {
	return &service{repository, post, topic}
}

func (s *service) InsertTopic(topic *TopicSpec) error {
	if err := validator.GetValidator().Struct(topic); err != nil {
		return business.ErrDataNotSpec
	}

	_, err := s.repository.FindTopicByName(&topic.Name)

	if err == nil {
		return business.ErrDataConflict

	} else if err != business.ErrDataNotFound {
		return err

	}

	return s.repository.InsertTopic(topic.toInsertTopic())
}

func (s *service) FindTopicById(id *string) (*Topic, error) {
	resultCache, err := s.cache.GetTopicById(id)
	if err != nil {
		return nil, err
	}

	if resultCache != nil {
		return cacheToTopic(resultCache), nil
	}

	resultTopic, err := s.repository.FindTopicById(id)
	if err != nil {
		return nil, err
	}

	if err := s.cache.SetNewTopic(toCacheTopic(resultTopic)); err != nil {
		return nil, err
	}

	return resultTopic, nil
}

func (s *service) FindTopicByNameWithAllPosts(name *string) (*TopicWithPosts, error) {
	if strings.TrimSpace(*name) == "" {
		return nil, business.ErrBadRequest
	}

	result, err := s.repository.FindTopicByName(name)
	if err != nil {
		return nil, err
	}

	posts, err := s.post.FindPostByTopicId(&result.ID)
	if err != nil {
		return nil, err
	}

	return GetTopicWithAllPosts(result, posts), nil
}

func (s *service) FindAllTopic() (*[]Topic, error) {
	return s.repository.FindAllTopic()
}

func (s *service) UpdateTopic(id *string, topic *TopicSpec) error {
	if err := validator.GetValidator().Struct(topic); err != nil {
		return business.ErrDataNotSpec
	}

	if _, err := s.repository.FindTopicById(id); err != nil {
		return err
	}

	result, err := s.repository.FindTopicByName(&topic.Name)
	if err == nil {
		if result.ID != *id {
			return business.ErrDataConflict
		}
	} else if err != business.ErrDataNotFound {
		return err
	}

	if err := s.cache.DeleteTopic(id); err != nil {
		return err
	}

	return s.repository.UpdateTopic(id, topic.toUpdateTopic())
}

func (s *service) DeleteTopic(id *string) error {
	if _, err := s.repository.FindTopicById(id); err != nil {
		return err
	}

	if err := s.cache.DeleteTopic(id); err != nil {
		return err
	}

	return s.repository.DeleteTopic(id, time.Now())
}
