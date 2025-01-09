package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Adress struct {
	Street      string `json:"street"`
	City        string `json:"city"`
	State       string `json:"state"`
	Postal_code string `json:"postal_code"`
}

type Course struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Credit int    `json:"credit"`
}

type Student struct {
	Id      int64    `json:"id"`
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Adress  Adress   `json:"address"`
	Major   string   `json:"major"`
	Courses []Course `json:"courses"`
}

type Error struct {
	Error string `json:"error"`
}

var index int = 1
var Students []Student = make([]Student, 0)

func studentshandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")

		jsonData, err := json.Marshal(Students)
		if err != nil {
			errrr, _ := json.Marshal(Error{"Unmarshalling went wrong"})
			w.WriteHeader(http.StatusInternalServerError) // 200 OK
			w.Write(errrr)
			return
		}

		w.WriteHeader(http.StatusOK) // 200 OK
		w.Write(jsonData)

	} else if r.Method == http.MethodPost {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			errrr, _ := json.Marshal(Error{"Unable to read request body"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errrr)
			return
		}
		defer r.Body.Close()

		requestData := Student{}

		err = json.Unmarshal(body, &requestData)
		if err != nil {
			errrr, _ := json.Marshal(Error{"Unable to parse JSON"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errrr)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		log.Println(requestData)

		//verify all fields manually
		if requestData.Name != "" && requestData.Age > 0 && requestData.Major != "" && requestData.Adress.Street != "" && requestData.Adress.City != "" && requestData.Adress.State != "" && requestData.Adress.Postal_code != "" {
			for _, course := range requestData.Courses {
				if course.Code == "" || course.Name == "" || course.Credit == 0 {
					errrr, _ := json.Marshal(Error{"Missing Required Values"})
					w.WriteHeader(http.StatusBadRequest)
					w.Write(errrr)
					return
				}
			}
			requestData.Id = int64(index)
			Students = append(Students, requestData)
			index++

		} else {
			errrr, _ := json.Marshal(Error{"Missing Required Values"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errrr)
			return
		}

		jsonData, err := json.Marshal(requestData)
		if err != nil {
			errrr, _ := json.Marshal(Error{"Unable to marshal response"})
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(errrr)
			return
		}

		w.WriteHeader(http.StatusCreated) // 201 OK
		w.Write(jsonData)
	} else {
		errrr, _ := json.Marshal(Error{"Invalid request method"})
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(errrr)
		return
	}

}

func studenthandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		requestPath := strings.TrimPrefix(r.URL.Path, "/students/")
		id, err := strconv.ParseInt(requestPath, 10, 64)

		if err != nil {
			errrr, _ := json.Marshal(Error{"Unable to read request path variable"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errrr)
			return
		}

		for _, student := range Students {
			if student.Id == id {
				w.Header().Set("Content-Type", "application/json")
				jsonData, err := json.Marshal(student)
				if err != nil {
					errrr, _ := json.Marshal(Error{"Unable to marshal response"})
					w.WriteHeader(http.StatusInternalServerError)
					w.Write(errrr)
					return
				}
				w.WriteHeader(http.StatusOK) // 200 OK
				w.Write(jsonData)
				return
			}
		}

		errrr, _ := json.Marshal(Error{"Student not found"})
		w.WriteHeader(http.StatusNotFound)
		w.Write(errrr)
		return

	} else if r.Method == http.MethodPut {

		requestPath := strings.TrimPrefix(r.URL.Path, "/students/")
		id, err := strconv.ParseInt(requestPath, 10, 64)

		if err != nil {
			errrr, _ := json.Marshal(Error{"Unable to read request path variable"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errrr)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			errrr, _ := json.Marshal(Error{"Unable to read request body"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errrr)
			return
		}
		defer r.Body.Close()

		requestData := Student{}

		err = json.Unmarshal(body, &requestData)
		if err != nil {
			errrr, _ := json.Marshal(Error{"Unable to parse JSON"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errrr)
			return
		}

		//verify all fields manually
		if requestData.Name != "" && requestData.Age > 0 && requestData.Major != "" && requestData.Adress.Street != "" && requestData.Adress.City != "" && requestData.Adress.State != "" && requestData.Adress.Postal_code != "" {
			for _, course := range requestData.Courses {
				if course.Code == "" || course.Name == "" || course.Credit == 0 {
					errrr, _ := json.Marshal(Error{"Missing Required Values"})
					w.WriteHeader(http.StatusBadRequest)
					w.Write(errrr)
					return
				}
			}
		} else {
			errrr, _ := json.Marshal(Error{"Missing Required Values"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errrr)
			return
		}

		for ii, student := range Students {
			if student.Id == id {
				requestData.Id = id
				Students[ii] = requestData
				w.Header().Set("Content-Type", "application/json")
				jsonData, err := json.Marshal(requestData)
				if err != nil {
					http.Error(w, "Unable to marshal response", http.StatusInternalServerError)
					return
				}

				w.WriteHeader(http.StatusOK) // 200 OK
				w.Write(jsonData)
				return

			}
		}

		errrr, _ := json.Marshal(Error{"Student Not Found"})
		w.WriteHeader(http.StatusNotFound)
		w.Write(errrr)
		return

	} else if r.Method == http.MethodDelete {
		requestPath := strings.TrimPrefix(r.URL.Path, "/students/")
		id, err := strconv.ParseInt(requestPath, 10, 64)

		if err != nil {
			errrr, _ := json.Marshal(Error{"Unable to read request path variable"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errrr)
			return
		}

		for indexx, student := range Students {
			if student.Id == id {
				w.Header().Set("Content-Type", "application/json")
				Students = append(Students[:indexx], Students[indexx+1:]...)
				w.WriteHeader(http.StatusOK) // 200 OK
				return
			}
		}
		errrr, _ := json.Marshal(Error{" Student not found"})
		w.WriteHeader(http.StatusNotFound)
		w.Write(errrr)
		return
	} else {

		errrr, _ := json.Marshal(Error{"Invalid request method"})
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(errrr)
		return
	}

}

func main() {
	http.HandleFunc("/students", studentshandler)
	http.HandleFunc("/students/", studenthandler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("error serving : ", err)
	}

}
