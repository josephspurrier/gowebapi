// Package logger standardizes the logging functions available to your team.
package logger

// ILog provides logging capabilities.
type ILog interface {
	Fatalf(format string, v ...interface{})
	Printf(format string, v ...interface{})
}

// New returns a new instance of a logger.
func New(i ILog) *Logger {
	return &Logger{
		out: i,
	}
}

// Logger will output to a writer.
type Logger struct {
	out ILog
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
