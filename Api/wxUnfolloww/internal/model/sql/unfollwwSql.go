package sql

import (
	"database/sql"
	"fmt"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

var (
	cacheUserNumberPrefix = "cache#User#number#"
	cacheUserIdPrefix     = "cache#User#id#"

)

type (
	UnFollowSql interface{
		UpdataIsValid(openid string) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}
)


func NewUnFollowSql(s sqlx.SqlConn, c cache.CacheConf) UnFollowSql {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(s, c),
		table: "TBSF_Conmmunicate_WXUser",
	}
}

func (d defaultUserModel) UpdataIsValid(openid string) error {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix,openid )
	_, err := d.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set IsValid =? where OpenId = ?", d.table)
		return conn.Exec(query,true, openid)
	}, userIdKey)
	return err
}