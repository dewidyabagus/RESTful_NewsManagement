package topic

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis"

	"RESTful/business/cache/topic"
)

type Repository struct {
	DB *redis.Client
}

func NewRepository(db *redis.Client) *Repository {
	return &Repository{db}
}

type Topic struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (t *Topic) toCacheTopic() *topic.CacheTopic {
	return &topic.CacheTopic{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

func toInsertTopic(t *topic.CacheTopic) *Topic {
	return &Topic{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

func (r *Repository) SetNewTopic(topic *topic.CacheTopic) error {
	jsonTopic, err := json.Marshal(toInsertTopic(topic))
	if err != nil {
		return err
	}

	return r.DB.Set("topics:"+topic.ID, jsonTopic, 0).Err()
}

func (r *Repository) GetTopicById(id *string) (*topic.CacheTopic, error) {
	var topic = new(Topic)

	jsonTopic, err := r.DB.Get("topics:" + *id).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}

	json.Unmarshal([]byte(jsonTopic), topic)

	return topic.toCacheTopic(), nil
}

func (r *Repository) DeleteTopic(id *string) error {
	return r.DB.Del("topics:" + *id).Err()
}
