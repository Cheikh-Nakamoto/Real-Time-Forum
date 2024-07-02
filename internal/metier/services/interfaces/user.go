package interfaces

import "bim/internal/metier/entities"

type IUser interface {
	AddUser(nickname, gender, firstname, lastname, email, password, image string, age uint) error
	GetUserById(id uint) (entities.User, error)
	GetUserByNickname(nickname string) (entities.User, error)
	UserExists(nickname string) (bool, error)
	SessionClear(id uint)
	// GetActiveUsers(active bool) ([]entities.User, error)
}
