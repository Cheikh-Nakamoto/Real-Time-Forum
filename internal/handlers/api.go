package handlers

import (
	"bim/internal/config"
	database "bim/internal/metier/database"
	"bim/internal/metier/dto"
	"bim/internal/metier/entities"
	"bim/internal/metier/services/implementations"
	service "bim/internal/metier/services/implementations"
	"bim/internal/pkg"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"
)

var postS = []dto.PostDAO{}
var Users = []dto.RequestDto{}

func initPost() ([]dto.PostDAO, error) {
	var Posts []dto.PostDAO
	// Save the post to the database (you need to implement this)
	DB, err := database.InitDatabase()
	if err != nil {
		return nil, err
	}
	defer DB.Close()

	postserveice := service.PostService{}
	postserveice.SetDB(DB.GetDatabase())
	AllPost, err := postserveice.GetAllPosts()

	for _, v := range AllPost {
		Posts = append(Posts, dto.PostDAO{
			ID:        v.GetID(),
			Title:     v.GetTitle(),
			Content:   v.GetContent(),
			Image:     v.GetImage(),
			CreatedAt: v.GetCreatedAt(),
			UpdatedAt: v.GetUpdatedAt(),
			Owner:     v.GetOwner(),
			// Comments:   v.GetComments(),
			// Like:       v.GetLike(),
			// Dislike:    v.GetDislike(),
		})
	}

	if err != nil {
		return nil, err
	}
	return Posts, nil
}

