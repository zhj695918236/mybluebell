package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"go.uber.org/zap"
	"mybluebell/models"
)

const secret = "hello world"

func InsertUser(user *models.User) error {
	// 对密码加密
	user.Password = encryptPassword(user.Password)
	sqlstr := "insert into  user(user_id,username,password) values(?,?,?)"
	_, err := db.Exec(sqlstr, user.UserID, user.Username, user.Password)
	return err
}

func encryptPassword(passowrd string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(passowrd)))
}

func CheckUserExist(username string) (err error) {
	sqlstr := "select count(user_id) from user where username=? "
	var count int64
	err = db.QueryRow(sqlstr, username).Scan(&count)
	if err != nil {
		zap.L().Error("CheckUserExist:mysql select error", zap.Error(err))
	}
	if count > 0 {
		return ErrorUserExist
	}
	return err
}

func Login(user *models.User) error {
	sqlstr := "select user_id,password from user where username=? "
	var (
		pwd    string
		userid int64
	)
	err := db.QueryRow(sqlstr, user.Username).Scan(&userid, &pwd)
	user.UserID = userid
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	if encryptPassword(user.Password) != pwd {
		return ErrorInvalidPassword
	}
	return err
}
