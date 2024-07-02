package interfaces

import "bim/internal/metier/entities"

type IMessage interface {
	SaveMessage(message entities.Message) error
}
