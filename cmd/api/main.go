package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tkdnbb/bookofben-api/internal/routes"
)

func main() {
	r := routes.SetupRoutes()

	fmt.Println("Bible API Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
