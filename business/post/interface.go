package post

type Service interface {
	InsertPost(post *PostSpec) error

	FindPostBySlug(slug *string) (*Post, error)

	FindPostByTopicId(topicId *string) (*[]Post, error)

	FindAllPost(status *string) (*[]Post, error)
}

type Repository interface {
	InsertPost(post *Post) error

	FindPostBySlug(slug *string) (*Post, error)

	FindPostByTopicId(topicId *string) (*[]Post, error)

	FindAllPost(status *string) (*[]Post, error)
}
