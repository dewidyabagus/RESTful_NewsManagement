package topic

import (
	"RESTful/business"
	"RESTful/utils/validator"
	"time"
)

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
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
	return s.repository.FindTopicById(id)
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

	return s.repository.UpdateTopic(id, topic.toUpdateTopic())
}

func (s *service) DeleteTopic(id *string) error {
	if _, err := s.repository.FindTopicById(id); err != nil {
		return err
	}

	return s.repository.DeleteTopic(id, time.Now())
}
