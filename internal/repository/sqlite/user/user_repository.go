package user

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"go-layered-architecture-practice/internal/domain/models/user"
)

var dbFileName = "test.db"

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository() (*userRepository, error) {
	db, err := sqlx.Connect("sqlite3", dbFileName)
	if err != nil {
		return nil, err
	}

	return &userRepository{db}, nil
}

func (r userRepository) Save(targetUser *user.User) error {
	var id, name, usertype string
	err := r.db.QueryRow("SELECT * FROM `user` WHERE `id`=?;", targetUser.Id()).Scan(&id, &name, &usertype)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == sql.ErrNoRows {
		_, err = r.db.NamedExec("INSERT INTO `user` (id, name, usertype) VALUES (:id, :name, :usertype);", targetUser)
		return err
	} else {
		uid, err := user.NewUserId(id)
		if err != nil {
			return err
		}
		uname, err := user.NewUserName(name)
		if err != nil {
			return err
		}
		utype, err := user.NewUserType(usertype)
		if err != nil {
			return err
		}

		u := user.NewUser(uid, uname, utype)
		_, err = r.db.NamedExec("UPDATE `user` SET `name`=:name, `usertype`=:usertype WHERE `id`=:id;", u)
		return err
	}
}

func (r userRepository) Find(targetUserId user.UserId) (*user.User, error) {
	var id, name, usertype string
	err := r.db.QueryRow("SELECT * FROM `user` WHERE `id`=?;", targetUserId).Scan(id, name, usertype)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	uid, _ := user.NewUserId(id)
	uname, _ := user.NewUserName(name)
	utype, _ := user.NewUserType(usertype)
	targetUser := user.NewUser(uid, uname, utype)
	return targetUser, nil
}

func (r userRepository) FindAll(targetUserName user.UserName) ([]*user.User, error) {
	rows, err := r.db.Query("SELECT * FROM `user` WHERE `name`=?;", targetUserName)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var users []*user.User
	for rows.Next() {
		var id, name, usertype string
		err = rows.Scan(id, name, usertype)
		if err != nil {
			return nil, err
		}
		uid, _ := user.NewUserId(id)
		uname, _ := user.NewUserName(name)
		utype, _ := user.NewUserType(usertype)
		users = append(users, user.NewUser(uid, uname, utype))
	}

	return users, nil
}
