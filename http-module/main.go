package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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


// request type
type TestRequest struct{
	Name string `json:"name"`
}

// helper function to take in data and encode to json
func WriteJSON(res http.ResponseWriter, status int, data any) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	_ = json.NewEncoder(res).Encode(data)
}

func testHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		WriteJSON(res, http.StatusMethodNotAllowed, map[string]any {
			"ok": "false",
			"error": "Only post is allowed",
		})
		return
	}
	defer req.Body.Close()

	var request TestRequest

	decoded := json.NewDecoder(req.Body)
	if error := decoded.Decode(&request); error != nil {
		WriteJSON(res, http.StatusBadRequest, map[string]any{
			"ok": false,
			"error": "Invalid json format",
		})
		return
	}
	request.Name = strings.TrimSpace(request.Name)

	if request.Name == " " {
		WriteJSON(res, http.StatusBadRequest, map[string]any{
			"ok": false,
			"error": "Fields must not ne empty",
		})
		return
	}

	WriteJSON(res, http.StatusOK, map[string]any{
		"ok": true,
		"data": request,
		"time": time.Now().UTC(),
	})

}


func main() {
    http.HandleFunc("/", rootRouteHandler)
	http.HandleFunc("/hello", welcomeMessage) // registering a route with a handler method -> welcomeMessage
	http.HandleFunc("/propergreet", welcomeQueryParam)
	http.HandleFunc("/testdecoding", testHandler)



    fmt.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil) // specifying a port to listen



	if err == nil {
		fmt.Println(err)
	}

}
