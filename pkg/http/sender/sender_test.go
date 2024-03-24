package sender_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

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

		res := w.Result()

		assert.Equal(t, c.expectedStatus, res.StatusCode)
		assert.Equal(t, c.expectedResponse, w.Body.String())

		_ = res.Body.Close()
	}
}

func TestSender_SendOK(t *testing.T) {
	cases := []struct {
		data             interface{}
		expectedResponse string
	}{
		{
			data:             nil,
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

		res := w.Result()

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, c.expectedResponse, w.Body.String())

		require.NoError(t, res.Body.Close())
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
			errors:           nil,
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

		res := w.Result()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Equal(t, c.expectedResponse, w.Body.String())

		_ = res.Body.Close()
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
			err:         nil,
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

			res := w.Result()

			assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
			assert.Contains(t, buf.String(), c.expectedLog)

			require.NoError(t, res.Body.Close())
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
			reason:           "",
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

			res := w.Result()

			assert.Equal(t, http.StatusNotFound, res.StatusCode)
			assert.Equal(t, c.expectedResponse, w.Body.String())

			require.NoError(t, res.Body.Close())
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

			res := w.Result()

			assert.Equal(t, http.StatusForbidden, res.StatusCode)

			require.NoError(t, res.Body.Close())
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
			t.Parallel()

			w := httptest.NewRecorder()

			buf := bytes.Buffer{}
			logger := log.New(&buf)
			s := sender.New(logger)

			s.SendNoContent(w)

			res := w.Result()

			assert.Equal(t, http.StatusNoContent, res.StatusCode)

			require.NoError(t, res.Body.Close())
		})
	}
}