func Api(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"messages": "http://localhost:1112/api/messages",
		"user":     "http://localhost:1112/api/user",
		"posts":    "http://localhost:1112/api/posts",
		"comments": "http://localhost:1112/api/comments",
	}
	if r.URL.Path == "/api" {
		// Écrire la réponse JSON
		body, err := json.Marshal(data)
		if err != nil {
			// Gérer l'erreur de marshaling JSON
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}

func ApiUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.URL.Path == "/api/user" && r.Method == http.MethodPost {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var ID map[string]string
		err = json.Unmarshal(data, &ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		db, err := database.InitDatabase()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()
		userService := implementations.UserService{}
		userService.SetDB(db.GetDatabase())
		nbr, err := strconv.Atoi(ID["ID"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		userinf, err := userService.GetUserById(uint(nbr))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var user dto.RequestDto
		user.Id = userinf.GetId()
		user.Nickname = userinf.GetNickname()
		user.Age = strconv.Itoa(int(userinf.GetAge()))
		user.Gender = userinf.GetGender()
		user.Firstname = userinf.GetFirstname()
		user.Lastname = userinf.GetLastname()
		user.Email = userinf.GetEmail()
		user.Image = userinf.GetImage()
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func ApiSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.URL.Path == "/api/session" && r.Method == http.MethodPost {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("ici1")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var ID map[string]string
		err = json.Unmarshal(data, &ID)
		if err != nil {
			fmt.Println("ici2")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		db, err := database.InitDatabase()
		if err != nil {
			fmt.Println("ici 3")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()
		userService := implementations.UserService{}
		userService.SetDB(db.GetDatabase())
		nbr, err := strconv.Atoi(ID["ID"])
		if err != nil {
			fmt.Println("ici 4")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		userService.SessionClear(uint(nbr))
	}
}

func ApiRegister(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api/register" && r.Method == http.MethodPost {
		er := pkg.Environment()
		if er != nil {
			http.Error(w, er.Error(), http.StatusInternalServerError)
			return
		}

		// if r.Header.Get(contentType) != appType {
		// 	http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		// 	return
		// }

		if r.URL.Path != os.Getenv("BASE_URL")+"/register" {
			http.Error(w, pageNotFound, http.StatusNotFound)
			return
		}

		//config.Cors(w, r)

		var (
			requestDto  dto.RequestDto
			user        entities.User
			userService implementations.UserService
		)

		db, err := database.InitDatabase()
		if err != nil {
			fmt.Println("Erreur d'initialisation de la bdd")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		userService.SetDB(db.GetDatabase())
		// r.ParseForm()
		// requestDto.Nickname = r.FormValue("nickname")
		// requestDto.Gender = r.FormValue("gender")
		// requestDto.Firstname = r.FormValue("firstname")
		// requestDto.Lastname = r.FormValue("lastname")
		// requestDto.Email = r.FormValue("email")
		// requestDto.Password = r.FormValue("password")
		// uin,_ := strconv.Atoi(r.FormValue("Age"))
		//requestDto.Age = uint(uin)
		//fmt.Println(requestDto.Nickname, requestDto.Gender, requestDto.Firstname, requestDto.Lastname, requestDto.Email)
		err = json.NewDecoder(r.Body).Decode(&requestDto)
		if err != nil {
			fmt.Println("l'erreur est  : ", err)
			fmt.Println("Erreur dans le decodage de ", r.Body)
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
			fmt.Println("L'erreur est au niveau de l'age : ", err)
		}
		user.SetAge(uint(nbr))
		fmt.Println(user)
		err = userService.AddUser(user.GetNickname(), user.GetGender(), user.GetFirstname(), user.GetLastname(), user.GetEmail(), user.GetPassword(), user.GetAge())
		if err != nil {
			fmt.Println("L'erreyr est :", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("Un utilisateur a été ajouté")
		json.NewEncoder(w).Encode(map[string]string{"messages": "Compte créé avec succes !"})
		w.WriteHeader(http.StatusCreated)

	} else if r.URL.Path == "/api/register" && r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(map[string]string{"messages": "http://localhost:1112/api/messages"})
	}
}

func ApiLogin(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api/login" && r.Method == http.MethodPost {

		er := pkg.Environment()
		if er != nil {
			http.Error(w, er.Error(), http.StatusInternalServerError)
			return
		}

		// if r.Header.Get(contentType) != appType {
		// 	http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		// 	return
		// }

		// if r.URL.Path != os.Getenv("BASE_URL")+"/login" {
		// 	http.Error(w, pageNotFound, http.StatusNotFound)
		// 	return
		// }

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Erreur lors de la lecture du body")
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
			fmt.Println("Erreur lors du decodage de la unmarshal")
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			return
		}
		fmt.Println("nickname :", requestDto.Nickname)
		user, err := userService.GetUserByNickname(requestDto.Nickname)
		if err != nil {
			fmt.Println("Erreur lors de la recuperation de l'utilisateur")
			http.Error(w, "Incorrect nickname or password", http.StatusUnauthorized)
			return
		}

		err = pkg.Decrypt(user.GetPassword(), requestDto.Password)
		if err != nil {
			fmt.Println("Erreur lors de la decryption du mot de passe")
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
			Age:       string(user.GetAge()),
		}

		token := config.BuildJWT(config.CreateHeader(), config.EncodePayload(userDto), os.Getenv("SECRET_KEY"))

		w.Header().Set(contentType, appType)

		w.Header().Set(contentType, appType)
		err = json.NewEncoder(w).Encode(map[string]string{"token": token, "object": strconv.Itoa(int(userDto.Id))})
		if err != nil {
			fmt.Println("echec de l'encodage ")
			return
		}
		err = userService.UpdateUserActivity(user.GetId(), true)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func ApiMessage(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"message": "Hello, world!",
	}
	if r.URL.Path == "/api/messages" {
		// Écrire la réponse JSON
		body, err := json.Marshal(data)
		if err != nil {
			// Gérer l'erreur de marshaling JSON
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}

func ApiComment(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"message": "Hello, world!",
	}
	if r.URL.Path == "/api/comments" {
		// Écrire la réponse JSON
		body, err := json.Marshal(data)
		if err != nil {
			// Gérer l'erreur de marshaling JSON
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}

func ApiPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api/posts" && r.Method == http.MethodPost {

		fmt.Println("Requête POST reçue sur /api/posts")

		// Parse the multipart form
		err := r.ParseMultipartForm(10 << 20) // 10 MB
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Get the form values
		title := r.FormValue("title")
		content := r.FormValue("content")
		category := r.FormValue("categorie")
		file, handler, err := r.FormFile("image")
		if err != nil {
			fmt.Println("Error uploading file")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Create the file path
		filePath := fmt.Sprintf("./assets/img/%s", handler.Filename)
		fmt.Println(filePath)
		dst, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Error creating file")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		// Copy the uploaded file to the destination
		if _, err := io.Copy(dst, file); err != nil {
			fmt.Println("Error copying file")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Create the post object
		post := dto.PostDAO{
			Title:    title,
			Content:  content,
			Image:    filePath,
			Category: category,
		}
		postS = append(postS, post)
		// Save the post to the database (you need to implement this)
		DB, err := database.NewDatabase()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer DB.Close()

		postserveice := service.PostService{}
		postserveice.SetDB(DB.GetDatabase())
		err = postserveice.CreatePost(title, content, filePath, uuid.Nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Return a response with the image path
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message":   "Form submitted successfully!",
			"imagePath": filePath,
		})
	} else if r.URL.Path == "/api/posts" && r.Method == http.MethodGet {
		postes, err := initPost()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(postes)

	} else {
		fmt.Println("Requête non supportée ou mauvaise méthode")
		http.Error(w, "Invalid request method or path", http.StatusMethodNotAllowed)
	}
}
