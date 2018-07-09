package logger_test

import (
	"os"
	"testing"

	"app/webapi/pkg/logger"

	"github.com/stretchr/testify/assert"
)

type mockLogger struct {
	printfCalled bool
	printfFormat string
	printfV      []interface{}

	fatalfCalled bool
	fatalfFormat string
	fatalfV      []interface{}
}

func (l *mockLogger) Fatalf(format string, v ...interface{}) {
	l.fatalfCalled = true
	l.fatalfFormat = format
	l.fatalfV = v
}

func (l *mockLogger) Printf(format string, v ...interface{}) {
	l.printfCalled = true
	l.printfFormat = format
	l.printfV = v
}
func TestPrintf(t *testing.T) {
	m := new(mockLogger)
	l := logger.New(m)

	var arr []interface{}
	arr = append(arr, "A")
	arr = append(arr, 2)
	arr = append(arr, "c")

	l.Printf("test %v %v %v", arr[0], arr[1], arr[2])
	assert.Equal(t, true, m.printfCalled)
	assert.Equal(t, arr, m.printfV)

	// Clear the logger.
	m = new(mockLogger)

	os.Setenv("WEBAPI_LOG_LEVEL", "none")
	l.Printf("test %v %v %v", arr[0], arr[1], arr[2])
	assert.Equal(t, false, m.printfCalled)
	assert.Equal(t, "", m.printfFormat)
	os.Unsetenv("WEBAPI_LOG_LEVEL")
}

func TestFatalf(t *testing.T) {
	m := new(mockLogger)
	l := logger.New(m)

	var arr []interface{}
	arr = append(arr, "A")
	arr = append(arr, 2)
	arr = append(arr, "c")

	l.Fatalf("test %v %v %v", arr[0], arr[1], arr[2])
	assert.Equal(t, true, m.fatalfCalled)
	assert.Equal(t, arr, m.fatalfV)

	// Clear the logger.
	m = new(mockLogger)

	os.Setenv("WEBAPI_LOG_LEVEL", "none")
	l.Fatalf("test %v %v %v", arr[0], arr[1], arr[2])
	assert.Equal(t, false, m.fatalfCalled)
	assert.Equal(t, "", m.fatalfFormat)
	os.Unsetenv("WEBAPI_LOG_LEVEL")
}
