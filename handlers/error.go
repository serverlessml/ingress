/*
Copyright © 2020 Dmitry Kisler <admin@dkisler.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package handlers

import (
	"fmt"
)

// Error defines the error.
type Error struct {
	// Type defines the error type.
	Type string
	// Message defines human readable error message.
	Message string
	// Details defines details of the error.
	Details interface{}
}

// Error returns a human readable error message.
func (e Error) Error() string {
	prefix := fmt.Sprintf("[%s]", e.Type)
	if e.Details == nil {
		return fmt.Sprintf("%s %s", prefix, e.Message)
	}
	return fmt.Sprintf("%s %s. Details:\n%v", prefix, e.Message, e.Details)
}

// ErrorPush defines the errors of submitting jobs.
type ErrorPush struct {
	// contains error message
	Message string `json:"message"`
	// contains pipeline config
	Details interface{} `json:"details"`
}

// Error returns a human readable error message.
func (e ErrorPush) Error() string {
	if e.Details == nil {
		return e.Message
	}
	return fmt.Sprintf("%s. Details:\n%v", e.Message, e.Details)
}

// ErrorArray returns human readable error messages.
func ErrorArray(errors []error) []string {
	e := []string{}
	for _, err := range errors {
		e = append(e, err.Error())
	}
	return e
}
