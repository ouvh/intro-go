package httpStructVerifier

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func Contains(s string, container []string) bool {
	for _, value := range container {
		if value == s {
			return true
		}
	}
	return false

}

func ValidateSearchCriteria(fields []string, criterias map[string]any, ignoreFields []string) error {

	for key := range criterias {
		found := Contains(key, fields)
		if found {
			continue
		} else if !found && Contains(key, ignoreFields) {
			continue
		} else {
			return (fmt.Errorf("unknown parameter %s", key))
		}

	}
	return nil

}

func ValidateStruct(s interface{}) error {
	return validateField(reflect.ValueOf(s), "")
}

func validateField(val reflect.Value, parentField string) error {
	// Handle pointer input
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			fieldName := parentField
			if fieldName == "" {
				fieldName = "input"
			}
			return fmt.Errorf("this field %s is not set", fieldName)
		}
		val = val.Elem()
	}

	// Handle different types
	switch val.Kind() {
	case reflect.Struct:
		// Get type for field names
		typ := val.Type()

		// Check each field
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldName := typ.Field(i).Name

			// Skip ID field
			if fieldName == "ID" {
				continue
			}

			// Skip unexported fields
			if !typ.Field(i).IsExported() {
				continue
			}

			// Build field path for nested structs
			fullFieldName := fieldName
			if parentField != "" {
				fullFieldName = fmt.Sprintf("%s.%s", parentField, fieldName)
			}

			// Validate the field
			if err := validateField(field, fullFieldName); err != nil {
				return err
			}
		}

	case reflect.Slice, reflect.Array:
		// Check if slice/array is nil or empty
		if !val.IsValid() || val.IsNil() || val.Len() == 0 {
			fieldName := parentField
			if fieldName == "" {
				fieldName = "slice"
			}
			return fmt.Errorf("this field %s is not set", fieldName)
		}

		// Validate each element in the slice/array
		for i := 0; i < val.Len(); i++ {
			elem := val.Index(i)
			fieldName := fmt.Sprintf("%s[%d]", parentField, i)
			if err := validateField(elem, fieldName); err != nil {
				return err
			}
		}

	default:
		// For basic types, check if they're set
		if !val.IsValid() || val.IsZero() {
			fieldName := parentField
			if fieldName == "" {
				fieldName = "field"
			}
			return fmt.Errorf("this field %s is not set", fieldName)
		}
	}

	return nil
}

func ValidatePathId(r *http.Request, prefix string) (int64, error) {
	requestPath := strings.TrimPrefix(r.URL.Path, prefix)
	id, err := strconv.ParseInt(requestPath, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("bad path parameter")
	}
	return id, err
}
