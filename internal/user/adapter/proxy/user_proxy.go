package proxy

import (
	"context"
	"github.com/core-go/sql"

	. "go-service/internal/user/domain"
)

const (
	tableUsers = "users"
)

func NewSqlLayerAdapter(proxy sql.Proxy) *SqlLayerAdapter {
	return &SqlLayerAdapter{
		proxy: proxy,
	}
}

type SqlLayerAdapter struct {
	proxy sql.Proxy
}

func (r *SqlLayerAdapter) Load(ctx context.Context, id string) (*User, error) {
	query := "select id, username, email, phone, date_of_birth from users where id = ?"
	var users []User

	err := r.proxy.Query(ctx, &users, query, id)
	if err != nil {
		return nil, err
	}
	return &users[0], nil
}

func (r *SqlLayerAdapter) Create(ctx context.Context, user *User) (int64, error) {
	txn := GetTxId(ctx)
	stmt, args := sql.BuildToInsert(tableUsers, user, sql.BuildParam, nil)

	if txn == nil {
		return r.proxy.Exec(ctx, stmt, args...)
	}
	return r.proxy.ExecTx(ctx, *txn, false, stmt, args...)
}

func (r *SqlLayerAdapter) Update(ctx context.Context, user *User) (int64, error) {
	txn := GetTxId(ctx)
	stmt, args := sql.BuildToUpdate(tableUsers, user, sql.BuildParam, nil)

	if txn == nil {
		return r.proxy.Exec(ctx, stmt, args...)
	}
	return r.proxy.ExecTx(ctx, *txn, false, stmt, args...)
}

func (r *SqlLayerAdapter) Patch(ctx context.Context, user map[string]interface{}) (int64, error) {
	txn := GetTxId(ctx)
	keys := []string{"id"}
	stmt, args := sql.BuildToPatch(tableUsers, user, keys, sql.BuildParam, nil)

	if txn == nil {
		return r.proxy.Exec(ctx, stmt, args...)
	}
	return r.proxy.ExecTx(ctx, *txn, false, stmt, args...)
}

func (r *SqlLayerAdapter) Delete(ctx context.Context, id string) (int64, error) {
	txn := GetTxId(ctx)
	ids := map[string]interface{}{"id": id}
	stmt, args := sql.BuildToDelete(tableUsers, ids, sql.BuildParam)

	if txn == nil {
		return r.proxy.Exec(ctx, stmt, args...)
	}
	return r.proxy.ExecTx(ctx, *txn, false, stmt, args...)
}

func GetTxId(ctx context.Context) *string {
	txi := ctx.Value("txId")
	if txi != nil {
		txx, ok := txi.(*string)
		if ok {
			return txx
		}
	}
	return nil
}
