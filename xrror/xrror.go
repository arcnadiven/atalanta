package xrror

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	std_time_format = `2006/01/02 15:04:05`
)

var (
	default_code        = 0
	default_file_length = -1
	default_stack_depth = 2
)

type Xrror struct {
	Code  int
	Time  string
	Stack string
	Err   string
}

func (x *Xrror) Error() string {
	if x.Code == default_code {
		return fmt.Sprintf(`%s [%s] %s`, x.Time, x.Stack, x.Err)
	}
	return fmt.Sprintf(`(%d) %s [%s] %s`, x.Code, x.Time, x.Stack, x.Err)
}

func SetFileLength(n int) {
	default_file_length = n
}

func SetStackDepth(n int) {
	default_stack_depth = n
}

func String(str string) error {
	return genXrror(default_code, default_stack_depth, default_file_length, str)
}

func New(err error) error {
	return genXrror(default_code, default_stack_depth, default_file_length, err.Error())
}

func NewWithCode(code int, err error) error {
	return genXrror(code, default_stack_depth, default_file_length, err.Error())
}

func NewWithDepth(depth int, err error) error {
	return genXrror(default_code, depth, default_file_length, err.Error())
}

func NewWithFileLen(fl int, err error) error {
	return genXrror(default_code, default_stack_depth, fl, err.Error())
}

func genXrror(code, depth, fl int, str string) error {
	_, file, line, ok := runtime.Caller(depth)
	if !ok {
		file = `unknown file`
		line = 0
	}
	if fl > 0 {
		pathList := strings.Split(file, string(os.PathSeparator))
		if len(pathList) > fl {
			file = strings.Join(pathList[len(pathList)-fl:], string(os.PathSeparator))
		}
	}
	return &Xrror{
		Code:  code,
		Time:  time.Now().Format(std_time_format),
		Stack: file + `:` + strconv.Itoa(line),
		Err:   str,
	}
}
