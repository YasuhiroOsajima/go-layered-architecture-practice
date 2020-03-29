package user

import (
	"database/sql"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"go-layered-architecture-practice/internal/domain/models/user"
)

var dbFileName = os.Getenv("SQLITE_PATH")

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository() *UserRepository {
	db, err := sqlx.Connect("sqlite3", dbFileName)
	if err != nil {
		panic(err.Error)
	}

	return &UserRepository{db}
}

func (r UserRepository) Save(targetUser *user.User) (err error) {
	found, err := r.Find(targetUser.Id())
	if err != nil {
		return err
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	if found == nil {
		_, err = tx.NamedExec("INSERT INTO `user` (`id`, `name`, `mailaddress`, `usertype`) VALUES (:uid, :uname, :umail, :utype);",
			map[string]interface{}{
				"uid":   targetUser.Id(),
				"uname": targetUser.Name(),
				"umail": targetUser.MailAddress(),
				"utype": targetUser.Type(),
			})
	} else {
		_, err = tx.NamedExec("UPDATE `user` SET `name`=:uname, `mailaddress`=:umail, `usertype`=:utype WHERE `id`=:uid;",
			map[string]interface{}{
				"uid":   targetUser.Id(),
				"uname": targetUser.Name(),
				"umail": targetUser.MailAddress(),
				"utype": targetUser.Type(),
			})
	}

	var txErr error
	if err != nil {
		txErr = tx.Rollback()
	} else {
		txErr = tx.Commit()
	}

	if err == nil && txErr != nil {
		return txErr
	}

	return err
}

func (r UserRepository) Find(targetUserId user.UserId) (*user.User, error) {
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

func (r UserRepository) FindAll(targetUserName user.UserName) ([]*user.User, error) {
	rows, err := r.db.Query("SELECT * FROM `user` WHERE `name`=?;", targetUserName)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var users []*user.User
	for rows.Next() {
		var id, name, mailaddress, usertype string
		err = rows.Scan(&id, &name, &mailaddress, &usertype)
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
func (r *UserRepository) Delete(targetUser *user.User) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.NamedExec("DELETE FROM `user` WHERE `id`=:id;", targetUser)

	var txErr error
	if err != nil {
		txErr = tx.Rollback()
	} else {
		txErr = tx.Commit()
	}

	if err == nil && txErr != nil {
		return txErr
	}

	return err
}
