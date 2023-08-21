package social

import (
	"context"
	"github.com/cold-runner/simpleTikTok/kitex_gen/SocialService"
	"reflect"
	"testing"
)

func TestSocialServiceImpl_CommentAction(t *testing.T) {
	type args struct {
		ctx context.Context
		req *SocialService.CommentActionRequest
	}
	var tests []struct {
		name     string
		args     args
		wantResp *SocialService.CommentActionResponse
		wantErr  bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ServiceImpl{}
			gotResp, err := s.CommentAction(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateComment() errno = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("CreateComment() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
