package main

import (
	SocialService "github.com/cold-runner/simpleTikTok/kitex_gen/SocialService/socialservice"
	"github.com/cold-runner/simpleTikTok/service/social"
	"log"
)

func main() {
	svr := SocialService.NewServer(new(social.SocialServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
