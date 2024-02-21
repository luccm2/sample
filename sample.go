// Package sample ...
package sample

// Example represents ...
type Example struct {
	Name string `json:"name"`
}

// NewExample creates a new instance of Example.
func NewExample(name string) *Example {
	return &Example{Name: name}
}
