package blog

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type repository struct {
	db        *gorm.DB
	tableName string
}

func NewRepository(db *gorm.DB, tableName string) Repository {
	return &repository{
		db:        db,
		tableName: tableName,
	}
}

func (r *repository) CreateBlog(ctx context.Context, b *Blog) (*Blog, error) {
	b.CreatedAt = time.Now().Unix()
	b.UpdatedAt = time.Now().Unix()

	err := r.db.WithContext(ctx).Create(&b).Error

	if err != nil {
		return nil, err
	}

	return b, nil
}

func (r *repository) ReadBlogByID(ctx context.Context, id int64) (*Blog, error) {
	var b Blog
	err := r.db.WithContext(ctx).First(&b, &Blog{ID: id}).Error

	if err != nil {
		return nil, err
	}

	return &b, nil
}

func (r *repository) UpdateBlog(ctx context.Context, b *Blog) (*Blog, error) {
	b.UpdatedAt = time.Now().Unix()

	err := r.db.WithContext(ctx).Where(&Blog{ID: b.ID}).Updates(&b).Error

	if err != nil {
		return nil, err
	}

	return b, nil
}

func (r *repository) DeleteBlog(ctx context.Context, id int64) error {
	err := r.db.WithContext(ctx).Delete(&Blog{ID: id}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) ReadBlogList(ctx context.Context) ([]*Blog, error) {
	var b []*Blog
	err := r.db.WithContext(ctx).Find(&b).Error

	if err != nil {
		return nil, err
	}

	return b, nil
}
