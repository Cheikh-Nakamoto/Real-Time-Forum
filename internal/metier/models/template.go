package models

type TemplateData struct {
	StringData map[string]string
	IntData    map[string]int
	PageName   string
	Users      []User
	Posts      []Post
	Categories []Category
	Category   Category
	Actions    Action
	Action     Action
	User       User
	Post       Post
	Error      ErrorData
	Short      ShortMessage
	IsOnline   bool
}

type ErrorData struct {
	Code    string
	Message string
}

type ShortMessage struct {
	Type    string
	Message string
}
