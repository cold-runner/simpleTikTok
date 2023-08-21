package errno

import (
	"github.com/marmotedu/errors"
	"github.com/novalagung/gubrak"
)

type Err struct {
	HTTP    int
	C       int
	Message string
}

func (e Err) HTTPStatus() int {
	return e.HTTP
}

func (e Err) String() string {
	return e.Message
}

func (e Err) Reference() string {
	return ""
}

func (e Err) Code() int {
	return e.C
}

// 编译检查
var _ errors.Coder = &Err{}

// nolint: unparam
func register(code int, httpStatus int, message string, refs ...string) {
	found, _ := gubrak.Includes([]int{200, 400, 401, 403, 404, 500}, httpStatus)
	if !found {
		panic("http code not in `200, 400, 401, 403, 404, 500`")
	}

	coder := &Err{
		C:    code,
		HTTP: httpStatus,
	}

	errors.MustRegister(coder)
}
