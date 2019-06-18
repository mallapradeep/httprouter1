package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

//Laying up foundation of get/post/delete endpoints

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	//we want to get the URl Parameters that were passed into that endpoint
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	//encodes ppl endpoint to json n return it to the user
	json.NewEncoder(w).Encode(people)

}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	//we r gonna take the json data from the post body and MARSHAL it to the structure
	_ = json.NewDecoder(req.Body).Decode(&person)
	//we r passing in a custom made ID
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	//setting up mux router
	router := mux.NewRouter()

	//creating sample set data
	people = append(people, Person{ID: "1", Firstname: "Pradeep", Lastname: "Malla"})
	people = append(people, Person{ID: "2", Firstname: "Priyamka", Lastname: "KArki"})

	//setting up endpoints
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")

	//setting up our http server
	log.Fatal(http.ListenAndServe(":8080", router))
}
