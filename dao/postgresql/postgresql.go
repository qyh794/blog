package postgresql

import (
	"blog/settings"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init(cfg *settings.PostgresqlConfig) (err error) {
	if cfg == nil {
		return errors.New("postgresql config is nil")
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Dbname,
		cfg.SslMode,
	)
	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		zap.L().Error("postgresql connect failed, err:", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(5)
	return
}

func Close() {
	_ = db.Close()
}
