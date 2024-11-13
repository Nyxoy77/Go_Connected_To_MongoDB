package router

import (
	controllers "github.com/Nyxoy77/mongoDB/Controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controllers.InsertoneMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controllers.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controllers.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteAll", controllers.DeleteAllMovies).Methods("DELETE")

	return router

}
