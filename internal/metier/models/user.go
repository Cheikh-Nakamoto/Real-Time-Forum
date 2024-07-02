package models

import (
	"github.com/google/uuid"
)

type User struct {
	id       uuid.UUID
	username string
	name     string
	email    string
	password string
	image    string //default image : "https://cdn.pixabay.com/photo/2020/07/01/12/58/icon-5359553_960_720.png"
}

// GetId returns the id of the user
func (u *User) GetId() uuid.UUID {
	return u.id
}

// SetId sets the id of the user
func (u *User) SetId(id uuid.UUID) {
	u.id = id
}

// GetUsername returns the username
func (u *User) GetUsername() string {
	return u.username
}

// SetUsername sets the username
func (u *User) SetUsername(username string) {
	u.username = username
}

// GetName returns the name of the user
func (u *User) GetName() string {
	return u.name
}

// SetName sets the name of the user
func (u *User) SetName(name string) {
	u.name = name
}

// GetEmail returns the email
func (u *User) GetEmail() string {
	return u.email
}

// SetEmail sets the email
func (u *User) SetEmail(email string) {
	u.email = email
}

// GetPassword returns the password
func (u *User) GetPassword() string {
	return u.password
}

// SetPassword sets the password
func (u *User) SetPassword(password string) {
	u.password = password
}

// GetImage returns the image associated with the user
func (u *User) GetImage() string {
	return u.image
}

// SetImage sets the image associated with the user
func (u *User) SetImage(image string) {
	u.image = image
}
