package main

import (
	_ "github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
)

type str1 struct {
	UserId     int64 `json:"user_id"`
	ToUserId   int64 `json:"to_user_id"`
	ActionType int32 `json:"action_type"`
}

type str2 struct {
	UserId   int64 `json:"user_id"`
	ToUserId int64 `json:"to_user_id"`
}

func main() {

	//data, _ := json.Marshal(&str1{
	//	UserId:     1,
	//	ToUserId:   2,
	//	ActionType: 3,
	//})
	//var req interface{}
	//err := json.Unmarshal(data, &req)
	//if err != nil {
	//	fmt.Println(err)
	////}
	//encryption.Encrypt()
	//fmt.Println(req)
}

//
//type FavoriteService struct{}
//
//func NewFavoriteService() *FavoriteService {
//	return &FavoriteService{}
//}
//
//func (s *FavoriteService) FavoriteAction(videoID int64, userID int64) {
//	dtmcli.Logf("beginning a dtm saga business")
//	// 初始化DTM客户端
//	//DtmServer := "http://localhost:36789/api/dtmsvr"
//	//dtm := dtmcli.NewSaga(DtmServer, "1").
//	//	Add()
//	//dtmcli.
//	//// 创建Saga
//	//
//	//// 提交Saga
//	//err := saga.Submit()
//	//if err != nil {
//	//	log.Fatalf("saga submit failed: %v", err)
//	//}
//	//
//	dtmgrpc.NewSagaGrpc().Add()
//}
