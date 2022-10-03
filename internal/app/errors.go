package app

import (
	"bytes"
	"fmt"
)

type Error struct {
	Code    string
	Message string
	Op      string
	Err     error
}

const (
	ECONFLICT = "conflict"  // action cannot be performed
	EINTERNAL = "internal"  // internal error
	EINVALID  = "invalid"   // validation failed
	ENOTFOUND = "not_found" // entity does not exist
)

func ErrorCode(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Code != "" {
		return e.Code
	} else if ok && e.Err != nil {
		return ErrorCode(e.Err)
	}
	return EINTERNAL
}

func ErrorMessage(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Message != "" {
		return e.Message
	} else if ok && e.Err != nil {
		return ErrorMessage(e.Err)
	}
	return "An internal error has occurred."
}

func (e *Error) Error() string {
	var buf bytes.Buffer

	if e.Op != "" {
		fmt.Fprintf(&buf, "%s: ", e.Op)
	}

	// assumes that Err cannot coexist with Code or Message on any given error
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != "" {
			fmt.Fprintf(&buf, "<%s> ", e.Code)
		}
		buf.WriteString(e.Message)
	}
	return buf.String()
}
