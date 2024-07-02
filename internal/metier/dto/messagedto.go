package dto

type MessageDto struct {
	Id         uint   `json:"id" db:"id"`
	SenderID   uint   `json:"sender_id" db:"sender_id"`
	ReceiverID uint   `json:"receiver_id" db:"receiver"`
	Content    string `json:"content" db:"content"`
	SendedAt   string `json:"sended_at" db:"sended_at"`
}
