package models

type FollowRequest struct {
	Follower  string `json:"follower" binding:"required"`
	Following string `json:"following" binding:"required"`
}
