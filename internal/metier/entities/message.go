package entities

import "time"

type Message struct {
	id         uint
	senderID   uint
	receiverID uint
	content    string
	sendedAt   time.Time
}

func (m *Message) GetID() uint {
	return m.id
}

func (m *Message) SetID(id uint) {
	m.id = id
}

func (m *Message) GetSenderID() uint {
	return m.senderID
}

func (m *Message) SetSenderID(senderID uint) {
	m.senderID = senderID
}

func (m *Message) GetReceiverID() uint {
	return m.receiverID
}

func (m *Message) SetReceiverID(receiverID uint) {
	m.receiverID = receiverID
}

func (m *Message) GetContent() string {
	return m.content
}

func (m *Message) SetContent(content string) {
	m.content = content
}

func (m *Message) GetSendedAt() time.Time {
	return m.sendedAt
}

func (m *Message) SetSendedAt() {
	m.sendedAt = time.Now()
}
