package models

type CreateUserRequest struct {
	Name   string `json:"name" binding:"required"`
	Email  string `json:"email" binding:"required"`
	Gender string `json:"gender" binding:"required"`
	Image  string `json:"image" binding:"required"`
}
