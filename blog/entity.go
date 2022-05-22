package blog

import (
	"errors"
)

var (
	ErrBlogTitleEmpty   = errors.New("blog title is empty")
	ErrBlogContentEmpty = errors.New("blog content is empty")
	ErrBlogAuthorEmpty  = errors.New("blog author is empty")
)

type Blog struct {
	ID        int64  `json:"id" gorm:"column:id;autoIncrement;primaryKey"`
	Title     string `json:"title" gorm:"column:title"`
	Slug      string `json:"slug" gorm:"column:slug"`
	Content   string `json:"content" gorm:"column:content;type:text"`
	Author    string `json:"author" gorm:"column:author"`
	CreatedAt int64  `json:"created_at" gorm:"column:created_at;type:bigint"`
	UpdatedAt int64  `json:"updated_at" gorm:"column:updated_at;type:bigint"`
}

func (b *Blog) validate() error {
	if b.Title == "" {
		return ErrBlogTitleEmpty
	}
	if b.Content == "" {
		return ErrBlogContentEmpty
	}
	if b.Author == "" {
		return ErrBlogAuthorEmpty
	}
	return nil
}
