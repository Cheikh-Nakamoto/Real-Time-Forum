package implementations

import (
	"bim/internal/metier/entities"
	"bim/internal/pkg"
	"database/sql"
	"errors"
)

type UserService struct {
	db *sql.DB
}

func (u *UserService) GetDB() *sql.DB {
	return u.db
}

func (u *UserService) SetDB(db *sql.DB) {
	u.db = db
}

func (u *UserService) AddUser(nickname, gender, firstname, lastname, email, password string, age uint) error {
	var user entities.User // Instance of user

	if !pkg.CheckNickname(nickname) || !pkg.CheckEmail(email) || !pkg.CheckName(firstname) || !pkg.CheckName(lastname) || !pkg.CheckPassword(password) || !pkg.CheckAge(age) || !pkg.CheckGender(gender) {
		return errors.New("nickname, email, firstname, lastname, password and age must be valid")
	}

	pass, err := pkg.Hash(password)
	if err != nil {
		return err
	}

	user.SetNickname(nickname)
	user.SetGender(gender)
	user.SetFirstname(firstname)
	user.SetLastname(lastname)
	user.SetEmail(email)
	user.SetPassword(pass)
	user.SetAge(age)

	check1, err1 := u.UserExists(nickname)
	check2, err2 := u.UserExists(email)
	if err1 != nil || err2 != nil {
		return err
	}
	if check1 || check2 {
		return errors.New("user already exists")
	}

	_, err = u.GetDB().Exec(`INSERT INTO users (nickname, age, gender, firstname, lastname, email, password) VALUES (?, ?, ?, ?, ?, ?, ?)`, user.GetNickname(), user.GetAge(), user.GetGender(), user.GetFirstname(), user.GetLastname(), user.GetEmail(), user.GetPassword())
	return err
}

func (u *UserService) GetUserById(id uint) (entities.User, error) {
	var (
		user      entities.User
		nickname  string
		age       uint
		gender    string
		firstname string
		lastname  string
		email     string
		image     string
	)

	err := u.GetDB().QueryRow(`SELECT nickname, age, gender, firstname, lastname, email, image FROM users WHERE id = ?`, id).Scan(&nickname, &age, &gender, &firstname, &lastname, &email, &image)
	if err != nil {
		return user, err
	}

	if !pkg.CheckNickname(nickname) || !pkg.CheckEmail(email) || !pkg.CheckName(firstname) || !pkg.CheckName(lastname) /* || !pkg.CheckPassword(password) */ || !pkg.CheckAge(age) || !pkg.CheckGender(gender) {
		return user, errors.New("data is invalid")
	}

	user.SetId(id)
	user.SetNickname(nickname)
	user.SetAge(age)
	user.SetGender(gender)
	user.SetFirstname(firstname)
	user.SetLastname(lastname)
	user.SetEmail(email)
	user.SetImage(image)

	return user, nil
}

func (u *UserService) GetUserByNickname(username string) (entities.User, error) {
	var (
		user      entities.User
		id        uint
		nickname  string
		age       uint
		gender    string
		firstname string
		lastname  string
		email     string
		password  string
		image     string
	)

	err := u.GetDB().QueryRow(`SELECT id, nickname, age, gender, firstname, lastname, email, password, image FROM users WHERE nickname = ? OR email = ?`, username, username).Scan(&id, &nickname, &age, &gender, &firstname, &lastname, &email, &password, &image)
	if err != nil {
		return user, err // User does not exist
	}

	if !pkg.CheckNickname(nickname) || !pkg.CheckEmail(email) {
		return user, errors.New("data is invalid")
	}

	user.SetId(id)
	user.SetNickname(nickname)
	user.SetAge(age)
	user.SetGender(gender)
	user.SetFirstname(firstname)
	user.SetLastname(lastname)
	user.SetEmail(email)
	user.SetPassword(password)
	user.SetImage(image)
	return user, nil
}

func (u *UserService) GetAllUsers() ([]entities.User, error) {
	rows, err := u.GetDB().Query(`SELECT id, nickname, age, gender, firstname, lastname, email, image FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var (
			user      entities.User
			id        uint
			nickname  string
			gender    string
			firstname string
			lastname  string
			email     string
			image     string
			age       uint
		)
		err = rows.Scan(&id, &nickname, &age, &gender, &firstname, &lastname, &email, &image)
		if err != nil {
			return nil, err
		}

		if !pkg.CheckNickname(nickname) || !pkg.CheckEmail(email) || !pkg.CheckName(firstname) || !pkg.CheckName(lastname) /* || !pkg.CheckPassword(password) */ || !pkg.CheckAge(age) || !pkg.CheckGender(gender) {
			return users, errors.New("data is invalid")
		}

		user.SetId(id)
		user.SetNickname(nickname)
		user.SetAge(age)
		user.SetGender(gender)
		user.SetFirstname(firstname)
		user.SetLastname(lastname)
		user.SetEmail(email)
		user.SetImage(image)
		users = append(users, user)
	}
	return users, nil
}

func (u *UserService) UserExists(nickname string) (bool, error) {
	var id uint
	err := u.GetDB().QueryRow(`SELECT id FROM users WHERE nickname = ? OR email = ?`, nickname, nickname).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil // User does not exist
		}
		return false, errors.New("unexpected error") // Unexpected error
	}
	return true, nil // User already exists
}
