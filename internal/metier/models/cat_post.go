package models

type CategoryPost struct {
	categoryId uint
	postId     uint
}

func (c *CategoryPost) GetCategoryID() uint {
	return c.categoryId
}

func (c *CategoryPost) SetCategoryID(id uint) {
	c.categoryId = id
}

func (c *CategoryPost) GetPostId() uint {
	return c.postId
}

func (c *CategoryPost) SetPostId(id uint) {
	c.postId = id
}
