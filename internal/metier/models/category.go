package models

import (
	"bim/internal/pkg"
	"html"
)

type Category struct {
	id   uint
	name string
}

// GetID returns the id of the category
func (c *Category) GetID() uint {
	return c.id
}

// SetID sets the id of the category
func (c *Category) SetID(id uint) {
	c.id = id
}

// GetName returns the name of the category
func (c *Category) GetName() string {
	return html.EscapeString(pkg.Capitalize(c.name))
}

// SetName sets the name of the category
func (c *Category) SetName(name string) {
	c.name = pkg.Capitalize(name)
}
