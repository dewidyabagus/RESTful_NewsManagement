package post

import (
	"time"

	"github.com/google/uuid"

	"RESTful/business/cache/post"
)

type Post struct {
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

type PostSpec struct {
	TopicId string   `validate:"required"`
	Title   string   `validate:"required"`
	Slug    string   `validate:"required"`
	Excerpt string   `validate:"required"`
	Body    string   `validate:"required"`
	Tags    []string `validate:"required"`
}

func (p *PostSpec) toInsertPost() *Post {
	return &Post{
		ID:        uuid.New().String(),
		TopicId:   p.TopicId,
		Title:     p.Title,
		Slug:      p.Slug,
		Excerpt:   p.Excerpt,
		Body:      p.Body,
		Tags:      p.Tags,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (p *PostSpec) toUpdatePost() *Post {
	return &Post{
		TopicId:   p.TopicId,
		Title:     p.Title,
		Slug:      p.Slug,
		Excerpt:   p.Excerpt,
		Body:      p.Body,
		Tags:      p.Tags,
		UpdatedAt: time.Now(),
	}
}

func postCacheToPost(p *post.CachePost) *Post {
	return &Post{
		ID:              p.ID,
		TopicId:         p.TopicId,
		TopicName:       p.TopicName,
		TopicDesciption: p.TopicDesciption,
		Title:           p.Title,
		Slug:            p.Slug,
		Excerpt:         p.Excerpt,
		Body:            p.Body,
		Tags:            p.Tags,
		Published:       p.Published,
		PublishedAt:     p.PublishedAt,
		CreatedAt:       p.CreatedAt,
		UpdatedAt:       p.UpdatedAt,
	}
}

func toCachePost(p *Post) *post.CachePost {
	return &post.CachePost{
		ID:              p.ID,
		TopicId:         p.TopicId,
		TopicName:       p.TopicName,
		TopicDesciption: p.TopicDesciption,
		Title:           p.Title,
		Slug:            p.Slug,
		Excerpt:         p.Excerpt,
		Body:            p.Body,
		Tags:            p.Tags,
		Published:       p.Published,
		PublishedAt:     p.PublishedAt,
		CreatedAt:       p.CreatedAt,
		UpdatedAt:       p.UpdatedAt,
	}
}
