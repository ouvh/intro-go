package author

import "fmt"

type Author struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Bio       string `json:"bio"`
}

var Fields = []string{"ID", "FirstName", "LastName", "Bio"}

var ComparableFields = map[string]int{}

func GetField(author Author, field string) (interface{}, error) {
	switch field {
	case "ID":
		return author.ID, nil
	case "FirstName":
		return author.FirstName, nil
	case "LastName":
		return author.LastName, nil
	case "Bio":
		return author.Bio, nil
	default:
		return nil, fmt.Errorf("field '%s' does not exist", field)
	}
}
