package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/lachikhin-mikhail/medods_test/api"
	"github.com/lachikhin-mikhail/medods_test/internal/db"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}

	conn, err := db.ConnectDB()
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close(context.Background())

	// Адрес для запуска сервера
	ip := ""
	port := os.Getenv("PORT")
	addr := fmt.Sprintf("%s:%s", ip, port)

	// Router
	r := chi.NewRouter()

	r.Post("/api/signin", api.PostSigninHandler)
	r.Post("/api/refresh", api.PostRefreshHandler)

	// Запуск сервера
	err = http.ListenAndServe(addr, r)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Server running on %s\n", port)

}
