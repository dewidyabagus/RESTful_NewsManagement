package post

import "time"

type CachePost struct {
	ID              string
	TopicId         string
	TopicName       string
	TopicDesciption string
	Title           string
	Slug            string
	Excerpt         string
	Body            string
	Tags            []string
	Published       bool
	PublishedAt     time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
