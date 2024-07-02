package handlers

import (
	database "bim/internal/metier/database"
	"bim/internal/metier/dto"
	"bim/internal/metier/entities"
	"bim/internal/metier/services/implementations"
	"bim/internal/config"
	"bim/internal/pkg"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// var Sessions = make(map[uint]time.Time)

const (
	contentType    = "Content-Type"
	appType        = "application/json"
	methNotAllowed = "Method not allowed"
	pageNotFound   = "Page Not Found"
	noToken        = "No token provided"
	tokenPayload   = "Invalid token payload"
)

func Register(w http.ResponseWriter, r *http.Request) {
	er := pkg.Environment()
	if er != nil {
		http.Error(w, er.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, methNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get(contentType) != appType {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	if r.URL.Path != os.Getenv("BASE_URL")+"/register" {
		http.Error(w, pageNotFound, http.StatusNotFound)
		return
	}

	config.Cors(w, r)

	var (
		requestDto  dto.RequestDto
		user        entities.User
		userService implementations.UserService
	)

	db, err := database.InitDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userService.SetDB(db.GetDatabase())

	err = json.NewDecoder(r.Body).Decode(&requestDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.SetNickname(requestDto.Nickname)
	user.SetGender(requestDto.Gender)
	user.SetFirstname(requestDto.Firstname)
	user.SetLastname(requestDto.Lastname)
	user.SetEmail(requestDto.Email)
	user.SetPassword(requestDto.Password)
	nbr, err := strconv.Atoi(requestDto.Age)
	if err != nil {
		fmt.Println("erreur de conversion de l'age")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.SetAge(uint(nbr))

	err = userService.AddUser(user.GetNickname(), user.GetGender(), user.GetFirstname(), user.GetLastname(), user.GetEmail(), user.GetPassword(), (user.GetAge()))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	er := pkg.Environment()
	if er != nil {
		http.Error(w, er.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, methNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get(contentType) != appType {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	if r.URL.Path != os.Getenv("BASE_URL")+"/login" {
		http.Error(w, pageNotFound, http.StatusNotFound)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	config.Cors(w, r)

	var (
		requestDto  dto.RequestDto
		userService implementations.UserService
	)

	db, err := database.InitDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userService.SetDB(db.GetDatabase())

	err = json.Unmarshal(body, &requestDto)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	user, err := userService.GetUserByNickname(requestDto.Nickname)
	if err != nil {
		http.Error(w, "Incorrect nickname or password", http.StatusUnauthorized)
		return
	}

	err = pkg.Decrypt(user.GetPassword(), requestDto.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	userDto := dto.RequestDto{
		Id:        user.GetId(),
		Nickname:  user.GetNickname(),
		Gender:    user.GetGender(),
		Firstname: user.GetFirstname(),
		Lastname:  user.GetLastname(),
		Email:     user.GetEmail(),
		Age:       strconv.Itoa(int(user.GetAge())),
	}

	token := config.BuildJWT(config.CreateHeader(), config.EncodePayload(userDto), os.Getenv("SECRET_KEY"))
	//Sessions[user.GetId()] = time.Now()
	err = userService.UpdateUserActivity(user.GetId(), true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(contentType, appType)
	err = json.NewEncoder(w).Encode(map[string]string{"token": token})
	if err != nil {
		return
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	er := pkg.Environment()
	if er != nil {
		http.Error(w, er.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, methNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != os.Getenv("BASE_URL")+"/logout" {
		http.Error(w, pageNotFound, http.StatusNotFound)
		return
	}

	config.Cors(w, r)

	token := r.Header.Get("Authorization")
	parts := strings.Split(token, " ")
	if len(parts) != 2 && parts[0] != "Bearer" || token == "" {
		http.Error(w, noToken, http.StatusBadRequest)
		return
	}

	tokenParts := strings.Split(parts[1], ".")
	if len(tokenParts) != 3 {
		http.Error(w, "Invalid token format", http.StatusUnauthorized)
		return
	}

	if !config.CheckToken(tokenParts[0], tokenParts[1], tokenParts[2], os.Getenv("SECRET_KEY")) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	payload, err := base64.RawURLEncoding.DecodeString(tokenParts[1])
	if err != nil {
		http.Error(w, tokenPayload, http.StatusUnauthorized)
		return
	}

	var userService implementations.UserService
	db, err := database.InitDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userService.SetDB(db.GetDatabase())

	var payloadData map[string]interface{}
	err = json.Unmarshal(payload, &payloadData)
	if err != nil {
		http.Error(w, tokenPayload, http.StatusUnauthorized)
		return
	}
	fmt.Println(payloadData)

	userID, ok := payloadData["id"].(float64)
	if !ok {
		http.Error(w, tokenPayload, http.StatusUnauthorized)
		return
	}

	err = userService.UpdateUserActivity(uint(userID), false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Logged out successfully"})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func OnlineUsers(w http.ResponseWriter, r *http.Request) {
	er := pkg.Environment()
	if er != nil {
		http.Error(w, er.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, methNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != os.Getenv("BASE_URL")+"/online-users" {
		http.Error(w, pageNotFound, http.StatusNotFound)
		return
	}

	config.Cors(w, r)

	var (
		users       []dto.RequestDto
		userService implementations.UserService
	)

	db, err := database.InitDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userService.SetDB(db.GetDatabase())

	onlineUsers, err := userService.GetOnlineUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, user := range onlineUsers {
		userDto := dto.RequestDto{
			Id:        user.GetId(),
			Nickname:  user.GetNickname(),
			Gender:    user.GetGender(),
			Firstname: user.GetFirstname(),
			Lastname:  user.GetLastname(),
			Email:     user.GetEmail(),
			Age:       strconv.Itoa(int(user.GetAge())),
			Image:     user.GetImage(),
		}
		users = append(users, userDto)
	}

	w.Header().Set(contentType, appType)
	if len(users) == 0 {
		err = json.NewEncoder(w).Encode(map[string]string{"message": "No online users found!"})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else {
		err = json.NewEncoder(w).Encode(users)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func OfflineUsers(w http.ResponseWriter, r *http.Request) {
	er := pkg.Environment()
	if er != nil {
		http.Error(w, er.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, methNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != os.Getenv("BASE_URL")+"/offline-users" {
		http.Error(w, pageNotFound, http.StatusNotFound)
		return
	}

	config.Cors(w, r)

	var (
		usersDto    []dto.RequestDto
		userService implementations.UserService
	)

	db, err := database.InitDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userService.SetDB(db.GetDatabase())

	offlineUsers, err := userService.GetOfflineUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, user := range offlineUsers {
		userDto := dto.RequestDto{
			Id:        user.GetId(),
			Nickname:  user.GetNickname(),
			Gender:    user.GetGender(),
			Firstname: user.GetFirstname(),
			Lastname:  user.GetLastname(),
			Email:     user.GetEmail(),
			Age:       strconv.Itoa(int(user.GetAge())),
			Image:     user.GetImage(),
		}
		usersDto = append(usersDto, userDto)
	}

	w.Header().Set(contentType, appType)
	if len(usersDto) == 0 {
		err = json.NewEncoder(w).Encode(map[string]string{"message": "No offline users found!"})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else {
		err = json.NewEncoder(w).Encode(usersDto)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func CheckSession(w http.ResponseWriter, r *http.Request) {
	er := pkg.Environment()
	if er != nil {
		http.Error(w, er.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, methNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != os.Getenv("BASE_URL")+"/check-session" {
		http.Error(w, pageNotFound, http.StatusNotFound)
		return
	}

	config.Cors(w, r)

	token := r.Header.Get("Authorization")
	if strings.TrimSpace(token) == "" {
		http.Error(w, noToken, http.StatusUnauthorized)
		return
	}

	parts := strings.Split(token, " ")
	if len(parts) != 2 && parts[0] != "Bearer" || token == "" {
		http.Error(w, noToken, http.StatusBadRequest)
		return
	}

	header, payload, signature := config.SplitToken(token)
	if !config.CheckToken(header, payload, signature, os.Getenv("SECRET_KEY")) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
