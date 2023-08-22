package main

import (
	"errors"
	"fmt"
	_ "github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
)

func main() {
	err := errno.ErrUserAlredyExist

	var e *errno.Errno
	if errors.As(err, &e) {
		fmt.Println("Yes")
		fmt.Println(e)
	}
	fmt.Println("No")
}
