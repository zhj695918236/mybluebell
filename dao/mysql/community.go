package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"mybluebell/models"
)

func GetComunityList() (com []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	if err := db.QueryRow(sqlStr).Scan(&com); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}
