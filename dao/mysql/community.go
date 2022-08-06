package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"mybluebell/models"
)

func GetComunityList() (com []*models.Community) {
	sqlStr := "select community_id, community_name from community"
	if err := db.Select(&com, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			return
		}
	}
	return
}
