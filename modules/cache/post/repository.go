package post

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis"

	"RESTful/business/cache/post"
)

type Repository struct {
	DB *redis.Client
}

func NewRepository(db *redis.Client) *Repository {
	return &Repository{db}
}

type Post struct {
	ID              string    `json:"id"`
	TopicId         string    `json:"topic_id"`
	TopicName       string    `json:"topic_name"`
	TopicDesciption string    `json:"topic_description"`
	Title           string    `json:"title"`
	Slug            string    `json:"slug"`
	Excerpt         string    `json:"excerpt"`
	Body            string    `json:"body"`
	Tags            []string  `json:"tags"`
	Published       bool      `json:"published"`
	PublishedAt     time.Time `json:"published_at"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (p *Post) toCachePost() *post.CachePost {
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

func toInsertPost(p *post.CachePost) *Post {
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

func (r *Repository) SetNewPost(newsPost *post.CachePost) error {
	jsonPost, err := json.Marshal(toInsertPost(newsPost))
	if err != nil {
		return err
	}

	err = r.DB.HSet("slugs", newsPost.Slug, newsPost.ID).Err()
	if err != nil {
		return err
	}

	return r.DB.Set("posts:"+newsPost.ID, jsonPost, 0).Err()
}

func (r *Repository) GetPostById(id *string) (*post.CachePost, error) {
	var newsPost = new(Post)

	jsonPost, err := r.DB.Get("posts:" + *id).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}

	json.Unmarshal([]byte(jsonPost), newsPost)

	return newsPost.toCachePost(), nil
}

func (r *Repository) GetPostBySlug(slug *string) (*post.CachePost, error) {
	cachePostId, err := r.DB.HGet("slugs", *slug).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
	}

	return r.GetPostById(&cachePostId)
}

func (r *Repository) DeletePost(id *string) error {
	newsPost, err := r.GetPostById(id)
	if err != nil {
		return err
	}

	if err := r.DB.Del("posts:" + newsPost.ID).Err(); err != nil {
		return err
	}

	return r.DB.HDel("slugs", newsPost.Slug).Err()
}
