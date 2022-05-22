package blog

import "context"

type Repository interface {
	CreateBlog(ctx context.Context, b *Blog) (*Blog, error)
	ReadBlogByID(ctx context.Context, id int64) (*Blog, error)
	UpdateBlog(ctx context.Context, b *Blog) (*Blog, error)
	DeleteBlog(ctx context.Context, id int64) error
	ReadBlogList(ctx context.Context) ([]*Blog, error)
}
