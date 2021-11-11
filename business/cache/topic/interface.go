package topic

type Service interface {
	SetNewTopic(topic *CacheTopic) error

	GetTopicById(id *string) (*CacheTopic, error)

	DeleteTopic(id *string) error
}

type Repository interface {
	SetNewTopic(topic *CacheTopic) error

	GetTopicById(id *string) (*CacheTopic, error)

	DeleteTopic(id *string) error
}
