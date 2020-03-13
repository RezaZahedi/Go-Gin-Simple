package rest_service

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestNewFiboServiceParallel(t *testing.T) {

	t.Parallel()
	router := setupRouter()
	endpointUrl := "/"
	method := "POST"

	tests := []struct {
		name           string
		args           string
		want           string
		wantStatusCode int
	}{
		{"1", "1", "\"1\"", 200},
		{"2", "2", "\"1\"", 200},
		{"8", "8", "\"21\"", 200},
		{"70", "70", "\"190392490709135\"", 200},
		{"10", "10", "\"55\"", 200},
		{"empty", "", "", 400},
		{"negative", "-10", "", 400},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			w := httptest.NewRecorder()
			form := url.Values{}
			form.Add("number", test.args)
			req, _ := http.NewRequest(method, endpointUrl, nil)
			req.PostForm = form
			router.ServeHTTP(w, req)

			assert.Equal(t, test.wantStatusCode, w.Code)
			assert.Equal(t, test.want, strings.TrimSpace(w.Body.String()))
		})
	}
}
