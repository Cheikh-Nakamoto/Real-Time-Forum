package models

import "github.com/google/uuid"

type Action struct {
	id        uint
	like      bool
	dislike   bool
	postId    uint
	commentId uint
	userId    uuid.UUID
}

// GetID returns the id of the action
func (a *Action) GetID() uint {
	return a.id
}

// SetID sets the id of the action
func (a *Action) SetID(id uint) {
	a.id = id
}

// GetLike returns the like of the action
func (a *Action) GetLike() bool {
	return a.like
}

// SetLike sets the like of the action
func (a *Action) SetLike(like bool) {
	a.like = like
}

// GetDislike returns the dislike of the action
func (a *Action) GetDislike() bool {
	return a.dislike
}

// SetDislike sets the dislike of the action
func (a *Action) SetDislike(dislike bool) {
	a.dislike = dislike
}

// GetPostID returns the type of content to like
func (a *Action) GetPostID() uint {
	return a.postId
}

// SetPostID sets the type of content to like
func (a *Action) SetPostID(postId uint) {
	a.postId = postId
}

// GetCommentID returns the type of content to like
func (a *Action) GetCommentID() uint {
	return a.commentId
}

// SetCommentID sets the type of content to like
func (a *Action) SetCommentID(commentId uint) {
	a.commentId = commentId
}

// GetUserID returns the id of the user who liked it
func (a *Action) GetUserID() uuid.UUID {
	return a.userId
}

// SetUserID sets the id of the user who liked it
func (a *Action) SetUserID(userId uuid.UUID) {
	a.userId = userId
}
