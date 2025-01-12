package storeJsonParser

import (
	"encoding/json"
	"os"
)

func SaveJson(s interface{}, file string) error {
	f, _ := os.Create(file)
	defer f.Close()
	as_json, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		return err
	}
	f.Write(as_json)
	return nil
}

func LoadJson(DAO interface{}, file string) error {

	f, err := os.ReadFile(file)

	if err != nil {
		return err
	}

	err = json.Unmarshal(f, DAO)
	if err != nil {
		return err
	}

	return nil

}
