package services

import (
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"go-layered-architecture-practice/internal/domain/models/user"
)

var dbFileName = "test.db"

func Exists(targetUser *user.User) (bool, error) {
	db, err := sqlx.Connect("sqlite3", dbFileName)
	if err != nil {
		return false, err
	}

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM `user` WHERE `id`=?;", targetUser.Id()).Scan(&count)
	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, nil
	} else if count == 1 {
		return true, nil
	} else {
		return false, errors.New("target user is invalid status")
	}
}
