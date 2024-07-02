package dto

type PostDAO struct {
	ID        uint
	Title     string
	Content   string
	Image     string
	Category  string
	Owner     string
	CreatedAt string
	UpdatedAt string
	Comments  []uint
	Likes     uint
	Dislikes  uint
}
