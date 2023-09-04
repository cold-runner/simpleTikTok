package main

import (
	SocialService "github.com/cold-runner/simpleTikTok/kitex_gen/SocialService/socialservice"
	"log"
)

func main() {
	svr := SocialService.NewServer(new(SocialServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
