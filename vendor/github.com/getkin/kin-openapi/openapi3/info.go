package openapi3

import (
	"context"
	"errors"
	"fmt"

	"github.com/getkin/kin-openapi/jsoninfo"
)

// Info is specified by OpenAPI/Swagger standard version 3.0.
type Info struct {
	ExtensionProps
	Title          string   `json:"title" yaml:"title"` // Required
	Description    string   `json:"description,omitempty" yaml:"description,omitempty"`
	TermsOfService string   `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        *License `json:"license,omitempty" yaml:"license,omitempty"`
	Version        string   `json:"version" yaml:"version"` // Required
}

func (value *Info) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(value)
}

func (value *Info) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, value)
}

func (value *Info) Validate(c context.Context) error {
	if contact := value.Contact; contact != nil {
		if err := contact.Validate(c); err != nil {
			return fmt.Errorf("Error when validating Contact: %s", err.Error())
		}
	}

	if license := value.License; license != nil {
		if err := license.Validate(c); err != nil {
			return fmt.Errorf("Error when validating License: %s", err.Error())
		}
	}

	if value.Version == "" {
		return errors.New("Variable 'version' must be a non-empty JSON string")
	}

	if value.Title == "" {
		return errors.New("Variable 'title' must be a non-empty JSON string")
	}

	return nil
}

// Contact is specified by OpenAPI/Swagger standard version 3.0.
type Contact struct {
	ExtensionProps
	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	URL   string `json:"url,omitempty" yaml:"url,omitempty"`
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
}

func (value *Contact) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(value)
}

func (value *Contact) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, value)
}

func (value *Contact) Validate(c context.Context) error {
	return nil
}

// License is specified by OpenAPI/Swagger standard version 3.0.
type License struct {
	ExtensionProps
	Name string `json:"name" yaml:"name"` // Required
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
}

func (value *License) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(value)
}

func (value *License) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, value)
}

func (value *License) Validate(c context.Context) error {
	if value.Name == "" {
		return errors.New("Variable 'name' must be a non-empty JSON string")
	}
	return nil
}