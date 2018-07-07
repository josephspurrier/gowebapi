package testutil

// MockLogger is a mocked logger.
type MockLogger struct {
	FatalfFunc FatalfFuncType
	PrintfFunc PrintfFuncType
}

// FatalfFuncType .
type FatalfFuncType func(format string, v ...interface{})

// FatalfFuncDefault .
var FatalfFuncDefault = func(format string, v ...interface{}) {}

// Fatalf .
func (l *MockLogger) Fatalf(format string, v ...interface{}) {
	if l.FatalfFunc != nil {
		l.FatalfFunc(format, v...)
	}
	FatalfFuncDefault(format, v...)
}

// PrintfFuncType .
type PrintfFuncType func(format string, v ...interface{})

// PrintfFuncDefault .
var PrintfFuncDefault = func(format string, v ...interface{}) {}

// Printf .
func (l *MockLogger) Printf(format string, v ...interface{}) {
	if l.PrintfFunc != nil {
		l.PrintfFunc(format, v...)
	}
	PrintfFuncDefault(format, v...)
}
