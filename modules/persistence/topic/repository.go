package topic

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"RESTful/business"
	"RESTful/business/topic"
)

type Topic struct {
	ID          string    `gorm:"id;type:uuid;primaryKey"`
	Name        string    `gorm:"name;type:varchar(75);unique:topics_name_uniq;not null"`
	Description string    `gorm:"description;type:varchar(150);not null"`
	CreatedAt   time.Time `gorm:"created_at;type:timestamp;not null"`
	UpdatedAt   time.Time `gorm:"updated_at;type:timestamp;not null"`
	DeletedAt   time.Time `gorm:"deleted_at;type:timestamp"`
}

func (t *Topic) toBusinessTopic() *topic.Topic {
	return &topic.Topic{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

func toInsertTopic(t *topic.Topic) *Topic {
	return &Topic{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

func toAllBusinessTopic(topics *[]Topic) *[]topic.Topic {
	var response = []topic.Topic{}

	for _, item := range *topics {
		response = append(response, *item.toBusinessTopic())
	}

	return &response
}

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) InsertTopic(data *topic.Topic) error {
	return r.DB.Create(toInsertTopic(data)).Error
}

func (r *Repository) FindTopicByName(name *string) (*topic.Topic, error) {
	var data = new(Topic)

	if err := r.DB.First(data, "name = ?", name).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, business.ErrDataNotFound
		}
		return nil, err
	}

	return data.toBusinessTopic(), nil
}

func (r *Repository) FindAllTopic() (*[]topic.Topic, error) {
	var data = new([]Topic)

	err := r.DB.Find(data).Where("to_char(deleted_at, 'YYYY') = '0001' or deleted_at is null").Order("name asc").Error
	if err != nil {
		return nil, err
	}

	return toAllBusinessTopic(data), nil
}
