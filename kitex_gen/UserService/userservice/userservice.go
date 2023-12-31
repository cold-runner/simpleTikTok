// Code generated by Kitex v0.6.2. DO NOT EDIT.

package userservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	BaseResponse "github.com/cold-runner/simpleTikTok/kitex_gen/BaseResponse"
	UserService "github.com/cold-runner/simpleTikTok/kitex_gen/UserService"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*UserService.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Register":                      kitex.NewMethodInfo(registerHandler, newRegisterArgs, newRegisterResult, false),
		"Login":                         kitex.NewMethodInfo(loginHandler, newLoginArgs, newLoginResult, false),
		"GetUserInfo":                   kitex.NewMethodInfo(getUserInfoHandler, newGetUserInfoArgs, newGetUserInfoResult, false),
		"MGetUserInfo":                  kitex.NewMethodInfo(mGetUserInfoHandler, newMGetUserInfoArgs, newMGetUserInfoResult, false),
		"ChangeUserFollowCount":         kitex.NewMethodInfo(changeUserFollowCountHandler, newChangeUserFollowCountArgs, newChangeUserFollowCountResult, false),
		"ChangeUserTotalFavoritedCount": kitex.NewMethodInfo(changeUserTotalFavoritedCountHandler, newChangeUserTotalFavoritedCountArgs, newChangeUserTotalFavoritedCountResult, false),
		"ChangeUserWorkCount":           kitex.NewMethodInfo(changeUserWorkCountHandler, newChangeUserWorkCountArgs, newChangeUserWorkCountResult, false),
		"ChangeUserFavoriteCount":       kitex.NewMethodInfo(changeUserFavoriteCountHandler, newChangeUserFavoriteCountArgs, newChangeUserFavoriteCountResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.6.2",
		Extra:           extra,
	}
	return svcInfo
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(UserService.UserRegisterRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(UserService.UserService).Register(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *RegisterArgs:
		success, err := handler.(UserService.UserService).Register(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RegisterResult)
		realResult.Success = success
	}
	return nil
}
func newRegisterArgs() interface{} {
	return &RegisterArgs{}
}

func newRegisterResult() interface{} {
	return &RegisterResult{}
}

type RegisterArgs struct {
	Req *UserService.UserRegisterRequest
}

