package v1

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
	"github.com/cold-runner/simpleTikTok/service/social"
	"github.com/cold-runner/simpleTikTok/service/social/dal"
	"github.com/cold-runner/simpleTikTok/service/social/rpc"
	"github.com/cold-runner/simpleTikTok/service/social/rpc/mocks"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
)

func TestCommentActionService_CommentAction(t *testing.T) {
	// mock用户服务接口
	ctrl := gomock.NewController(t)
	m := mocks.NewMockUserRpc(ctrl)
	m.
		EXPECT().
		GetUserInfo(gomock.Eq(int64(1))).
		Return(&SocialService.User{
			Id:            1,
			Name:          "金天翊",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
			Avatar:        "www.carrol-stamm.co",
		}, nil).
		Times(1)
	m.
		EXPECT().
		GetUserInfo(gomock.Eq(2)).
		Return(&SocialService.User{
			Id:            2,
			Name:          "孟嘉熙",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
			Avatar:        "www.heike-swaniawski.net",
		}, nil).
		Times(0)

	type fields struct {
		ctx     context.Context
		Dal     dal.Dal
		UserRpc rpc.UserRpc
	}
	type args struct {
		ctx context.Context
		req *SocialService.CommentActionRequest
	}
	social.Init()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SocialService.Comment
		wantErr bool
	}{
		{
			name: "id=1",
			fields: fields{
				ctx:     nil,
				Dal:     dal.MysqlIns,
				UserRpc: m,
			},
			args: args{
				ctx: nil,
				req: &SocialService.CommentActionRequest{
					UserId:      1,
					VideoId:     1,
					ActionType:  1,
					CommentText: "some comments",
					CommentId:   233,
				},
			},
			want: &SocialService.Comment{
				Id: 0,
				User: &SocialService.User{
					Id:            int64(1),
					Name:          "金天翊",
					FollowCount:   0,
					FollowerCount: 0,
					IsFollow:      false,
					Avatar:        "www.carrol-stamm.co",
				},
				Content:    "some comments",
				CreateDate: "",
			}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CommentActionService{
				ctx:     tt.fields.ctx,
				Dal:     tt.fields.Dal,
				UserRpc: tt.fields.UserRpc,
			}
			got, err := s.CommentAction(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommentAction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if deepEqualCommentResp(tt.want, got) {
				t.Errorf("CommentAction() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func deepEqualCommentResp(want *SocialService.Comment, got *SocialService.Comment) bool {
	if want.Id != got.Id {
		return false
	}
	if want.Content != got.Content {
		return false
	}
	if !reflect.DeepEqual(want.User, got.User) {
		return false
	}
	return true
}
