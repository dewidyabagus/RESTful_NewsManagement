package post

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) SetNewPost(postNews *CachePost) error {
	return s.repository.SetNewPost(postNews)
}

func (s *service) GetPostById(id *string) (*CachePost, error) {
	return s.repository.GetPostById(id)
}

func (s *service) GetPostBySlug(slug *string) (*CachePost, error) {
	return s.repository.GetPostBySlug(slug)
}

func (s *service) DeletePost(id *string) error {
	return s.repository.DeletePost(id)
}
