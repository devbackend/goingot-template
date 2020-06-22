package httpresponse

import (
	"encoding/json"
	"github.com/devbackend/goingot/pkg/logger"
	"net/http"
)

type Sender struct {
	logger logger.Logger
}

func NewSender(logger logger.Logger) *Sender {
	return &Sender{logger: logger}
}

func (s *Sender) JSON(w http.ResponseWriter, status int, data interface{}) {
	resp, err := json.Marshal(data)
	if err != nil {
		s.logger.Error("http response json marshaling error", err)
		status = http.StatusInternalServerError
		resp = []byte(`Internal error`)
	}

	w.WriteHeader(status)
	_, err = w.Write(resp)
	if err != nil {
		s.logger.Error("write http response error", err)
	}
}
