package logger_test

import (
	"app/webapi/pkg/logger"
	"testing"

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

	l.Printf("test", arr[0], arr[1], arr[2])
	assert.Equal(t, true, m.printfCalled)
	assert.Equal(t, arr, m.printfV)
}

func TestFatalf(t *testing.T) {
	m := new(mockLogger)
	l := logger.New(m)

	var arr []interface{}
	arr = append(arr, "A")
	arr = append(arr, 2)
	arr = append(arr, "c")

	l.Fatalf("test", arr[0], arr[1], arr[2])
	assert.Equal(t, true, m.fatalfCalled)
	assert.Equal(t, arr, m.fatalfV)
}
