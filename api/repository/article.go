package repository

import (
	"backend_test_sharing_vision/infrastructure"
	"backend_test_sharing_vision/models"
)

// PostRepository -> PostRepository
type PostRepository struct {
	db infrastructure.Database
}

// NewPostRepository : fetching database
func NewPostRepository(db infrastructure.Database) PostRepository {
	return PostRepository{
		db: db,
	}
}

// Save -> Method for saving post to database
func (p PostRepository) Save(post models.Post) error {
	return p.db.DB.Create(&post).Error
}

// FindAll -> Method for fetching all posts from database
func (p PostRepository) FindAll(post models.Post, limit int, offset int, status string) (*[]models.Post, int64, error) {
	var posts []models.Post
	var totalRows int64 = 0

	queryBuilder := p.db.DB.Order("created_date desc").Model(&models.Post{})

	if status != "" {
		queryKeyword := "%" + status + "%"
		queryBuilder = queryBuilder.Where(
			p.db.DB.Where("posts.status LIKE ? ", queryKeyword))
	}

	queryBuilder = queryBuilder.Limit(limit).Offset(offset)

	err := queryBuilder.
		Where(post).
		Find(&posts).
		Count(&totalRows).Error
	return &posts, totalRows, err
}

// Update -> Method for updating Post
func (p PostRepository) Update(post models.Post) error {
	return p.db.DB.Save(&post).Error
}

// Find -> Method for fetching post by id
func (p PostRepository) Find(post models.Post) (models.Post, error) {
	var posts models.Post
	err := p.db.DB.
		Debug().
		Model(&models.Post{}).
		Where(&post).
		Take(&posts).Error
	return posts, err
}

// Delete Deletes Post
func (p PostRepository) Trash(post models.Post) error {
	return p.db.DB.Model(&post).Update("status", "Trash").Error
}
