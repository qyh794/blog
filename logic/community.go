package logic

import (
	"blog/dao/mysql"
	"blog/models"
)

func CommunityList() ([]*models.Community, error) {
	return mysql.CommunityList()
}

func GetCommunityDetailByID(id int64) (models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
