package post

import "time"

type Service interface {
	InsertPost(post *PostSpec) error

	FindPostBySlug(slug *string) (*Post, error)

	FindPostById(id *string) (*Post, error)

	FindPostByTopicId(topicId *string) (*[]Post, error)

	FindAllPost(status *string) (*[]Post, error)

	PublishPost(id *string) error

	// UpdatePost(id *string, post *PostSpec) error

	// DeletePost(id *string) error
}

type Repository interface {
	InsertPost(post *Post) error

	FindPostBySlug(slug *string) (*Post, error)

	FindPostById(id *string) (*Post, error)

	FindPostByTopicId(topicId *string) (*[]Post, error)

	FindAllPost(status *string) (*[]Post, error)

	PublishPost(id *string, published time.Time) error

	// UpdatePost(id *string, post *Post) error

	// DeletePost(id *string) error
}
