package logic

import (
	"blog/dao/postgresql"
	"blog/models"
)

func CommunityList() ([]*models.Community, error) {
	return postgresql.CommunityList()
}

func GetCommunityDetailByID(id int64) (models.CommunityDetail, error) {
	return postgresql.GetCommunityDetailByID(id)
}
