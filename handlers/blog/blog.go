package blog

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/afikrim/protobuf-tutorial/blog"
	pbBlog "github.com/afikrim/protobuf-tutorial/pb/blog"
	"google.golang.org/protobuf/types/known/emptypb"
)

type handler struct {
	blogSvc blog.Service
}

func NewHandler(blogSvc blog.Service) *handler {
	return &handler{
		blogSvc: blogSvc,
	}
}

func (h *handler) CreateBlog(ctx context.Context, req *pbBlog.CreateBlogRequest) (*pbBlog.Blog, error) {
	b := parsePbCreateBlogRequestDataToBlog(req.Data)
	b, err := h.blogSvc.CreateBlog(ctx, b)

	if err != nil {
		return nil, err
	}

	return parseBlogToPbBlog(b), nil
}

func (h *handler) ReadBlogByID(ctx context.Context, req *pbBlog.ReadBlogByIDRequest) (*pbBlog.Blog, error) {
	b, err := h.blogSvc.ReadBlogByID(ctx, req.Id)

	if err != nil {
		return nil, err
	}

	return parseBlogToPbBlog(b), nil
}

func (h *handler) UpdateBlog(ctx context.Context, req *pbBlog.UpdateBlogRequest) (*pbBlog.Blog, error) {
	if req.Data.Id == 0 {
		req.Data.Id = req.Id
	}

	b := parsePbUpdateBlogRequestDataToBlog(req.Data)
	b, err := h.blogSvc.UpdateBlog(ctx, b)

	if err != nil {
		return nil, err
	}

	return parseBlogToPbBlog(b), nil
}

func (h *handler) DeleteBlog(ctx context.Context, req *pbBlog.DeleteBlogRequest) (*emptypb.Empty, error) {
	err := h.blogSvc.DeleteBlog(ctx, req.Id)

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h *handler) ReadBlogList(ctx context.Context, req *pbBlog.ReadBlogListRequest) (*pbBlog.ReadBlogListResponse, error) {
	b, err := h.blogSvc.ReadBlogList(ctx)

	if err != nil {
		return nil, err
	}

	return &pbBlog.ReadBlogListResponse{
		Blogs: parseBlogsToPbBlogs(b),
	}, nil
}

func parsePbBlogToBlog(pbBlog *pbBlog.Blog) *blog.Blog {
	return &blog.Blog{
		ID:        pbBlog.Id,
		Title:     pbBlog.Title,
		Slug:      pbBlog.Slug,
		Content:   pbBlog.Content,
		Author:    pbBlog.Author,
		CreatedAt: pbBlog.CreatedAt,
		UpdatedAt: pbBlog.UpdatedAt,
	}
}

func parsePbCreateBlogRequestDataToBlog(pbBlog *pbBlog.CreateBlogRequestData) *blog.Blog {
	return &blog.Blog{
		Title:   pbBlog.Title,
		Slug:    fmt.Sprintf("%d--%s", time.Now().Unix(), strings.Replace(pbBlog.Title, " ", "-", -1)),
		Content: pbBlog.Content,
		Author:  pbBlog.Author,
	}
}

func parsePbUpdateBlogRequestDataToBlog(pbBlog *pbBlog.UpdateBlogRequestData) *blog.Blog {
	return &blog.Blog{
		ID:      pbBlog.Id,
		Title:   pbBlog.Title,
		Slug:    fmt.Sprintf("%d--%s", pbBlog.Id, strings.Replace(pbBlog.Title, " ", "-", -1)),
		Content: pbBlog.Content,
		Author:  pbBlog.Author,
	}
}

func parseBlogToPbBlog(blog *blog.Blog) *pbBlog.Blog {
	return &pbBlog.Blog{
		Id:        blog.ID,
		Title:     blog.Title,
		Slug:      blog.Slug,
		Content:   blog.Content,
		Author:    blog.Author,
		CreatedAt: blog.CreatedAt,
		UpdatedAt: blog.UpdatedAt,
	}
}

func parseBlogsToPbBlogs(blogs []*blog.Blog) []*pbBlog.Blog {
	pbBlogs := make([]*pbBlog.Blog, len(blogs))
	for i, blog := range blogs {
		pbBlogs[i] = parseBlogToPbBlog(blog)
	}
	return pbBlogs
}
