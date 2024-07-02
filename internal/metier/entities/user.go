package entities

// User entity
type User struct {
	id        uint   `json:"id"`
	nickname  string `json:"nickname"`
	age       uint   `json:"age"`
	gender    string `json:"gender"`
	firstname string `json:"firstname"`
	lastname  string `json:"lastname"`
	email     string `json:"email"`
	password  string `json:"password"`
	image     string `json:"image"`
}

// Getters and Setters
func (u *User) GetId() uint {
	return u.id
}

func (u *User) SetId(id uint) {
	u.id = id
}

func (u *User) GetNickname() string {
	return u.nickname
}

func (u *User) SetNickname(nickname string) {
	u.nickname = nickname
}

func (u *User) GetAge() uint {
	return u.age
}

func (u *User) SetAge(age uint) {
	u.age = age
}

func (u *User) GetGender() string {
	return u.gender
}

func (u *User) SetGender(gender string) {
	u.gender = gender
}

func (u *User) GetFirstname() string {
	return u.firstname
}

func (u *User) SetFirstname(firstname string) {
	u.firstname = firstname
}

func (u *User) GetLastname() string {
	return u.lastname
}

func (u *User) SetLastname(lastname string) {
	u.lastname = lastname
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) GetImage() string {
	return u.image
}

func (u *User) SetImage(image string) {
	u.image = image
}
