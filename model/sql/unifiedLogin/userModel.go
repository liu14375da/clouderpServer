package unifiedLogin

import (
	"ZeroProject/model/entity/tbsfUser"
	"ZeroProject/model/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
	"strings"
)

var (
	userFieldNames        = builderx.RawFieldNames(&tbsfUser.TbsfUser{})
	userRows              = strings.Join(userFieldNames, ",")
	cacheUserNumberPrefix = "cache#User#number#"
	cacheUserIdPrefix     = "cache#User#id#"
)

type (
	UserModel interface {
		FindUnified(username string) (*tbsfUser.TbsfUser, error)
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}
)

/*
	初始化数据库表
*/
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`tbsf_user`",
	}
}

//根据登录名称查询数据库对应的信息
func (m *defaultUserModel) FindUnified(username string) (*tbsfUser.TbsfUser, error) {
	var resp tbsfUser.TbsfUser
	//userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, username+"44")
	err := m.QueryRow(&resp, username, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where user_name = ? ", userRows, m.table)
		str := strings.Replace(query, "`", "", -1)
		return conn.QueryRow(v, str, username)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, sql.ErrNotFound
	default:
		return nil, err
	}
}
