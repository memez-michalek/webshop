package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type product struct{
	name string
	quantity int
	price float64
	category string
	id string
}
type key struct{
	id string
	username string
	value string

}
func generateAPIkeyHandler(w http.ResponseWriter, r * http.Request){
	



}
func main() {
	router := mux.NewRouter()
	router.



}
