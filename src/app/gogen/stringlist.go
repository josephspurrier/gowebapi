package gogen

import (
	"strings"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// StringList is a string array.
type StringList []string

// NewStringList accepts one or more strings as arguments.
func NewStringList(s kingpin.Settings) (target *StringList) {
	target = new(StringList)
	s.SetValue((*StringList)(target))
	return
}

// Set appends the string to the list.
func (i *StringList) Set(value string) error {
	*i = append(*i, value)
	return nil
}

// String returns the list.
func (i *StringList) String() string {
	return strings.Join(*i, " ")
}

// IsCumulative allows more than one value to be passed.
func (i *StringList) IsCumulative() bool {
	return true
}
