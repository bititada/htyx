package errno

import (
	"fmt"
)

type Errno struct {
	Code    int
	Message string
}

func (err *Errno) Error() string {
	return err.Message
}

type LogErr struct {
	Code    int
	Message string
	Err     error
}

func NewLogErr(errno *Errno, err error) *LogErr {
	return &LogErr{Code: errno.Code, Message: errno.Message, Err: err}
}
func (err *LogErr) Error() string {
	return fmt.Sprintf("Err - code:%d,message:%s,error:%s", err.Code, err.Message, err.Err)
}
func (err *LogErr) Add(message string) error {
	err.Message += " " + message
	return err
}
func (err *LogErr) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}
	switch typed := err.(type) {
	case *LogErr:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:

	}
	return InternalServerError.Code, err.Error()
}
func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFound.Code
}
