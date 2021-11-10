package post

type Service interface {
	InsertPost(post *PostSpec) error

	FindAllPost() (*[]Post, error)
}

type Repository interface {
	InsertPost(post *Post) error

	FindPostBySlug(slug *string) (*Post, error)

	FindAllPost() (*[]Post, error)
}
