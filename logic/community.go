package logic

import (
	"mybluebell/dao/mysql"
	"mybluebell/models"
)

func GetCommunityList() ([]*models.Community) {
	return mysql.GetComunityList()
}
