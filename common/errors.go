package common

import "fmt"

type ErrorCollection struct {
	errors []error
}

func (ec *ErrorCollection) Add(err error) {
	if err == nil {
		return
	}
	ec.errors = append(ec.errors, err)
}

func (ec *ErrorCollection) FirstError() error {
	if !ec.HasErrors() {
		return nil
	}
	return ec.errors[0]
}

func (ec *ErrorCollection) Errors() []error {
	return ec.errors
}

func (ec *ErrorCollection) HasErrors() bool {
	return len(ec.errors) > 0
}

func (ec *ErrorCollection) String() string {
	return fmt.Sprintf("ErrorCollection{errors: %v}", len(ec.errors))
}

func NewErrorCollection() *ErrorCollection {
	return &ErrorCollection{}
}
