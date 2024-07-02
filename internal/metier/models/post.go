package models

import (
	"bim/internal/pkg"
	"html"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	id         uint
	title      string
	content    string
	image      string // image : "https://cdn.pixabay.com/photo/2017/02/21/05/59/label-2084756_1280.jpg"
	createdAt  time.Time
	updatedAt  time.Time
	userId     uuid.UUID
	owner      string
	categories []Category
	comments   []uint
	like       uint
	dislike    uint
	comment    uint
}

// GetID returns the id of the post
func (p *Post) GetID() uint {
	return p.id
}

// SetID sets the id of the post
func (p *Post) SetID(id uint) {
	p.id = id
}

// GetTitle returns the title of the post
func (p *Post) GetTitle() string {
	return html.EscapeString(p.title)
}

// SetTitle sets the title of the post
func (p *Post) SetTitle(title string) {
	p.title = title
}

// GetContent returns the content of the post
func (p *Post) GetContent() string {
	return html.EscapeString(p.content)
}

// SetContent sets the content of the post
func (p *Post) SetContent(content string) {
	p.content = content
}

// GetImage returns the image of the post
func (p *Post) GetImage() string {
	return p.image
}

// SetImage sets the image of the post
func (p *Post) SetImage(image string) {
	p.image = image
}

// GetCreatedAt returns the created time of the post
func (p *Post) GetCreatedAt() string {
	return pkg.FormatDateTime(p.createdAt)
}

// SetCreatedAt sets the created time of the post
func (p *Post) SetCreatedAt(createdAt time.Time) {
	p.createdAt = createdAt
}

// GetUpdatedAt returns the updated time of the post
func (p *Post) GetUpdatedAt() string {
	return pkg.FormatDateTime(p.updatedAt)
}

// SetUpdatedAt sets the updated time of the post
func (p *Post) SetUpdatedAt(updateAt time.Time) {
	p.updatedAt = updateAt
}

// GetUserId returns the id of the user who created the post
func (p *Post) GetUserId() uuid.UUID {
	return p.userId
}

// SetUserId sets the id of the user who created the post
func (p *Post) SetUserId(userId uuid.UUID) {
	p.userId = userId
}

// GetOwner returns the user who created the post
func (p *Post) GetOwner() string {
	return p.owner
}

// SetOwner sets the user who created the post
func (p *Post) SetOwner(owner string) {
	p.owner = owner
}
func (p *Post) GetCategories() []Category {
	return p.categories
}
func (p *Post) SetCategories(categories []Category) {
	p.categories = categories
}
func (p *Post) GetCommentsid() []uint {
	return p.comments
}

func (p *Post) SetCommentsId(comments []uint) {
	p.comments = comments
}
func (p *Post) GetLike() uint {
	return p.like
}
func (p *Post) SetLike(like uint) {
	p.like = like
}
func (p *Post) GetDislike() uint {
	return p.dislike
}
func (p *Post) SetDislike(dislike uint) {
	p.dislike = dislike
}
func (p *Post) GetComment() uint {
	return p.comment
}
func (p *Post) SetComment(comment uint) {
	p.comment = comment
}
