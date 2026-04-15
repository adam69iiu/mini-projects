package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)


func Health(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("this is amazing"))
}
func main()  {
	r := chi.NewRouter()
	r.Get("/health", Health)
	fmt.Println("giga niga")
	fmt.Println("http://localhost:6767")
 if err := http.ListenAndServe(":6767",r); err !=nil{
	 fmt.Println("port is reseved")
 }
}
