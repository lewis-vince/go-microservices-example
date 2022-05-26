package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func returnAllPets(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Pets)
}

func returnSinglePet(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["id"]

     // Loop over all of our Articles
    // if the article.Id equals the key we pass in
    // return the article encoded as JSON
    for _, article := range Pets {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}

func createNewPet(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // unmarshal this into a new Article struct
    // append this to our Articles array.    
    reqBody, _ := ioutil.ReadAll(r.Body)
    var article Pet 
    json.Unmarshal(reqBody, &article)
    // update our global Articles array to include
    // our new Article
    Pets = append(Pets, article)

    json.NewEncoder(w).Encode(article)
}

func deletePet(w http.ResponseWriter, r *http.Request) {
    // once again, we will need to parse the path parameters
    vars := mux.Vars(r)
    // we will need to extract the `id` of the article we
    // wish to delete
    id := vars["id"]

    // we then need to loop through all our articles
    for index, article := range Pets {
        // if our id path parameter matches one of our
        // articles
        if article.Id == id {
            // updates our Articles array to remove the 
            // article
            Pets = append(Pets[:index], Pets[index+1:]...)
        }
    }

}
func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the Pet Shop! pages include pets,pet, GET/DELETE pet/id")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/pets", returnAllPets).Methods(("GET"))
    // NOTE: Ordering is important here! This has to be defined before
    // the other `/article` endpoint. 
    myRouter.HandleFunc("/pet", createNewPet).Methods("POST")
    myRouter.HandleFunc("/pet/{id}", deletePet).Methods("DELETE")
    myRouter.HandleFunc("/pet/{id}", returnSinglePet).Methods("GET")
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Pets = []Pet{
        Pet{Id: "1", Name: "Lilly", Species: "Dog", DateOfBirth: "2011-01-01"},
        Pet{Id: "2", Name: "Petal", Species: "Pangolin", DateOfBirth: "2010-12-25"},
        Pet{Id: "3", Name: "Poppy", Species: "Capybara", DateOfBirth: "2010-12-24"},

    }
    handleRequests()
}

type Pet struct {
	Id      string `json:"Id"`
    Name   string `json:"Name"`
    Species    string `json:"Species"`
    DateOfBirth string `json:"DOB"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Pets []Pet