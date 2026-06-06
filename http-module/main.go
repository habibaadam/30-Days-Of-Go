package main

import (
	"fmt"
	"net/http"
)

func welcomeMessage(res http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodGet {
		http.Error(res, "Only Get Methodis allowed", http.StatusMethodNotAllowed)
		return
	}

	_,_ = res.Write([]byte("Hello From GO net/http server"))

}

func main() {

	http.HandleFunc("/hello", welcomeMessage) // registering a route with a handler method -> welcomeMessage

	
    fmt.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil) // specifying a port to listen



	if err == nil {
		fmt.Println(err)
	}

}
