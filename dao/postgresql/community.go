package postgresql

import (
	"blog/models"
	"database/sql"

	"go.uber.org/zap"
)

func CommunityList() (communityList []*models.Community, err error) {

	sqlStr := `select 
    community_id, community_name 
from 
   "community"`
	// 查询多行用select
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("no date")
			err = nil
		}
	}
	return
}

func GetCommunityDetailByID(id int64) (data models.CommunityDetail, err error) {

	sqlStr := `select 
    community_id, community_name, introduction, create_time 
from 
    "community" 
where 
    community_id = $1`
	if err = db.Get(&data, sqlStr, id); err != nil {
		zap.L().Error("db.select failed err:", zap.Error(err))
		if err == sql.ErrNoRows {
			zap.L().Warn("no data")
			err = ErrorInvalidCommunityID
		}
	}
	return data, err
}
