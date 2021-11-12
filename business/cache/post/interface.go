package post

type Service interface {
	SetNewPost(postNews *CachePost) error

	GetPostById(id *string) (*CachePost, error)

	GetPostBySlug(slug *string) (*CachePost, error)

	DeletePost(id *string) error
}

type Repository interface {
	SetNewPost(postNews *CachePost) error

	GetPostById(id *string) (*CachePost, error)

	GetPostBySlug(slug *string) (*CachePost, error)

	DeletePost(id *string) error
}
