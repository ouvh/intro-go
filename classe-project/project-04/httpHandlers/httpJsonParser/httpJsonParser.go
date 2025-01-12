package httpJsonParser

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SetJson(s interface{}, w http.ResponseWriter, status int) error {
	jsonData, err := json.Marshal(s)

	if err != nil {
		return fmt.Errorf("unable to marshal response")
	}
	w.WriteHeader(status)
	w.Write(jsonData)
	return nil
}

func LoadJson(r *http.Request, DAO interface{}) error {

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		return fmt.Errorf("unable to read request body")
	}
	errr := json.Unmarshal(body, &DAO)
	if errr != nil {
		return fmt.Errorf("unable to parse json")
	}
	return nil

}
