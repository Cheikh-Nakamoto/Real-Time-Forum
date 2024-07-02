package dto

type RequestDto struct {
	Id        uint   `json:"id"`
	Nickname  string `json:"nickname"`
	Age       string   `json:"age"`
	Gender    string `json:"gender"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Image     string `json:"image"`
}
