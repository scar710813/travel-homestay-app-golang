// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"golodge/common/globalkey"
	"golodge/deploy/script/mysql/genModel"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	historyFieldNames          = builder.RawFieldNames(&History{})
	historyRows                = strings.Join(historyFieldNames, ",")
	historyRowsExpectAutoSet   = strings.Join(stringx.Remove(historyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`last_browsing_time`", "`updated_at`"), ",")
	historyRowsWithPlaceHolder = strings.Join(stringx.Remove(historyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`last_browsing_time`", "`updated_at`"), "=?,") + "=?"

	cacheLooklookTravelHistoryIdPrefix = "cache:looklookTravel:history:id:"
)

type (
	historyModel interface {
		Insert(ctx context.Context, data *History) (sql.Result, error)
		SelectBuilder() squirrel.SelectBuilder
		FindOne(ctx context.Context, historyId int64) (*History, error)
		FindOneByUserId(ctx context.Context, userId int64) (*History, error)
		FindOneByHomestayIdAndUserId(ctx context.Context, homestayId, userId int64) (*History, error)
		FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*History, error)
		Update(ctx context.Context, data *History) error
		Delete(ctx context.Context, historyId int64) error
		DeleteSoft(ctx context.Context, session sqlx.Session, data *History) error
		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *History) error
		DeleteAll(ctx context.Context, userId int64) error
	}

	defaultHistoryModel struct {
		sqlc.CachedConn
		table string
	}

	//UpdateTime         time.Time `db:"update_time"`
	History struct {
		Id                 int64     `db:"id"`
		CreateTime         time.Time `db:"create_time"`
		LastBrowsingTime   time.Time `db:"last_browsing_time"`
		Title              string    `db:"title"`
		Cover              string    `db:"cover"`
		Intro              string    `db:"intro"`
		Location           string    `db:"location"`
		PriceBefore        int64     `db:"price_before"`
		PriceAfter         int64     `db:"price_after"`
		RowState           int64     `db:"row_state"`
		HomestayBusinessId int64     `db:"homestay_business_id"`
		RatingStars        float32   `db:"rating_stars"`
		UserId             int64     `db:"user_id"`
		HomestayId         int64     `db:"homestay_id"`
		DelState           int64     `db:"del_state"`
		Version            int64     `db:"version"`
		DeleteTime         time.Time `db:"delete_time"`
	}
)

func newHistoryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultHistoryModel {
	return &defaultHistoryModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`history`",
	}
}

func (m *defaultHistoryModel) Delete(ctx context.Context, historyId int64) error {
	looklookTravelHistoryIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHistoryIdPrefix, historyId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		// bug: 这里写的是historyId, 而数据库里根本没有historyId, 应该是id
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, historyId)
	}, looklookTravelHistoryIdKey)
	return err
}

func (m *defaultHistoryModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, data *History) error {

	// bug 这里误注释了
	oldVersion := data.Version
	data.Version += 1

	var sqlResult sql.Result
	var err error

	looklookTravelHistoryIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHistoryIdPrefix, data.Id)
	sqlResult, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and version = ? ", m.table, historyRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Title, data.Cover, data.Intro, data.Location, data.PriceBefore, data.PriceAfter, data.RowState, data.HomestayBusinessId, data.RatingStars, data.UserId, data.HomestayId, data.DelState, data.Version, data.DeleteTime, data.Id, oldVersion)
		}
		// bug最后得带data.id和oldVersion，不然会报参数无法满足
		// ctx, query, data.DeleteTime, data.DelState, data.Version, data.Title, data.SubTitle, data.Banner, data.Info, data.PeopleNum, data.HomestayBusinessId, data.UserId, data.RowState, data.RowType, data.FoodInfo, data.FoodPrice, data.HomestayPrice, data.MarketHomestayPrice, data.Id, oldVersion
		return conn.ExecCtx(ctx, query, data.Title, data.Cover, data.Intro, data.Location, data.PriceBefore, data.PriceAfter, data.RowState, data.HomestayBusinessId, data.RatingStars, data.UserId, data.HomestayId, data.DelState, data.Version, data.DeleteTime, data.Id, oldVersion)
	}, looklookTravelHistoryIdKey)
	if err != nil {
		return err
	}
	updateCount, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if updateCount == 0 {
		return genModel.ErrNoRowsUpdate
	}

	return nil
}

