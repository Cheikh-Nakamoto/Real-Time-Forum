package implementations

import (
	"bim/internal/metier/dto"
	"bim/internal/metier/entities"
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

const (
	invalidData = "invalid data"
)

func (u *UserService) UpdateUserActivity(userID uint, online bool) error {
	var err error
	if online {
		timestamp := time.Now()
		_, err = u.GetDB().Exec(`INSERT OR REPLACE INTO sessions (user_id, last_active) VALUES (?,?)`, userID, timestamp)
	} else {
		_, err = u.GetDB().Exec(`DELETE FROM sessions WHERE user_id = ?`, userID)
	}
	return err
}

func (u *UserService) CleanUpSessions() {
	for {
		_, err := u.GetDB().Exec(`DELETE  FROM sessions WHERE last_active <?`, time.Now().Add(-5*time.Minute))
		if err != nil {
			fmt.Println("Cannot clean up sessions: ", err)
			return
		}
		time.Sleep(1 * time.Minute)
	}
}

func (u *UserService) SessionClear(id uint) {
	_, err := u.GetDB().Exec(`DELETE  FROM sessions WHERE user_id = ?`, id)
    if err!= nil {
        fmt.Println("Cannot clean up sessions: ", err)
        return
    }
}

func (u *UserService) GetOnlineUsers() ([]entities.User, error) {
	rows, err := u.GetDB().Query(`SELECT id, nickname, age, gender, firstname, lastname, email, image FROM users WHERE id IN (SELECT user_id FROM sessions WHERE last_active > ?)`, time.Now().Add(-5*time.Minute))
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Cannot close rows: ", err)
			return
		}
	}(rows)

	var users []entities.User
	for rows.Next() {
		var (
			userTemp dto.RequestDto
			user     entities.User
		)
		if err := rows.Scan(&userTemp.Id, &userTemp.Nickname, &userTemp.Age, &userTemp.Gender, &userTemp.Firstname, &userTemp.Lastname, &userTemp.Email, &userTemp.Image); err != nil {
			return nil, err
		}
		user.SetId(userTemp.Id)
		user.SetNickname(userTemp.Nickname)
		nbr, _ := strconv.Atoi(userTemp.Age)
		user.SetAge(uint(nbr))
		user.SetGender(userTemp.Gender)
		user.SetFirstname(userTemp.Firstname)
		user.SetLastname(userTemp.Lastname)
		user.SetEmail(userTemp.Email)
		user.SetImage(userTemp.Image)
		users = append(users, user)
	}
	return users, nil
}

func (u *UserService) GetOfflineUsers() ([]entities.User, error) {
	rows, err := u.GetDB().Query(`SELECT id, nickname, age, gender, firstname, lastname, email, image FROM users WHERE id NOT IN (SELECT user_id FROM sessions WHERE last_active > ?)`, time.Now().Add(-5*time.Minute))
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Cannot close rows: ", err)
			return
		}
	}(rows)

	var users []entities.User
	for rows.Next() {
		var (
			userTemp dto.RequestDto
			user     entities.User
		)
		if err := rows.Scan(&userTemp.Id, &userTemp.Nickname, &userTemp.Age, &userTemp.Gender, &userTemp.Firstname, &userTemp.Lastname, &userTemp.Email, &userTemp.Image); err != nil {
			return nil, err
		}
		user.SetId(userTemp.Id)
		user.SetNickname(userTemp.Nickname)
		nbr, _ := strconv.Atoi(userTemp.Age)
		user.SetAge(uint(nbr))
		user.SetGender(userTemp.Gender)
		user.SetFirstname(userTemp.Firstname)
		user.SetLastname(userTemp.Lastname)
		user.SetEmail(userTemp.Email)
		user.SetImage(userTemp.Image)
		users = append(users, user)
	}
	return users, nil
}
