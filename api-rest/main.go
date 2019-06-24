package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main(){
	//Creacion de ruta con libreria
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/" , Index)
	router.HandleFunc("/contacto", Contact)
	router.HandleFunc("/peliculas", MovieList)
	router.HandleFunc("/pelicula/{id}", MovieShow)
	//creacion de la ruta sin libreria
	/*http.HandleFunc("/" , func(w http.ResponseWriter , r *http.Request){
		fmt.Fprintf(w , "Hola mundo desde mi servidor en GO")
	})*/
	fmt.Println("Se inicia una ejecucion")
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)

}//termina func

func Index (w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w , "Hola mundo desde mi servidor en GO")
}

func Contact (w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w , "Esta es la pagina de contacto")
}

func MovieList (w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w , "Listado de peliculas")
}

func MovieShow (w http.ResponseWriter , r *http.Request){
	params := mux.Vars(r)
	movie_id := params["id"]

	fmt.Fprintf(w , "Has cargado la pelicula numero: %s",movie_id)
}