func (p *RegisterArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(UserService.UserRegisterRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RegisterArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RegisterArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RegisterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in RegisterArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *RegisterArgs) Unmarshal(in []byte) error {
	msg := new(UserService.UserRegisterRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RegisterArgs_Req_DEFAULT *UserService.UserRegisterRequest

func (p *RegisterArgs) GetReq() *UserService.UserRegisterRequest {
	if !p.IsSetReq() {
		return RegisterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RegisterArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RegisterArgs) GetFirstArgument() interface{} {
	return p.Req
}

type RegisterResult struct {
	Success *UserService.UserRegisterResponse
}

var RegisterResult_Success_DEFAULT *UserService.UserRegisterResponse

func (p *RegisterResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(UserService.UserRegisterResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RegisterResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RegisterResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RegisterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in RegisterResult")
	}
	return proto.Marshal(p.Success)
}

func (p *RegisterResult) Unmarshal(in []byte) error {
	msg := new(UserService.UserRegisterResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RegisterResult) GetSuccess() *UserService.UserRegisterResponse {
	if !p.IsSetSuccess() {
		return RegisterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RegisterResult) SetSuccess(x interface{}) {
	p.Success = x.(*UserService.UserRegisterResponse)
}

func (p *RegisterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RegisterResult) GetResult() interface{} {
	return p.Success
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(UserService.UserLoginRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(UserService.UserService).Login(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *LoginArgs:
		success, err := handler.(UserService.UserService).Login(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*LoginResult)
		realResult.Success = success
	}
	return nil
}
func newLoginArgs() interface{} {
	return &LoginArgs{}
}

func newLoginResult() interface{} {
	return &LoginResult{}
}

type LoginArgs struct {
	Req *UserService.UserLoginRequest
}

func (p *LoginArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(UserService.UserLoginRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *LoginArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *LoginArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *LoginArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in LoginArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *LoginArgs) Unmarshal(in []byte) error {
	msg := new(UserService.UserLoginRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var LoginArgs_Req_DEFAULT *UserService.UserLoginRequest

func (p *LoginArgs) GetReq() *UserService.UserLoginRequest {
	if !p.IsSetReq() {
		return LoginArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *LoginArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *LoginArgs) GetFirstArgument() interface{} {
	return p.Req
}

type LoginResult struct {
	Success *UserService.UserLoginResponse
}

var LoginResult_Success_DEFAULT *UserService.UserLoginResponse

func (p *LoginResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(UserService.UserLoginResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *LoginResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *LoginResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *LoginResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in LoginResult")
	}
	return proto.Marshal(p.Success)
}

func (p *LoginResult) Unmarshal(in []byte) error {
	msg := new(UserService.UserLoginResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *LoginResult) GetSuccess() *UserService.UserLoginResponse {
	if !p.IsSetSuccess() {
		return LoginResult_Success_DEFAULT
	}
	return p.Success
}

func (p *LoginResult) SetSuccess(x interface{}) {
	p.Success = x.(*UserService.UserLoginResponse)
}

func (p *LoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *LoginResult) GetResult() interface{} {
	return p.Success
}

func getUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(UserService.UserInfoRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(UserService.UserService).GetUserInfo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetUserInfoArgs:
		success, err := handler.(UserService.UserService).GetUserInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserInfoResult)
		realResult.Success = success
	}
	return nil
}
func newGetUserInfoArgs() interface{} {
	return &GetUserInfoArgs{}
}

func newGetUserInfoResult() interface{} {
	return &GetUserInfoResult{}
}

type GetUserInfoArgs struct {
	Req *UserService.UserInfoRequest
}

func (p *GetUserInfoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(UserService.UserInfoRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetUserInfoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetUserInfoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetUserInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUserInfoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserInfoArgs) Unmarshal(in []byte) error {
	msg := new(UserService.UserInfoRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserInfoArgs_Req_DEFAULT *UserService.UserInfoRequest

func (p *GetUserInfoArgs) GetReq() *UserService.UserInfoRequest {
	if !p.IsSetReq() {
		return GetUserInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetUserInfoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetUserInfoResult struct {
	Success *UserService.UserInfoResponse
}

var GetUserInfoResult_Success_DEFAULT *UserService.UserInfoResponse

func (p *GetUserInfoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(UserService.UserInfoResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetUserInfoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetUserInfoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetUserInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUserInfoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserInfoResult) Unmarshal(in []byte) error {
	msg := new(UserService.UserInfoResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserInfoResult) GetSuccess() *UserService.UserInfoResponse {
	if !p.IsSetSuccess() {
		return GetUserInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*UserService.UserInfoResponse)
}

func (p *GetUserInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserInfoResult) GetResult() interface{} {
	return p.Success
}

func mGetUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(UserService.MGetUserInfoRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(UserService.UserService).MGetUserInfo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *MGetUserInfoArgs:
		success, err := handler.(UserService.UserService).MGetUserInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*MGetUserInfoResult)
		realResult.Success = success
	}
	return nil
}
func newMGetUserInfoArgs() interface{} {
	return &MGetUserInfoArgs{}
}

func newMGetUserInfoResult() interface{} {
	return &MGetUserInfoResult{}
}

type MGetUserInfoArgs struct {
	Req *UserService.MGetUserInfoRequest
}

func (p *MGetUserInfoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(UserService.MGetUserInfoRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *MGetUserInfoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *MGetUserInfoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *MGetUserInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MGetUserInfoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *MGetUserInfoArgs) Unmarshal(in []byte) error {
	msg := new(UserService.MGetUserInfoRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MGetUserInfoArgs_Req_DEFAULT *UserService.MGetUserInfoRequest

func (p *MGetUserInfoArgs) GetReq() *UserService.MGetUserInfoRequest {
	if !p.IsSetReq() {
		return MGetUserInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MGetUserInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *MGetUserInfoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type MGetUserInfoResult struct {
	Success *UserService.MGetUserInfoResponse
}

var MGetUserInfoResult_Success_DEFAULT *UserService.MGetUserInfoResponse

func (p *MGetUserInfoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(UserService.MGetUserInfoResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *MGetUserInfoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *MGetUserInfoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *MGetUserInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MGetUserInfoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *MGetUserInfoResult) Unmarshal(in []byte) error {
	msg := new(UserService.MGetUserInfoResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MGetUserInfoResult) GetSuccess() *UserService.MGetUserInfoResponse {
	if !p.IsSetSuccess() {
		return MGetUserInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MGetUserInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*UserService.MGetUserInfoResponse)
}

func (p *MGetUserInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MGetUserInfoResult) GetResult() interface{} {
	return p.Success
}

func changeUserFollowCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(UserService.ChangeUserFollowCountRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(UserService.UserService).ChangeUserFollowCount(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *ChangeUserFollowCountArgs:
		success, err := handler.(UserService.UserService).ChangeUserFollowCount(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ChangeUserFollowCountResult)
		realResult.Success = success
	}
	return nil
}
func newChangeUserFollowCountArgs() interface{} {
	return &ChangeUserFollowCountArgs{}
}

func newChangeUserFollowCountResult() interface{} {
	return &ChangeUserFollowCountResult{}
}

type ChangeUserFollowCountArgs struct {
	Req *UserService.ChangeUserFollowCountRequest
}

func (p *ChangeUserFollowCountArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(UserService.ChangeUserFollowCountRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ChangeUserFollowCountArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ChangeUserFollowCountArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ChangeUserFollowCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ChangeUserFollowCountArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *ChangeUserFollowCountArgs) Unmarshal(in []byte) error {
	msg := new(UserService.ChangeUserFollowCountRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ChangeUserFollowCountArgs_Req_DEFAULT *UserService.ChangeUserFollowCountRequest

func (p *ChangeUserFollowCountArgs) GetReq() *UserService.ChangeUserFollowCountRequest {
	if !p.IsSetReq() {
		return ChangeUserFollowCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChangeUserFollowCountArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ChangeUserFollowCountArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ChangeUserFollowCountResult struct {
	Success *BaseResponse.BaseResp
}

var ChangeUserFollowCountResult_Success_DEFAULT *BaseResponse.BaseResp

func (p *ChangeUserFollowCountResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(BaseResponse.BaseResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ChangeUserFollowCountResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ChangeUserFollowCountResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ChangeUserFollowCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ChangeUserFollowCountResult")
	}
	return proto.Marshal(p.Success)
}

func (p *ChangeUserFollowCountResult) Unmarshal(in []byte) error {
	msg := new(BaseResponse.BaseResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChangeUserFollowCountResult) GetSuccess() *BaseResponse.BaseResp {
	if !p.IsSetSuccess() {
		return ChangeUserFollowCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChangeUserFollowCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*BaseResponse.BaseResp)
}

func (p *ChangeUserFollowCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChangeUserFollowCountResult) GetResult() interface{} {
	return p.Success
}

func changeUserTotalFavoritedCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(UserService.ChangeUserTotalFavoritedCountRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(UserService.UserService).ChangeUserTotalFavoritedCount(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *ChangeUserTotalFavoritedCountArgs:
		success, err := handler.(UserService.UserService).ChangeUserTotalFavoritedCount(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ChangeUserTotalFavoritedCountResult)
		realResult.Success = success
	}
	return nil
}
func newChangeUserTotalFavoritedCountArgs() interface{} {
	return &ChangeUserTotalFavoritedCountArgs{}
}

func newChangeUserTotalFavoritedCountResult() interface{} {
	return &ChangeUserTotalFavoritedCountResult{}
}

type ChangeUserTotalFavoritedCountArgs struct {
	Req *UserService.ChangeUserTotalFavoritedCountRequest
}

func (p *ChangeUserTotalFavoritedCountArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(UserService.ChangeUserTotalFavoritedCountRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ChangeUserTotalFavoritedCountArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ChangeUserTotalFavoritedCountArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ChangeUserTotalFavoritedCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ChangeUserTotalFavoritedCountArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *ChangeUserTotalFavoritedCountArgs) Unmarshal(in []byte) error {
	msg := new(UserService.ChangeUserTotalFavoritedCountRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ChangeUserTotalFavoritedCountArgs_Req_DEFAULT *UserService.ChangeUserTotalFavoritedCountRequest

func (p *ChangeUserTotalFavoritedCountArgs) GetReq() *UserService.ChangeUserTotalFavoritedCountRequest {
	if !p.IsSetReq() {
		return ChangeUserTotalFavoritedCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChangeUserTotalFavoritedCountArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ChangeUserTotalFavoritedCountArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ChangeUserTotalFavoritedCountResult struct {
	Success *UserService.UpdateUserInfoResponse
}

var ChangeUserTotalFavoritedCountResult_Success_DEFAULT *UserService.UpdateUserInfoResponse

func (p *ChangeUserTotalFavoritedCountResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(UserService.UpdateUserInfoResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ChangeUserTotalFavoritedCountResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ChangeUserTotalFavoritedCountResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ChangeUserTotalFavoritedCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ChangeUserTotalFavoritedCountResult")
	}
	return proto.Marshal(p.Success)
}

func (p *ChangeUserTotalFavoritedCountResult) Unmarshal(in []byte) error {
	msg := new(UserService.UpdateUserInfoResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChangeUserTotalFavoritedCountResult) GetSuccess() *UserService.UpdateUserInfoResponse {
	if !p.IsSetSuccess() {
		return ChangeUserTotalFavoritedCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChangeUserTotalFavoritedCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*UserService.UpdateUserInfoResponse)
}

func (p *ChangeUserTotalFavoritedCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChangeUserTotalFavoritedCountResult) GetResult() interface{} {
	return p.Success
}

func changeUserWorkCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(UserService.ChangeUserWorkCountRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(UserService.UserService).ChangeUserWorkCount(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *ChangeUserWorkCountArgs:
		success, err := handler.(UserService.UserService).ChangeUserWorkCount(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ChangeUserWorkCountResult)
		realResult.Success = success
	}
	return nil
}
func newChangeUserWorkCountArgs() interface{} {
	return &ChangeUserWorkCountArgs{}
}

func newChangeUserWorkCountResult() interface{} {
	return &ChangeUserWorkCountResult{}
}

type ChangeUserWorkCountArgs struct {
	Req *UserService.ChangeUserWorkCountRequest
}

func (p *ChangeUserWorkCountArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(UserService.ChangeUserWorkCountRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ChangeUserWorkCountArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ChangeUserWorkCountArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ChangeUserWorkCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ChangeUserWorkCountArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *ChangeUserWorkCountArgs) Unmarshal(in []byte) error {
	msg := new(UserService.ChangeUserWorkCountRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ChangeUserWorkCountArgs_Req_DEFAULT *UserService.ChangeUserWorkCountRequest

func (p *ChangeUserWorkCountArgs) GetReq() *UserService.ChangeUserWorkCountRequest {
	if !p.IsSetReq() {
		return ChangeUserWorkCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChangeUserWorkCountArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ChangeUserWorkCountArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ChangeUserWorkCountResult struct {
	Success *UserService.UpdateUserInfoResponse
}

var ChangeUserWorkCountResult_Success_DEFAULT *UserService.UpdateUserInfoResponse

func (p *ChangeUserWorkCountResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(UserService.UpdateUserInfoResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ChangeUserWorkCountResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ChangeUserWorkCountResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ChangeUserWorkCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ChangeUserWorkCountResult")
	}
	return proto.Marshal(p.Success)
}

func (p *ChangeUserWorkCountResult) Unmarshal(in []byte) error {
	msg := new(UserService.UpdateUserInfoResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChangeUserWorkCountResult) GetSuccess() *UserService.UpdateUserInfoResponse {
	if !p.IsSetSuccess() {
		return ChangeUserWorkCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChangeUserWorkCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*UserService.UpdateUserInfoResponse)
}

func (p *ChangeUserWorkCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChangeUserWorkCountResult) GetResult() interface{} {
	return p.Success
}

func changeUserFavoriteCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(UserService.ChangeUserFavoriteCountRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(UserService.UserService).ChangeUserFavoriteCount(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *ChangeUserFavoriteCountArgs:
		success, err := handler.(UserService.UserService).ChangeUserFavoriteCount(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ChangeUserFavoriteCountResult)
		realResult.Success = success
	}
	return nil
}
func newChangeUserFavoriteCountArgs() interface{} {
	return &ChangeUserFavoriteCountArgs{}
}

func newChangeUserFavoriteCountResult() interface{} {
	return &ChangeUserFavoriteCountResult{}
}

type ChangeUserFavoriteCountArgs struct {
	Req *UserService.ChangeUserFavoriteCountRequest
}

func (p *ChangeUserFavoriteCountArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(UserService.ChangeUserFavoriteCountRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ChangeUserFavoriteCountArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ChangeUserFavoriteCountArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ChangeUserFavoriteCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ChangeUserFavoriteCountArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *ChangeUserFavoriteCountArgs) Unmarshal(in []byte) error {
	msg := new(UserService.ChangeUserFavoriteCountRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ChangeUserFavoriteCountArgs_Req_DEFAULT *UserService.ChangeUserFavoriteCountRequest

func (p *ChangeUserFavoriteCountArgs) GetReq() *UserService.ChangeUserFavoriteCountRequest {
	if !p.IsSetReq() {
		return ChangeUserFavoriteCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChangeUserFavoriteCountArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ChangeUserFavoriteCountArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ChangeUserFavoriteCountResult struct {
	Success *UserService.UpdateUserInfoResponse
}

var ChangeUserFavoriteCountResult_Success_DEFAULT *UserService.UpdateUserInfoResponse

func (p *ChangeUserFavoriteCountResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(UserService.UpdateUserInfoResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ChangeUserFavoriteCountResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ChangeUserFavoriteCountResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ChangeUserFavoriteCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ChangeUserFavoriteCountResult")
	}
	return proto.Marshal(p.Success)
}

func (p *ChangeUserFavoriteCountResult) Unmarshal(in []byte) error {
	msg := new(UserService.UpdateUserInfoResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChangeUserFavoriteCountResult) GetSuccess() *UserService.UpdateUserInfoResponse {
	if !p.IsSetSuccess() {
		return ChangeUserFavoriteCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChangeUserFavoriteCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*UserService.UpdateUserInfoResponse)
}

func (p *ChangeUserFavoriteCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChangeUserFavoriteCountResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, Req *UserService.UserRegisterRequest) (r *UserService.UserRegisterResponse, err error) {
	var _args RegisterArgs
	_args.Req = Req
	var _result RegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, Req *UserService.UserLoginRequest) (r *UserService.UserLoginResponse, err error) {
	var _args LoginArgs
	_args.Req = Req
	var _result LoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserInfo(ctx context.Context, Req *UserService.UserInfoRequest) (r *UserService.UserInfoResponse, err error) {
	var _args GetUserInfoArgs
	_args.Req = Req
	var _result GetUserInfoResult
	if err = p.c.Call(ctx, "GetUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MGetUserInfo(ctx context.Context, Req *UserService.MGetUserInfoRequest) (r *UserService.MGetUserInfoResponse, err error) {
	var _args MGetUserInfoArgs
	_args.Req = Req
	var _result MGetUserInfoResult
	if err = p.c.Call(ctx, "MGetUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChangeUserFollowCount(ctx context.Context, Req *UserService.ChangeUserFollowCountRequest) (r *BaseResponse.BaseResp, err error) {
	var _args ChangeUserFollowCountArgs
	_args.Req = Req
	var _result ChangeUserFollowCountResult
	if err = p.c.Call(ctx, "ChangeUserFollowCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChangeUserTotalFavoritedCount(ctx context.Context, Req *UserService.ChangeUserTotalFavoritedCountRequest) (r *UserService.UpdateUserInfoResponse, err error) {
	var _args ChangeUserTotalFavoritedCountArgs
	_args.Req = Req
	var _result ChangeUserTotalFavoritedCountResult
	if err = p.c.Call(ctx, "ChangeUserTotalFavoritedCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChangeUserWorkCount(ctx context.Context, Req *UserService.ChangeUserWorkCountRequest) (r *UserService.UpdateUserInfoResponse, err error) {
	var _args ChangeUserWorkCountArgs
	_args.Req = Req
	var _result ChangeUserWorkCountResult
	if err = p.c.Call(ctx, "ChangeUserWorkCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChangeUserFavoriteCount(ctx context.Context, Req *UserService.ChangeUserFavoriteCountRequest) (r *UserService.UpdateUserInfoResponse, err error) {
	var _args ChangeUserFavoriteCountArgs
	_args.Req = Req
	var _result ChangeUserFavoriteCountResult
	if err = p.c.Call(ctx, "ChangeUserFavoriteCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
