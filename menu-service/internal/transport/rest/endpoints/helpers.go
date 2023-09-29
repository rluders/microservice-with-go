package endpoints

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Payload    interface{} `json:"payload,omitempty"`
}

func (r *Response) Marshal() []byte {
	jsonResponse, err := json.Marshal(r)
	if err != nil {
		log.Printf("failed to marshal response: %v", err)
	}

	return jsonResponse
}

func sendValidationError(w http.ResponseWriter, err error) {
	validationErrors := err.(validator.ValidationErrors)

	fieldErrors := make(map[string][]string)
	for _, vErr := range validationErrors {
		fieldName := vErr.Field()
		fieldError := fieldName + " " + vErr.Tag()

		fieldErrors[fieldName] = append(fieldErrors[fieldName], fieldError)
	}

	response := &Response{
		StatusCode: http.StatusBadRequest,
		Message:    "Vaidation error",
		Payload: map[string]interface{}{
			"errors": fieldErrors,
		},
	}

	writeResponse(w, response)
}

func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	response := &Response{
		StatusCode: statusCode,
		Message:    message,
	}

	writeResponse(w, response)
}

func sendDataResponse[T interface{}](w http.ResponseWriter, message string, statusCode int, payload *T) {
	response := &Response{
		StatusCode: statusCode,
		Message:    message,
		Payload:    payload,
	}

	writeResponse(w, response)
}

func sendResponse(w http.ResponseWriter, message string, statusCode int) {
	response := &Response{
		StatusCode: statusCode,
		Message:    message,
	}

	writeResponse(w, response)
}

func writeResponse(w http.ResponseWriter, r *Response) {
	jsonResponse := r.Marshal()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)

	_, err := w.Write(jsonResponse)
	if err != nil {
		log.Printf("Error writing response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func parseRequest[T interface{}](r *T, body io.ReadCloser) error {
	if err := json.NewDecoder(body).Decode(&r); err != nil {
		return err
	}
	return nil
}

func isRequestValid(request interface{}) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(request)
}
