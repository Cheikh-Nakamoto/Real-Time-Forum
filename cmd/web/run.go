package main

import (
	"bim/internal/handlers"
	"bim/internal/metier/database"
	"bim/internal/metier/services/implementations"
	"bim/internal/pkg"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/rs/cors"
)

func Run(tab []string) {
	if len(tab) <= 1 {
		var port string
		db, err := database.InitDatabase()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = pkg.Environment()
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(tab) == 1 {
			p, err := strconv.Atoi(tab[0])
			if err != nil || (p < 1024 || p > 65535) {
				fmt.Println("Port must be a number between 1024 and 65535")
				return
			} else {
				port = ":" + tab[0]
			}
		} else {
			port = os.Getenv("PORT")
		}
		var userService implementations.UserService
		userService.SetDB(db.GetDatabase())
		go userService.CleanUpSessions()

		router := http.NewServeMux() // Création du routeur

		// Passages du gestionnaires des templates
		router.HandleFunc("/api", handlers.Api)
		router.HandleFunc("/api/user", handlers.ApiUser)
		router.HandleFunc("/api/session", handlers.ApiSession)
		router.HandleFunc("/api/login", handlers.ApiLogin)
		router.HandleFunc("/api/register", handlers.ApiRegister)
		router.HandleFunc("/api/messages", handlers.ApiMessage)
		router.HandleFunc("/api/comments", handlers.ApiComment)
		router.HandleFunc("/api/posts", handlers.ApiPost)
		router.HandleFunc(os.Getenv("BASE_URL")+"/online-users", handlers.OnlineUsers)
		router.HandleFunc(os.Getenv("BASE_URL")+"/offline-users", handlers.OfflineUsers)
		router.HandleFunc(os.Getenv("BASE_URL")+"/check-session", handlers.CheckSession)
		router.HandleFunc(os.Getenv("BASE_URL")+"/websocket", handlers.SocketMessage)
		handler := cors.Default().Handler(router)
		// Lancement du serveur d'application
		fmt.Println("Server started on port http://" + os.Getenv("HOST") + port + "/api")
		err = http.ListenAndServe(port, handler)
		if err != nil {
			fmt.Println(err)
			return
		}
		os.Exit(0) // Sortie du programme
	}
	fmt.Println("Usage: go run ./cmd/web/. or go run ./cmd/web/. <port>")
}
