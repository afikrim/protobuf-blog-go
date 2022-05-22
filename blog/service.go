package blog

import "context"

type Service interface {
	CreateBlog(ctx context.Context, b *Blog) (*Blog, error)
	ReadBlogByID(ctx context.Context, id int64) (*Blog, error)
	UpdateBlog(ctx context.Context, b *Blog) (*Blog, error)
	DeleteBlog(ctx context.Context, id int64) error
	ReadBlogList(ctx context.Context) ([]*Blog, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateBlog(ctx context.Context, b *Blog) (*Blog, error) {
	if err := b.validate(); err != nil {
		return nil, err
	}
	return s.repo.CreateBlog(ctx, b)
}

func (s *service) ReadBlogByID(ctx context.Context, id int64) (*Blog, error) {
	return s.repo.ReadBlogByID(ctx, id)
}

func (s *service) UpdateBlog(ctx context.Context, b *Blog) (*Blog, error) {
	if err := b.validate(); err != nil {
		return nil, err
	}
	return s.repo.UpdateBlog(ctx, b)
}

func (s *service) DeleteBlog(ctx context.Context, id int64) error {
	return s.repo.DeleteBlog(ctx, id)
}

func (s *service) ReadBlogList(ctx context.Context) ([]*Blog, error) {
	return s.repo.ReadBlogList(ctx)
}
