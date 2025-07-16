//package postrgesql
//
//import (
//	"context"
//	"fmt"
//	"govno/internal/db"
//	"govno/internal/models"
//
//	"github.com/jackc/pgx/v5"
//	"github.com/jackc/pgx/v5/pgconn"
//	"github.com/jackc/pgx/v5/pgxpool"
//)
//
//var _ db.DB = (*repo)(nil)
//
//type repo struct {
//	conn connection
//}
//
//// dsn - строка подключения к БД
//func NewRepository(ctx context.Context, dsn string) (*repo, error) {
//	pgCfg, err := pgxpool.ParseConfig(dsn)
//	if err != nil {
//		return nil, fmt.Errorf("parse config: %w", err)
//	}
//
//	pool, err := pgxpool.NewWithConfig(ctx, pgCfg)
//	if err != nil {
//		return nil, fmt.Errorf("create pool: %w", err)
//	}
//
//	c, err := pool.Acquire(ctx)
//	if err != nil {
//		return nil, fmt.Errorf("pool acquire: %w", err)
//	}
//	c.Release()
//
//	return &repo{
//		conn: pool,
//	}, nil
//}
//
//
//type connection interface {
//	Begin(ctx context.Context) (pgx.Tx, error)
//	Close()
//	Ping(ctx context.Context) error
//
//	Exec(ctx context.Context, sql string, args ...any) (cTag pgconn.CommandTag, err error)
//	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
//	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
//}
//
//func (r *repo) GetAdminByLogpass(ctx context.Context, filter models.AdminModelLogpassFilter) ([]models.AdminModel, error) {
//	q := fmt.Sprintf("SELECT * FROM admins WHERE login=%s AND password=%s", filter.Login, filter.Password)
//
//	args := pgx.NamedArgs{}
//	rows, err := r.conn.Query(ctx, q, args)
//	if err != nil {
//		return nil, fmt.Errorf("get note from ...: %w", err)
//	}
//	defer rows.Close()
//
//	res := make([]models.AdminModel, 0)
//
//	for rows.Next() {
//		var admin models.AdminModel
//		if err := rows.Scan(&admin); err != nil {
//			return nil, fmt.Errorf("scan note ... : %w", err)
//		}
//		res = append(res, admin)
//	}
//
//	return res, nil
//}
//
//func (r *repo) Close() {
//	r.conn.Close()
//}