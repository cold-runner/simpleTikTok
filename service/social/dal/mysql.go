package dal

import (
	"context"
	"github.com/cold-runner/simpleTikTok/pkg/dao/model"
	dao "github.com/cold-runner/simpleTikTok/pkg/dao/query"
)

type Mysql struct {
	*dao.Query
}

var (
	MysqlIns *Mysql
)

// 编译器检查
var _ = Dal(&Mysql{})

func (m Mysql) CreateComment(ctx context.Context, comment *model.Comment) error {
	err := m.Transaction(func(tx *dao.Query) error {
		// 创建一条评论
		err := tx.WithContext(ctx).Comment.Create(comment)
		if err != nil {
			return err
		}
		// 更新video表，评论数量+1
		videoTable := tx.Video
		_, err = videoTable.
			WithContext(ctx).
			Where(videoTable.ID.Eq(comment.VedioID)).
			UpdateSimple(videoTable.CommentCount.Add(1))
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (m Mysql) CommentList(ctx context.Context, videoID int64) ([]*model.Comment, error) {
	commentTable := m.Query.Comment
	commentList, err := m.Query.WithContext(ctx).Comment.Where(commentTable.VedioID.Eq(videoID)).Find()
	if err != nil {
		return nil, err
	}
	return commentList, nil
}

func (m Mysql) DeleteComment(ctx context.Context, commentID int64) error {
	err := m.Transaction(func(tx *dao.Query) error {
		commentTable := tx.Comment
		// 查comment表 TODO 所有字段都查询了
		comment, err := commentTable.
			WithContext(ctx).
			Where(commentTable.ID.Eq(commentID)).
			First()
		if err != nil {
			return err
		}
		// 根据comment的videoID，更新video表，评论数量-1
		videoTable := tx.Video
		_, err = videoTable.
			WithContext(ctx).
			Where(videoTable.ID.Eq(comment.VedioID)).
			UpdateSimple(videoTable.CommentCount.Sub(1))
		if err != nil {
			return err
		}

		// 删除评论
		_, err = commentTable.
			WithContext(ctx).
			Delete(comment)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (m Mysql) LikeAction(ctx context.Context, like *model.VideoLike) error {
	err := m.Query.Transaction(func(tx *dao.Query) error {
		// 创建一条点赞记录
		if err := tx.
			WithContext(ctx).
			VideoLike.
			Create(like); err != nil {
			return err
		}

		// 更新视频表，点赞数+1
		videoTable := tx.Video
		if _, err := videoTable.
			WithContext(ctx).
			Where(videoTable.ID.Eq(like.VideoID)).
			UpdateSimple(videoTable.FavoriteCount.Add(1)); err != nil {
			return err
		}

		// 更新用户表，点赞数+1
		userTable := tx.User
		if _, err := userTable.
			WithContext(ctx).
			Where(userTable.ID.Eq(like.UserID)).
			UpdateSimple(userTable.FavoriteCount.Add(1)); err != nil {
			return err
		}
		return nil
	})

	// 事务失败
	if err != nil {
		return err
	}
	// 事务成功
	return nil
}

func (m Mysql) CancelLikeAction(ctx context.Context, userId int64, videoId int64) error {
	err := m.Transaction(func(tx *dao.Query) error {
		// video表视频点赞数-1
		videoTable := tx.Video
		if _, err := videoTable.
			WithContext(ctx).
			Where(videoTable.ID.Eq(videoId)).
			UpdateSimple(videoTable.FavoriteCount.Sub(1)); err != nil {
			return err
		}

		// 用户表喜欢数-1
		userTable := tx.User
		if _, err := userTable.
			WithContext(ctx).
			Where(userTable.ID.Eq(userId)).
			UpdateSimple(userTable.FavoriteCount.Sub(1)); err != nil {
			return err
		}

		// user-vide关系表删除记录
		relationTable := tx.VideoLike
		if _, err := relationTable.
			Unscoped().
			WithContext(ctx).
			Where(
				relationTable.UserID.Eq(userId),
				relationTable.VideoID.Eq(videoId)).
			Delete(); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (m Mysql) LikeList(ctx context.Context, userId int64) ([]*model.Video, error) {
	var videoList []*model.Video
	// 从video-user关系表中搜索出目标用户喜欢的所有视频Id，关联video表返回数据
	relation := m.Query.VideoLike
	videoTable := m.Query.Video

	if err := relation.
		WithContext(ctx).
		Select(videoTable.ALL).
		Where(relation.UserID.Eq(userId)).
		LeftJoin(videoTable, videoTable.ID.EqCol(relation.VideoID)).
		Scan(&videoList); err != nil {
		return nil, err
	}

	return videoList, nil
}
