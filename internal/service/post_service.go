package service

import (
	"gin-lecture/internal/model"
)

// Open interfaces where you need them
type PostRepository interface {
	GetAll() ([]model.Post, error)
	GetById(id int) (*model.Post, error)
	Create(post *model.Post) error
	Update(post *model.Post, id int) error
	Delete(postID int) error
}

type PostService struct {
	repo PostRepository
}

// NewPostService Constructor
func NewPostService(postRepo PostRepository) *PostService {
	return &PostService{repo: postRepo}
}

func (p *PostService) GetAllPosts() ([]model.Post, error) {
	posts, err := p.repo.GetAll()
	return posts, err
}

func (p *PostService) GetPostByID(id int) (*model.Post, error) {
	post, err := p.repo.GetById(id)
	return post, err
}

func (p *PostService) Create(title, description, content string) (*model.Post, error) {
	post := &model.Post{
		Title:       title,
		Description: description,
		Content:     content,
	}
	err := p.repo.Create(post)
	return post, err
}

func (p *PostService) Update(title, description, content string, id int) error {
	post := &model.Post{
		Title:       title,
		Description: description,
		Content:     content,
	}
	err := p.repo.Update(post, id)
	return err
}

func (p *PostService) Delete(postID int) error {
	return p.repo.Delete(postID)
}
