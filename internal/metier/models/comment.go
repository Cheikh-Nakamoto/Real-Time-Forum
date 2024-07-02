package models

import (
	"bim/internal/pkg"
	"github.com/google/uuid"
	"html"
	"time"
)

type Comment struct {
	id        uint
	content   string
	createdAt time.Time
	updatedAt time.Time
	postId    uint
	userId    uuid.UUID
	owner     string
	like      uint
	dislike   uint
}

// GetID returns the id of the comment
func (c *Comment) GetID() uint {
	return c.id
}

// SetID sets the id of the comment
func (c *Comment) SetID(id uint) {
	c.id = id
}

// GetContent returns the content of the comment
func (c *Comment) GetContent() string {
	return html.EscapeString(c.content)
}

// SetContent sets the content of the comment
func (c *Comment) SetContent(content string) {
	c.content = content
}

// GetCreatedAt returns the created time of the comment
func (c *Comment) GetCreatedAt() string {
	return pkg.FormatDateTime(c.createdAt)
}

// SetCreatedAt sets the created time of the comment
func (c *Comment) SetCreatedAt(date time.Time) {
	c.createdAt = date
}

// GetUpdatedAt returns the updated time of the comment
func (c *Comment) GetUpdatedAt() string {
	return pkg.FormatDateTime(c.updatedAt)
}

// SetUpdatedAt sets the updated time of the comment
func (c *Comment) SetUpdatedAt(date time.Time) {
	c.updatedAt = date
}

// GetUserId returns the id of the user who created the comment
func (c *Comment) GetUserId() uuid.UUID {
	return c.userId
}

// SetUserId sets the id of the user who created the comment
func (c *Comment) SetUserId(userId uuid.UUID) {
	c.userId = userId
}

// GetPostId returns the id of the post that was commented
func (c *Comment) GetPostId() uint {
	return c.postId
}

// SetPostId sets the id of the post that was commented
func (c *Comment) SetPostId(postId uint) {
	c.postId = postId
}

func (c *Comment) GetOwner() string {
	return c.owner
}

func (c *Comment) SetOwner(owner string) {
	c.owner = owner
}

func (c *Comment) GetLike() uint {
	return c.like
}

func (c *Comment) SetLike(like uint) {
	c.like = like
}

func (c *Comment) GetDislike() uint {
	return c.dislike
}

func (c *Comment) SetDislike(dislike uint) {
	c.dislike = dislike
}
