package httpresponse

import (
	"github.com/devbackend/goingot/pkg/logger/logger_mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type jsonTestCase struct {
	status           int
	data             interface{}
	expectedStatus   int
	expectedResponse string
}

func TestHTTPResponse_JSON(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logger_mock.NewMockLogger(ctrl)
	logger.EXPECT().Error(gomock.AssignableToTypeOf(""), gomock.Any()).AnyTimes()

	for _, c := range caseProviderJSON() {

		hr := NewSender(logger)

		w := httptest.NewRecorder()

		hr.JSON(w, c.status, c.data)

		assert.Equal(t, c.expectedStatus, w.Result().StatusCode)
		assert.Equal(t, c.expectedResponse, w.Body.String())
	}
}

func caseProviderJSON() []jsonTestCase {
	return []jsonTestCase{
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
			expectedResponse: `Internal error`,
		},
	}
}
