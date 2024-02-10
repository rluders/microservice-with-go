package endpoints

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Body    any    `json:"body,omitempty"`
}

func (r *Response) Marshal() []byte {
	jsonResponse, err := json.Marshal(r)
	if err != nil {
		log.Printf("failed to marshal response: %v", err)
	}

	return jsonResponse
}

type ValidationErrors struct {
	Errors map[string][]string `json:"errors,omitempty"`
}

func NewValidationErrors(err error) *ValidationErrors {
	var validationErrors validator.ValidationErrors
	errors.As(err, &validationErrors)

	fieldErrors := make(map[string][]string)
	for _, vErr := range validationErrors {
		fieldName := vErr.Field()
		fieldError := fieldName + " " + vErr.Tag()

		fieldErrors[fieldName] = append(fieldErrors[fieldName], fieldError)
	}

	return &ValidationErrors{Errors: fieldErrors}
}

func isRequestValid(request any) *ValidationErrors {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(request)
	if err != nil {
		return NewValidationErrors(err)
	}

	return nil
}

func sendResponse[T any](w http.ResponseWriter, message string, code int, body *T) {
	response := &Response{
		Code:    code,
		Message: message,
		Body:    body,
	}

	writeResponse(w, response)
}

func writeResponse(w http.ResponseWriter, r *Response) {
	jsonResponse := r.Marshal()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Code)

	_, err := w.Write(jsonResponse)
	if err != nil {
		log.Printf("Error writing response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func parseRequest[T any](r *T, body io.ReadCloser) error {
	if err := json.NewDecoder(body).Decode(&r); err != nil {
		return err
	}
	return nil
}
