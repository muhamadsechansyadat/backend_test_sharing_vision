package service

import (
	"backend_test_sharing_vision/api/repository"
	"backend_test_sharing_vision/models"
	"time"
)

// PostService PostService struct
type PostService struct {
	repository repository.PostRepository
}

// NewPostService : returns the PostService struct instance
func NewPostService(r repository.PostRepository) PostService {
	return PostService{
		repository: r,
	}
}

// Save -> calls post repository save method
func (p PostService) Save(post models.Post) error {
	return p.repository.Save(post)
}

// FindAll -> calls post repo find all method
func (p PostService) FindAll(post models.Post, limit int, offset int, status string) (*[]models.Post, int64, error) {
	return p.repository.FindAll(post, limit, offset, status)
}

// Update -> calls postrepo update method
func (p PostService) Update(post models.Post) error {
	post.UpdatedDate = time.Now()
	return p.repository.Update(post)
}

// Delete -> calls post repo delete method
func (p PostService) Delete(id int64) error {
	var post models.Post
	post.ID = id
	return p.repository.Trash(post)
}

// Find -> calls post repo find method
func (p PostService) Find(post models.Post) (models.Post, error) {
	return p.repository.Find(post)
}
