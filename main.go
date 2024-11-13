package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/Nyxoy77/mongoDB/Controllers/Router"
)

func main() {
	fmt.Println("MONGO DB API ")
	r := router.Router()

	log.Fatal(http.ListenAndServe(":8000", r))

}
