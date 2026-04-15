package server

import (
	"fmt"
	"net/http"

	"github.com/adam69iiu/backend/internal/config"
	"github.com/adam69iiu/backend/internal/database"
	"github.com/adam69iiu/backend/internal/db"
	"github.com/adam69iiu/backend/internal/repository"
	"github.com/adam69iiu/backend/internal/service"
	"github.com/go-chi/chi/v5"
)


func Health(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("this is amazing"))
}
func Run() error{
	cfg := config.Load()
	conn,err := database.ConnectDatabase(cfg.DtatabaseURL)
	if err != nil {
		return err
	}

	queries := db.New(conn) 
  userRepo := repository.NewUserRepo(queries) 
	authService := service.NewAuthService(userRepo, cfg.JWTSecret)
	r := chi.NewRouter()
	r.Get("/health", Health)
	fmt.Println("http://localhost:6767")
 if err := http.ListenAndServe(":6767",r); err !=nil{
	 return err
 }
 return nil
}
