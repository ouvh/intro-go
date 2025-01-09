package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

/*
resp,err := http.Get("http://wefwefwe")

if err != nil {
	log.Fatal("something went Wrong")
}

defer resp.Body.Close()
body,err := ioutil.ReadAll(resp.Body)





data := url.Values{}
data.Set("name","John")
data.Set("age","30")

resp,err := http.PostForm("rgrre/form")




var result APiresponse

json.NewDecoder(resp.Body).Decode(&result)


// for the error

http.Error(w,err.Error(),http.StatusInternalServerError)

queryParams := r.URL.Query()



*/

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func handlerjson(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Hello, World!",
		Code:    200,
	}
	w.Header().Set("Content-Type", "application/json")

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.WriteHeader(http.StatusOK) // 200 OK
	w.Write(jsonData)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Check if the method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Read the body of the request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Declare a variable of the Request type
	requestData := make(map[string]any, 10)

	// Unmarshal the body into the Request struct
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(w, "Unable to parse JSON", http.StatusBadRequest)
		return
	}

	// Do something with the parsed data
	fmt.Printf("Received: %+v\n", requestData)

	// Send a JSON response
	response := map[string]interface{}{
		"message": "Data received successfully",
		"data":    requestData,
	}

	// Set the response header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Marshal and send the response as JSON
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Unable to marshal response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK) // 200 OK
	w.Write(jsonData)
}

func handlerHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, `{"text":"Hello , world !"}`)
}

func handleWithQueryParam(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/home", handlerHome)

	err := http.ListenAndServe(":8085", nil)

	if err != nil {
		log.Fatal("error serving : ", err)
	}

}
