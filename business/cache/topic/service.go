package topic

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) SetNewTopic(topic *CacheTopic) error {
	return s.repository.SetNewTopic(topic)
}

func (s *service) GetTopicById(id *string) (*CacheTopic, error) {
	return s.repository.GetTopicById(id)
}

func (s *service) DeleteTopic(id *string) error {
	return s.repository.DeleteTopic(id)
}
