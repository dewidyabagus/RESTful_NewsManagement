package topic

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

// func (s *service) FindTopicById(id *string) (*Topic, error) {

// }

func (s *service) FindAllTopic() (*[]Topic, error) {
	return s.repository.FindAllTopic()
}

// func (s *service) UpdateTopic(id *string, topic *TopicSpec) error {

// }

// func (s *service) DeleteTopic(id *string) error {

// }
