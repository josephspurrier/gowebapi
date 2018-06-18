package bind

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"

	"github.com/matryer/way"
	"gopkg.in/go-playground/validator.v9"
)

// Binder contains the request bind an validator objects.
type Binder struct {
	validator *validator.Validate
}

// New returns a new binder for request bind and validation.
func New() *Binder {
	return &Binder{
		validator: validator.New(),
	}
}

// Validate will validate a struct using the validator.
func (b *Binder) Validate(s interface{}) error {
	return b.validator.Struct(s)
}

// JSONUnmarshal will perform an unmarshal on an interface using JSON.
/*func (b *Binder) JSONUnmarshal(iface interface{}, keys []string,
	values []string) (err error) {
	// Check for errors.
	v := reflect.ValueOf(iface)
	if v.Kind() != reflect.Ptr {
		return errors.New("must pass a pointer, not a value")
	} else if len(keys) != len(values) {
		return errors.New("keys and values array or not the same length")
	}

	// Load the map.
	m := make(map[string]interface{})
	for i := 0; i < len(keys); i++ {
		m[keys[i]] = values[i]
	}

	// Convert to JSON.
	var data []byte
	data, err = json.Marshal(m)
	if err != nil {
		return
	}

	// Unmarshal to the interface from JSON.
	return json.Unmarshal(data, &iface)
}*/

// FormUnmarshal will perform an unmarshal on an interface using a form.
func (b *Binder) FormUnmarshal(iface interface{}, r *http.Request) (err error) {
	// Check for errors.
	v := reflect.ValueOf(iface)
	if v.Kind() != reflect.Ptr {
		return errors.New("must pass a pointer, not a value")
	}

	// Parse the form.
	err = r.ParseForm()
	if err != nil {
		return err
	}

	// Load the map.
	m := make(map[string]interface{})
	for k, v := range r.Form {
		m[k] = v[0]
	}

	// Loop through each field to extract the URL parameter.
	elem := reflect.Indirect(v.Elem())
	keys := elem.Type()
	for j := 0; j < elem.NumField(); j++ {
		tag := keys.Field(j).Tag
		tagvalue := tag.Get("json")
		pathParam := way.Param(r.Context(), tagvalue)
		if len(pathParam) > 0 {
			m[tagvalue] = pathParam
		}
	}

	// Convert to JSON.
	var data []byte
	data, err = json.Marshal(m)
	if err != nil {
		return
	}

	// Unmarshal to the interface from JSON.
	return json.Unmarshal(data, &iface)
}
