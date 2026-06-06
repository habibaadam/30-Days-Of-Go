package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func welcomeMessage(res http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodGet {
		http.Error(res, "Only Get Method is allowed", http.StatusMethodNotAllowed)
		return
	}

	_,_ = res.Write([]byte("Hello From GO net/http server"))

}

func rootRouteHandler(res http.ResponseWriter, req *http.Request) {
	_, _ = res.Write([]byte("Welcome to my simple server with GO"))
}

func welcomeQueryParam(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json") //setting a header

	res.WriteHeader(http.StatusOK) // setting a status code
	name := req.URL.Query().Get("name") //specifies getting value of a query parameter

	if name == " " {
		name = "Guest User"
	}

	message := map[string]any{
		"name": name,
		"message": "JSON encode successful",
		"date": time.Now().UTC(),
	}

	_ = json.NewEncoder(res).Encode(message)
}



func main() {
    http.HandleFunc("/", rootRouteHandler)
	http.HandleFunc("/hello", welcomeMessage) // registering a route with a handler method -> welcomeMessage
	http.HandleFunc("/propergreet", welcomeQueryParam)



    fmt.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil) // specifying a port to listen



	if err == nil {
		fmt.Println(err)
	}

}
