package form

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

var (
	ErrNotSupported     = errors.New("type not supported for conversion")
	ErrWrongType        = errors.New("value is wrong type")
	ErrBadStruct        = errors.New("struct missing json tag")
	ErrNotStruct        = errors.New("model is not a struct")
	ErrRequiredMissing  = errors.New("required field missing")
	ErrWrongContentType = errors.New("content-type of request is incorrect")
)

// Prevent running on types other than struct
func structPtrCheck(i interface{}) bool {
	if reflect.TypeOf(i).Kind() != reflect.Ptr {
		return false
	} else if reflect.TypeOf(i).Elem().Kind() != reflect.Struct {
		return false
	}

	return true
}

// Validate returns true if the submitted form has the required fields
func Validate(r *http.Request, model interface{}) (error, string) {
	// Prevent running on types other than struct
	if !structPtrCheck(model) {
		return ErrNotStruct, ErrNotStruct.Error()
	}

	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		return ErrWrongContentType, ErrWrongContentType.Error()
	}

	// Parse Form
	r.ParseForm()

	// Get the struct type
	t := reflect.TypeOf(model).Elem()

	// Look through each struct field
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Get name tag
		name := field.Tag.Get("json")
		if len(name) == 0 {
			return ErrBadStruct, fmt.Sprintf("%v errored because the json tag is missing", field.Name)
		}

		// Check required tag
		if strings.ToLower(field.Tag.Get("require")) == "true" {
			sentVal := r.FormValue(name)
			if len(sentVal) == 0 {
				return ErrRequiredMissing, fmt.Sprintf("%v is missing", name)
			}
		}
	}

	return nil, ""
}

// TypeCopy copies values from request form to struct
func StructCopy(r *http.Request, model interface{}) (error, string) {
	// Prevent running on types other than struct
	if !structPtrCheck(model) {
		return ErrNotStruct, ErrNotStruct.Error()
	}

	// Parse Form
	r.ParseForm()

	// Get the struct type
	modelType := reflect.TypeOf(model).Elem()
	modelValue := reflect.ValueOf(model).Elem()

	// Loop through the submitted fields
	for keyName, value := range r.Form {

		// Find the field where the json value is equal
		for i := 0; i < modelType.NumField(); i++ {
			field := modelType.Field(i)

			// Get name tag
			name := field.Tag.Get("json")

			// If the form value matches a struct
			if name == keyName {
				fieldValue := modelValue.FieldByName(field.Name)
				// If the field exists and there is a passed value (length if zero is possible)
				if fieldValue.IsValid() && len(value) > 0 {

					// Get the first value even though there could be multiple values
					singleValue := value[0]

					// Set the value in the model to the correct value and type
					err, _ := typeConvert(singleValue, fieldValue)
					if err == ErrWrongType {
						return err, fmt.Sprintf("%v needs to be type: %v", name, fieldValue.Type())
					} else if err != nil {
						return err, fmt.Sprintf("%v errored because the type (%v) is not supported", name, fieldValue.Type())
					}
				}

				break
			}
		}

		// Anything that gets to here is not a value field so just drop it
		// You can return an error here if extra values are not permitted
	}

	return nil, ""
}

// convert safely converts the string to the value type and assigns the value
// Returns a standard error and error specific text
func typeConvert(s string, v reflect.Value) (error, string) {
	var err error

	// Convert to correct type
	switch v.Kind() {
	case reflect.String:
		v.SetString(s)

	case reflect.Bool:
		var parsed bool
		parsed, err = strconv.ParseBool(s)
		v.SetBool(parsed)

	case reflect.Float32:
		var parsed float64
		parsed, err = strconv.ParseFloat(s, 32)
		v.SetFloat(parsed)
	case reflect.Float64:
		var parsed float64
		parsed, err = strconv.ParseFloat(s, 64)
		v.SetFloat(parsed)

	case reflect.Int:
		var parsed int64
		parsed, err = strconv.ParseInt(s, 10, 0)
		v.SetInt(parsed)
	case reflect.Int8:
		var parsed int64
		parsed, err = strconv.ParseInt(s, 0, 8)
		v.SetInt(parsed)
	case reflect.Int16:
		var parsed int64
		parsed, err = strconv.ParseInt(s, 0, 16)
		v.SetInt(parsed)
	case reflect.Int32:
		var parsed int64
		parsed, err = strconv.ParseInt(s, 0, 32)
		v.SetInt(parsed)
	case reflect.Int64:
		var parsed int64
		parsed, err = strconv.ParseInt(s, 0, 64)
		v.SetInt(parsed)

	case reflect.Uint:
		var parsed uint64
		parsed, err = strconv.ParseUint(s, 10, 0)
		v.SetUint(parsed)
	case reflect.Uint8:
		var parsed uint64
		parsed, err = strconv.ParseUint(s, 0, 8)
		v.SetUint(parsed)
	case reflect.Uint16:
		var parsed uint64
		parsed, err = strconv.ParseUint(s, 0, 16)
		v.SetUint(parsed)
	case reflect.Uint32:
		var parsed uint64
		parsed, err = strconv.ParseUint(s, 0, 32)
		v.SetUint(parsed)
	case reflect.Uint64:
		var parsed uint64
		parsed, err = strconv.ParseUint(s, 0, 64)
		v.SetUint(parsed)

	default:
		return ErrNotSupported, fmt.Sprintf("Type conversion is not supported for type: %v", v.Type())
	}

	if err != nil {
		return ErrWrongType, err.Error()
	}
	return nil, ""
}

func StructTags(model interface{}, tag string) ([]string, error) {
	var arr []string

	// Prevent running on types other than struct
	if !structPtrCheck(model) {
		return arr, ErrNotStruct
	}

	// Get the struct type
	t := reflect.TypeOf(model).Elem()

	// Look through each struct field
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Get tag
		name := field.Tag.Get(tag)
		if len(name) != 0 {
			arr = append(arr, name)
		}
	}

	return arr, nil
}

func StructFields(model interface{}) ([]string, error) {
	var arr []string

	// Prevent running on types other than struct
	if !structPtrCheck(model) {
		return arr, ErrNotStruct
	}

	// Get the struct type
	t := reflect.TypeOf(model).Elem()

	// Look through each struct field
	for i := 0; i < t.NumField(); i++ {
		arr = append(arr, t.Field(i).Name)
	}

	return arr, nil
}
