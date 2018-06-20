package logrequest_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"app/webapi/middleware/logrequest"

	"github.com/stretchr/testify/assert"
)

type MockLogger struct {
	output string
}

func (l *MockLogger) Printf(format string, v ...interface{}) {
	l.output = fmt.Sprintf(format, v...)
}

type MockClock struct{}

func (c *MockClock) Now() time.Time {
	t := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	return t
}

func TestHandler(t *testing.T) {
	ml := new(MockLogger)
	c := new(MockClock)
	lr := logrequest.New()
	lr.SetClock(c)
	lr.SetLog(ml)

	called := false

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		called = true
	})

	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	lr.Handler(mux).ServeHTTP(w, r)

	assert.Equal(t, true, called)
	assert.Equal(t, ml.output, "2009-11-17 08:34:58 PM 192.0.2.1:1234 GET /")
}
