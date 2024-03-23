package sender

import (
	"encoding/json"
	"net/http"

	"github.com/devbackend/goingot/pkg/log"
)

type ErrorResponse struct {
	Reason string   `json:"reason"`
	Errors []string `json:"errors"`
}

type Sender struct {
	logger log.Logger
}

// New return sender
func New(logger log.Logger) *Sender {
	return &Sender{logger}
}

// SendJSON send response to HTTP client in content-type application/json
func (s *Sender) SendJSON(w http.ResponseWriter, status int, data interface{}) {
	resp, err := marshall(data)
	if err != nil {
		s.logger.Error().Msgf("http response json marshaling error - %v", err)

		status = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	if _, err := w.Write(resp); err != nil {
		s.logger.Error().Msgf("write http response error - %v", err)
	}
}

// SendOK send response with HTTP-status OK
func (s *Sender) SendOK(w http.ResponseWriter, data interface{}) {
	s.SendJSON(w, http.StatusOK, data)
}

// SendNoContent send response with HTTP-status NoContent
func (s *Sender) SendNoContent(w http.ResponseWriter) {
	s.SendJSON(w, http.StatusNoContent, nil)
}

// SendBadRequest send response with HTTP-status BadRequest
func (s *Sender) SendBadRequest(w http.ResponseWriter, reason string, errors []string) {
	if errors == nil {
		errors = []string{}
	}

	s.SendJSON(w, http.StatusBadRequest, ErrorResponse{
		Reason: reason,
		Errors: errors,
	})
}

// SendBadRequest send response with HTTP-status BadRequest
func (s *Sender) SendForbidden(w http.ResponseWriter) {
	s.SendJSON(w, http.StatusForbidden, ErrorResponse{
		Reason: "Forbidden",
	})
}

// SendNotFound send response with HTTP-status NotFound
func (s *Sender) SendNotFound(w http.ResponseWriter, reason string) {
	s.SendJSON(w, http.StatusNotFound, ErrorResponse{
		Reason: reason,
	})
}

// SendInternalError send response with HTTP-status InternalServerError
func (s *Sender) SendInternalError(w http.ResponseWriter, err error) {
	if err != nil {
		s.logger.Error().Msg(err.Error())
	}

	s.SendJSON(w, http.StatusInternalServerError, nil)
}

func marshall(data interface{}) (resp []byte, err error) {
	if data == nil {
		return
	}

	resp, err = json.Marshal(data)
	if err != nil {
		resp = []byte(`{"status": 500, "body": "internal error"}`)
	}

	return
}
