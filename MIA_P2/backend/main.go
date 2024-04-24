package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type responseList struct {
	Status int64    `json:"Status"`
	List   []string `json:"List"`
}

type responseString struct {
	Status int64  `json:"Status"`
	Value  string `json:"Value"`
}

type loginValues struct {
	User     string `json:"User"`
	Password string `json:"Password"`
}

func postMethod(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Data")
	}
	var newLoginValues loginValues
	json.Unmarshal(reqBody, &newLoginValues)

	fmt.Println(newLoginValues.User)
	fmt.Println(newLoginValues.Password)

	var newResponseList responseList
	newResponseList.Status = 200
	newResponseList.List = []string{"A", "B", "C", "D", "E", "F"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newResponseList)
}

func getMethod(w http.ResponseWriter, r *http.Request) {
	var newResponseString responseString
	newResponseString.Status = 200
	newResponseString.Value = "Hello World"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newResponseString)
}

func handleRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Weltome to my  API :D")
}

func main() {
	fmt.Println("Server started on port 4000")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handleRoute)
	router.HandleFunc("/tasks", postMethod).Methods("POST")
	router.HandleFunc("/tasks", getMethod).Methods("GET")

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":4000", handler))

}
