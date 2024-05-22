package main

import (
	"cmp"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Brawl345/gitwebhook/api"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	http.HandleFunc("/api/webhook", api.Handler)
	port := cmp.Or(os.Getenv("PORT"), "8080")
	log.Printf("Listening on %s", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)

	if err != nil {
		log.Fatal(err)
	}
}
