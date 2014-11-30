package swagger

import (
	"encoding/json"

	"github.com/casualjim/go-swagger/reflection"
)

// Info object provides metadata about the API.
// The metadata can be used by the clients if needed, and can be presented in the Swagger-UI for convenience.
//
// For more information: http://goo.gl/8us55a#infoObject
type Info struct {
	Extensions     map[string]interface{} `swagger:"-"` // custom extensions, omitted when empty
	Description    string                 `swagger:"description,omitempty"`
	Title          string                 `swagger:"title,omitempty"`
	TermsOfService string                 `swagger:"termsOfService,omitempty"`
	Contact        *ContactInfo           `swagger:"contact,omitempty"`
	License        *License               `swagger:"license,omitempty"`
	Version        string                 `swagger:"version,omitempty"`
}

// UnmarshalMap hydrates this info instance with the data from the map
func (i *Info) UnmarshalMap(data interface{}) error {
	dict := reflection.MarshalMap(data)
	if err := reflection.UnmarshalMapRecursed(dict, i); err != nil {
		return err
	}
	i.Extensions = readExtensions(dict)
	return nil
}

// UnmarshalJSON hydrates this info instance with the data from JSON
func (i *Info) UnmarshalJSON(data []byte) error {
	var value map[string]interface{}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	return i.UnmarshalMap(value)
}

// UnmarshalYAML hydrates this info instance with the data from YAML
func (i *Info) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value map[string]interface{}
	if err := unmarshal(&value); err != nil {
		return err
	}
	return i.UnmarshalMap(value)
}

// MarshalMap converts this info object to a map
func (i Info) MarshalMap() map[string]interface{} {
	res := reflection.MarshalMapRecursed(i)
	addExtensions(res, i.Extensions)
	return res
}

// MarshalJSON converts this info object to JSON
func (i Info) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.MarshalMap())
}

// MarshalYAML converts this info object to YAML
func (i Info) MarshalYAML() (interface{}, error) {
	return i.MarshalMap(), nil
}
