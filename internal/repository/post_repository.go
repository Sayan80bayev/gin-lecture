package repository

import (
	"gin-lecture/internal/model"
	"gorm.io/gorm"
)

type PostRepositoryImpl struct {
	db *gorm.DB
}

// NewPostRepository Constructor
func NewPostRepository(db *gorm.DB) *PostRepositoryImpl {
	return &PostRepositoryImpl{db: db}
}

func (p PostRepositoryImpl) GetAll() ([]model.Post, error) {
	var posts []model.Post
	err := p.db.Find(&posts).Error
	return posts, err
}

func (p PostRepositoryImpl) GetById(id int) (*model.Post, error) {
	var post model.Post
	err := p.db.First(&post, id).Error
	return &post, err
}

func (p PostRepositoryImpl) Create(post *model.Post) error {
	return p.db.Create(post).Error
}

func (p PostRepositoryImpl) Update(post *model.Post, id int) error {
	return p.db.Model(&model.Post{}).Where("id = ?", id).Omit("id, CreatedAt").Updates(post).Error
}

func (p PostRepositoryImpl) Delete(postID int) error {
	return p.db.Delete(&model.Post{}, postID).Error
}
