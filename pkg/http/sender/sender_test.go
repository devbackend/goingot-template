package sender_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/devbackend/goingot/pkg/http/sender"
	"github.com/devbackend/goingot/pkg/log"
)

func TestSender_SendJSON(t *testing.T) {
	cases := []struct {
		status           int
		data             interface{}
		expectedStatus   int
		expectedResponse string
	}{
		{
			status:           http.StatusOK,
			data:             "ok",
			expectedStatus:   http.StatusOK,
			expectedResponse: `"ok"`,
		},
		{
			status:           http.StatusOK,
			data:             func() {},
			expectedStatus:   http.StatusInternalServerError,
			expectedResponse: `{"status": 500, "body": "internal error"}`,
		},
		{
			status:           http.StatusCreated,
			data:             nil,
			expectedStatus:   http.StatusCreated,
			expectedResponse: "",
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, c := range cases {
		w := httptest.NewRecorder()

		buf := bytes.Buffer{}
		logger := log.New(&buf)
		s := sender.New(logger)

		s.SendJSON(w, c.status, c.data)

		assert.Equal(t, c.expectedStatus, w.Result().StatusCode)
		assert.Equal(t, c.expectedResponse, w.Body.String())
	}
}

func TestSender_SendOK(t *testing.T) {
	cases := []struct {
		data             interface{}
		expectedResponse string
	}{
		{
			expectedResponse: "",
		},
		{
			expectedResponse: `{"msg":"hello!"}`,
			data: struct {
				Msg string `json:"msg"`
			}{"hello!"},
		},
	}

	for _, c := range cases {
		w := httptest.NewRecorder()

		buf := bytes.Buffer{}
		logger := log.New(&buf)
		s := sender.New(logger)

		s.SendOK(w, c.data)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Equal(t, c.expectedResponse, w.Body.String())
	}
}

func TestSender_SendBadRequest(t *testing.T) {
	cases := []struct {
		reason           string
		errors           []string
		expectedResponse string
	}{
		{
			reason:           "test reason 1",
			expectedResponse: `{"reason":"test reason 1","errors":[]}`,
		},
		{
			reason:           "test reason 2",
			errors:           []string{"error 1"},
			expectedResponse: `{"reason":"test reason 2","errors":["error 1"]}`,
		},
	}

	for _, c := range cases {
		w := httptest.NewRecorder()

		buf := bytes.Buffer{}
		logger := log.New(&buf)
		s := sender.New(logger)

		s.SendBadRequest(w, c.reason, c.errors)

		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
		assert.Equal(t, c.expectedResponse, w.Body.String())
	}
}

func TestSender_SendInternalError(t *testing.T) {
	cases := []struct {
		name        string
		err         error
		expectedLog string
	}{
		{
			name:        "test error",
			err:         errors.New("test error"),
			expectedLog: "test error",
		},
		{
			name:        "NIL error",
			expectedLog: "",
		},
	}

	for _, c := range cases {
		c := c // scopelint mute

		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			buf := bytes.Buffer{}
			logger := log.New(&buf)
			s := sender.New(logger)

			s.SendInternalError(w, c.err)

			assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
			assert.Contains(t, buf.String(), c.expectedLog)
		})
	}
}

func TestSender_SendNotFound(t *testing.T) {
	tests := []struct {
		name             string
		reason           string
		expectedResponse string
	}{
		{
			name:             "empty reason",
			expectedResponse: `{"reason":"","errors":null}`,
		},
		{
			name:             "Test not found",
			reason:           "Test not found",
			expectedResponse: `{"reason":"Test not found","errors":null}`,
		},
	}
	for _, c := range tests {
		c := c // scopelint mute

		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			buf := bytes.Buffer{}
			logger := log.New(&buf)
			s := sender.New(logger)

			s.SendNotFound(w, c.reason)

			assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
			assert.Equal(t, c.expectedResponse, w.Body.String())
		})
	}
}

func TestSender_SendForbidden(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "uno",
		},
	}

	for _, c := range tests {
		c := c // scopelint mute

		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			buf := bytes.Buffer{}
			logger := log.New(&buf)
			s := sender.New(logger)

			s.SendForbidden(w)

			assert.Equal(t, http.StatusForbidden, w.Result().StatusCode)
		})
	}
}

func TestSender_SendNoContent(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "uno",
		},
	}

	for _, c := range tests {
		c := c // scopelint mute

		if c.name == "" {
			t.Errorf("test case name required!")
			continue
		}

		t.Run(c.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			buf := bytes.Buffer{}
			logger := log.New(&buf)
			s := sender.New(logger)

			s.SendNoContent(w)

			assert.Equal(t, http.StatusNoContent, w.Result().StatusCode)
		})
	}
}
