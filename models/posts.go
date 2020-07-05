package models

type CreatePostRequest struct {
	Description string `json:"description"`
	Content     string `json:"content" binding:"required"`
	Author      string `json:"author" binding:"required"`
}

type Author struct {
	Name   string `json:"name"`
	Image  string `json:"image"`
	Gender string `json:"gender"`
	Email  string `json:"email"`
}

type Post struct {
	Content     string `json:"content"`
	Description string `json:"description"`
	Timestamp   string `json:"postedAt"`
}
