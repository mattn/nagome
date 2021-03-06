package nicolive

import (
	"fmt"
	"runtime"
)

// NicoErrNum is an Enum to represent Nico.Errnum
type NicoErrNum int

// Enum NicoErrNum
const (
	NicoErrOther NicoErrNum = iota - 1
	NicoErrNicoLiveOther
	NicoErrNotLogin
	NicoErrClosed
)

// NicoError is error interface in nicolive
type NicoError interface {
	Code() string
	Description() string
	Where() string
}

// NicoErrStruct is an error struct in nicolive
type NicoErrStruct struct {
	errnum      NicoErrNum
	code        string
	description string
	where       string
}

func (n NicoErrStruct) Error() string {
	if n.description == "" {
		return fmt.Sprintf("[%s] %s", n.where, n.code)
	}
	return fmt.Sprintf("[%s] %s : %s", n.where, n.code, n.description)
}

// ErrorNum returns errorNum
func (n NicoErrStruct) ErrorNum() NicoErrNum {
	return n.errnum
}

// Code returns code
func (n NicoErrStruct) Code() string {
	return n.code
}

// Description returns description
func (n NicoErrStruct) Description() string {
	return n.description
}

// Where returns where
func (n NicoErrStruct) Where() string {
	return n.where
}

// NicoErr returns NicoErr that format as the given info
// and the code position.
func NicoErr(errNum NicoErrNum, code string, description string) NicoError {
	_, file, line, ok := runtime.Caller(1)
	short := file
	if !ok {
		short = "???"
		line = 0
	} else {
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
	}
	where := fmt.Sprintf("%s:%d", short, line)

	return &NicoErrStruct{errNum, code, description, where}
}

// NicoErrFromStdErr returns NicoErr converted from standard error.
func NicoErrFromStdErr(e error) NicoError {
	return NicoErr(NicoErrOther, "standard error", e.Error())
}
