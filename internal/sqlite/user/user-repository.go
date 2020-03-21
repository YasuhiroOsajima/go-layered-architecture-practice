package user

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"go-layered-architecture-practice/internal/domain/models/user"
)

var dbFileName = "test.db"

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository() userRepository {
	db, err := sqlx.Connect("sqlite3", dbFileName)
	if err != nil {
		return false, err
	}

	return userRepository{db}
}

func (r userRepository) Save(targetUser *user.User) error {
	return nil
}

func (r userRepository) Find(targetUserName *user.UserName) (*user.User, error) {
	id, _ := user.NewUserId("1")
	name, _ := user.NewUserName("aaa")
	usertype, _ := user.NewUserType(user.Normal)
	targetUser := user.NewUser(id, name, usertype)
	return targetUser, nil
}
