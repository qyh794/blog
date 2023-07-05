package mysql

import (
	_ "github.com/go-sql-driver/mysql"
)

//var db *sqlx.DB

//func Init(cfg *settings.PostgresqlConfig) (err error) {
//	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
//		cfg.User,
//		cfg.Password,
//		cfg.Host,
//		cfg.Port,
//		cfg.Dbname,
//	)
//	db, err = sqlx.Connect("mysql", dsn)
//	if err != nil {
//		zap.L().Error("sql connect failed, err:", zap.Error(err))
//		return
//	}
//	return
//}
//
//func Close() {
//	_ = db.Close()
//}
