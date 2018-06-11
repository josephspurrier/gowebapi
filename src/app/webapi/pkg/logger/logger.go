package logger

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// New returns a new instance of a logger.
func New() *Logger {
	l := log.New(os.Stderr, "", log.LstdFlags)
	return &Logger{
		out: l,
	}
}

// Logger will output to a writer.
type Logger struct {
	out *log.Logger
}

// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.out.Fatalf(format, v...)
}

// Printf calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Printf(format string, v ...interface{}) {
	l.out.Printf(format, v...)
}

// ControllerError will output the request, file, line number, and error.
func (l *Logger) ControllerError(r *http.Request, err error, a ...interface{}) {
	// Get the error from the previous function.
	_, fn, line, _ := runtime.Caller(1)
	f := filepath.Base(fn)
	p := filepath.Base(filepath.Dir(fn))

	out := fmt.Sprintf("%v %v | %s:%d | %v", r.Method, r.URL.Path, p+"/"+f, line, err)
	if len(a) > 0 {
		parts := make([]string, 0)
		for i := 0; i < len(a); i++ {
			parts = append(parts, fmt.Sprint(a[i]))
		}

		out = fmt.Sprintf("%v %v | %s:%d | %v | %v", r.Method, r.URL.Path, p+"/"+f, line, err, strings.Join(parts, ","))
	}

	l.out.Println(out)
}
