package main

import (
	"api/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)

	// Create a new router
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Go Sample API is running...")

	fmt.Println(`
 ______     ______        ______     ______   __    
/\  ___\   /\  __ \      /\  __ \   /\  == \ /\ \   
\ \ \__ \  \ \ \/\ \     \ \  __ \  \ \  _-/ \ \ \  
 \ \_____\  \ \_____\     \ \_\ \_\  \ \_\    \ \_\ 
  \/_____/   \/_____/      \/_/\/_/   \/_/     \/_/ `)

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
