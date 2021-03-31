package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type addressBook struct {
	Firstname string
	Lastname  string
	Code      int
	Phone     string
}

func getAddressBookAll(w http.ResponseWriter, r *http.Request) {
	addBook := addressBook{
		Firstname: "Arunee",
		Lastname:  "Poonsawat",
		Code:      2000,
		Phone:     "+61x-xxxx-xxxx",
	}
	json.NewEncoder(w).Encode(addBook)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage!")
}
func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("No Port In Heroku " + port)
	}
	return ":" + port
}
func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/getAddress", getAddressBookAll)
	//http.ListenAndServe(":8080", nil)
	http.ListenAndServe(getPort(), nil)
}
func main() {
	handleRequest()
}

// nameless-mountain-27438

// go get github.com/tools/godep
// godep save

// create Procfile // web go-api
// https://medium.com/odds-team/%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2-go-%E0%B8%95%E0%B8%AD%E0%B8%99-9-%E0%B8%A7%E0%B8%B4%E0%B8%98%E0%B8%B5-deploy-go-lang-%E0%B9%84%E0%B8%9B%E0%B8%A2%E0%B8%B1%E0%B8%87-heroku-cloud-%E0%B9%81%E0%B8%9A%E0%B8%9A-step-by-step-df9e52599291
