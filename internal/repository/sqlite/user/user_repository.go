package user

import (
	"database/sql"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"go-layered-architecture-practice/internal/domain/models/user"
)

var dbFileName = "../test.db"

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
	found, err := r.Find(targetUser.Id())
	if err != nil {
		return err
	}

	if found == nil {
		_, err = r.db.NamedExec("INSERT INTO `user` (`id`, `name`, `usertype`) VALUES (:uid, :uname, :utype);",
			map[string]interface{}{
				"uid":   targetUser.Id(),
				"uname": targetUser.Name(),
				"utype": targetUser.Type(),
			})
		return err
	} else {
		_, err = r.db.NamedExec("UPDATE `user` SET `name`=:uname, `usertype`=:utype WHERE `id`=:uid;",
			map[string]interface{}{
				"uid":   targetUser.Id(),
				"uname": targetUser.Name(),
				"utype": targetUser.Type(),
			})
		return err
	}
}

func (r userRepository) Find(targetUserId user.UserId) (*user.User, error) {
	var id, name, mailaddress, usertype string
	err := r.db.QueryRow("SELECT * FROM `user` WHERE `id`=?;", targetUserId).Scan(&id, &name, &mailaddress, &usertype)
	if err != nil {
		if err != sql.ErrNoRows && !strings.HasPrefix(err.Error(), "could not find name ") {
			return nil, err
		}
		return nil, nil
	}

	uid, _ := user.NewUserId(id)
	uname, _ := user.NewUserName(name)
	utype, _ := user.NewUserType(usertype)
	umail, _ := user.NewUserMailAddress(mailaddress)
	targetUser := user.NewUser(uid, uname, umail, utype)
	return targetUser, nil
}

func (r userRepository) FindAll(targetUserName user.UserName) ([]*user.User, error) {
	rows, err := r.db.Query("SELECT * FROM `user` WHERE `name`=?;", targetUserName)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var users []*user.User
	for rows.Next() {
		var id, name, mailaddress, usertype string
		err = rows.Scan(id, name, mailaddress, usertype)
		if err != nil {
			return nil, err
		}
		uid, _ := user.NewUserId(id)
		uname, _ := user.NewUserName(name)
		utype, _ := user.NewUserType(usertype)
		umail, _ := user.NewUserMailAddress(mailaddress)
		users = append(users, user.NewUser(uid, uname, umail, utype))
	}

	return users, nil
}
func (r *userRepository) Delete(targetUser *user.User) error {
	_, err := r.db.NamedExec("DELETE FROM `user` WHERE `id`=:id;", targetUser)
	return err
}
