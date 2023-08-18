package main

import (
	UserService "github.com/cold-runner/simpleTikTok/kitex_gen/UserService/userservice"
	"log"
)

func main() {
	svr := UserService.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
