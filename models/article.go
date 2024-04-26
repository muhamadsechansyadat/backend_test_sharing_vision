package models

import "time"

// Post Post Model
type Post struct {
	ID          int64     `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `gorm:"size:200" json:"title"`
	Content     string    `json:"content"`
	Category    string    `gorm:"size:100"`
	CreatedDate time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"created_date"`
	UpdatedDate time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"updated_date"`
	Status      string    `gorm:"type:enum('Publish', 'Draft', 'Trash');default:'Publish'"`
}

// TableName method sets table name for Post model
func (post *Post) TableName() string {
	return "posts"
}

// ResponseMap -> response map method of Post
func (post *Post) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = post.ID
	resp["title"] = post.Title
	resp["content"] = post.Content
	resp["category"] = post.Category
	resp["created_at"] = post.CreatedDate
	resp["updated_at"] = post.UpdatedDate
	resp["status"] = post.Status
	return resp
}
