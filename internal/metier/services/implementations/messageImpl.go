package implementations

import (
	"bim/internal/metier/entities"
	"database/sql"
)

type MessageService struct {
	db *sql.DB
}

func (m *MessageService) GetDB() *sql.DB {
	return m.db
}

func (m *MessageService) SetDB(db *sql.DB) {
	m.db = db
}

func (m *MessageService) SaveMessage(message entities.Message) error {
	_, err := m.GetDB().Exec(`INSERT INTO messages (sender_id, recevier_id, content) VALUES (?, ?, ?)`, message.GetSenderID(), message.GetReceiverID(), message.GetContent())
	return err
}
