// Code generated by goctl. DO NOT EDIT.

package videoModel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	videoFieldNames          = builder.RawFieldNames(&Video{})
	videoRows                = strings.Join(videoFieldNames, ",")
	videoRowsExpectAutoSet   = strings.Join(stringx.Remove(videoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	videoRowsWithPlaceHolder = strings.Join(stringx.Remove(videoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTiktokVideoIdPrefix = "cache:tiktok:video:id:"
)

type (
	videoModel interface {
		Insert(ctx context.Context, data *Video) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Video, error)
		FindManyByTime(ctx context.Context, time int64, num int64) ([]*Video, error)
		FindAllByUserId(ctx context.Context, userId string) ([]*Video, error)
		Update(ctx context.Context, data *Video) error
		Delete(ctx context.Context, id int64) error
	}

	defaultVideoModel struct {
		sqlc.CachedConn
		table string
	}

	Video struct {
		Id           int64  `db:"id"`
		AuthorId     int64  `db:"author_id"`
		PlayUrl      string `db:"play_url"`
		CoverUrl     string `db:"cover_url"`
		LikeCount    int64  `db:"like_count"`
		CommentCount int64  `db:"comment_count"`
		Time         int64  `db:"time"`
		Title        string `db:"title"`
		Removed      int64  `db:"removed"`
		Type         int64  `db:"type"`
	}
)

func newVideoModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultVideoModel {
	return &defaultVideoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`video`",
	}
}

func (m *defaultVideoModel) Delete(ctx context.Context, id int64) error {
	tiktokVideoIdKey := fmt.Sprintf("%s%v", cacheTiktokVideoIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, tiktokVideoIdKey)
	return err
}

func (m *defaultVideoModel) FindOne(ctx context.Context, id int64) (*Video, error) {
	tiktokVideoIdKey := fmt.Sprintf("%s%v", cacheTiktokVideoIdPrefix, id)
	var resp Video
	err := m.QueryRowCtx(ctx, &resp, tiktokVideoIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", videoRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultVideoModel) Insert(ctx context.Context, data *Video) (sql.Result, error) {
	tiktokVideoIdKey := fmt.Sprintf("%s%v", cacheTiktokVideoIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, videoRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.AuthorId, data.PlayUrl, data.CoverUrl, data.LikeCount, data.CommentCount, data.Time, data.Title, data.Removed, data.Type)
	}, tiktokVideoIdKey)
	return ret, err
}

func (m *defaultVideoModel) Update(ctx context.Context, data *Video) error {
	tiktokVideoIdKey := fmt.Sprintf("%s%v", cacheTiktokVideoIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, videoRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.AuthorId, data.PlayUrl, data.CoverUrl, data.LikeCount, data.CommentCount, data.Time, data.Title, data.Removed, data.Type, data.Id)
	}, tiktokVideoIdKey)
	return err
}

func (m *defaultVideoModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTiktokVideoIdPrefix, primary)
}

func (m *defaultVideoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", videoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultVideoModel) tableName() string {
	return m.table
}