func (m *defaultHistoryModel) DeleteSoft(ctx context.Context, session sqlx.Session, data *History) error {
	data.DelState = globalkey.DelStateYes
	data.DeleteTime = time.Now()
	if err := m.UpdateWithVersion(ctx, session, data); err != nil {
		return errors.Wrapf(errors.New("delete soft failed "), "HistoryModel delete err : %+v", err)
	}
	return nil
}

func (m *defaultHistoryModel) DeleteAll(ctx context.Context, userId int64) error {
	looklookTravelHistoryIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHistoryIdPrefix, userId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, userId)
	}, looklookTravelHistoryIdKey)
	return err
}

//func (m *defaultHistoryModel) DeleteAllSoft(ctx context.Context, session sqlx.Session, data *History) error {
//	looklookTravelHistoryIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHistoryIdPrefix, userId)
//	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
//		query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
//		return conn.ExecCtx(ctx, query, userId)
//	}, looklookTravelHistoryIdKey)
//	return err
//}

func (m *defaultHistoryModel) FindOne(ctx context.Context, historyId int64) (*History, error) {
	looklookTravelHistoryIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHistoryIdPrefix, historyId)
	var resp History
	err := m.QueryRowCtx(ctx, &resp, looklookTravelHistoryIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		// bug: 这里不能用history_id, 要用id
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", historyRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, historyId)
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

func (m *defaultHistoryModel) FindOneByHomestayIdAndUserId(ctx context.Context, homestayId, userId int64) (*History, error) {
	//looklookTravelHistoryIdKey := fmt.Sprintf("%s%v%v", cacheLooklookTravelHistoryIdPrefix, homestayId, userId)
	var resp History
	query := fmt.Sprintf("select %s from %s where `homestay_id` = ? and `user_id` = ? limit 1", historyRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, homestayId, userId)
	//err := m.QueryRowCtx(ctx, &resp, looklookTravelHistoryIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
	//	query := fmt.Sprintf("select %s from %s where `homestay_id` = ? and `user_id` = ? limit 1", historyRows, m.table)
	//	return conn.QueryRowCtx(ctx, v, query, homestayId, userId)
	//})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultHistoryModel) FindOneByUserId(ctx context.Context, userId int64) (*History, error) {
	looklookTravelHistoryIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHistoryIdPrefix, userId)
	var resp History
	err := m.QueryRowCtx(ctx, &resp, looklookTravelHistoryIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", historyRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, userId)
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

func (m *defaultHistoryModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*History, error) {

	builder = builder.Columns(homestayRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*History
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHistoryModel) Insert(ctx context.Context, data *History) (sql.Result, error) {
	// bug: 没给DelState和DeleteTime做初始化
	data.DelState = globalkey.DelStateNo
	data.DeleteTime = time.Unix(0, 0)
	looklookTravelHistoryIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHistoryIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, historyRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Title, data.Cover, data.Intro, data.Location, data.PriceBefore, data.PriceAfter, data.RowState, data.HomestayBusinessId, data.RatingStars, data.UserId, data.HomestayId, data.DelState, data.Version, data.DeleteTime)
	}, looklookTravelHistoryIdKey)
	return ret, err
}

func (m *defaultHistoryModel) SelectBuilder()squirrel.SelectBuilder{
	return squirrel.Select().From(m.table)
}

func (m *defaultHistoryModel) Update(ctx context.Context, data *History) error {
	looklookTravelHistoryIdKey := fmt.Sprintf("%s%v", cacheLooklookTravelHistoryIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, historyRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Title, data.Cover, data.Intro, data.Location, data.PriceBefore, data.PriceAfter, data.RowState, data.HomestayBusinessId, data.RatingStars, data.UserId, data.Id)
	}, looklookTravelHistoryIdKey)
	return err
}

func (m *defaultHistoryModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheLooklookTravelHistoryIdPrefix, primary)
}

func (m *defaultHistoryModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", historyRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultHistoryModel) tableName() string {
	return m.table
}
