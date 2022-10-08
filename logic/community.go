package logic

import (
	"mybluebell/dao/mysql"
	"mybluebell/models"
)

func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetComunityList()
}
