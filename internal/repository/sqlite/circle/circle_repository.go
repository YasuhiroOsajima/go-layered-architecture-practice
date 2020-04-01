package circle

import (
	"database/sql"
	"errors"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"

	"go-layered-architecture-practice/internal/domain/models/circle"
	"go-layered-architecture-practice/internal/domain/models/user"
)

var dbFileName = os.Getenv("SQLITE_PATH")

type CircleRepository struct {
	db             *sqlx.DB
	userRepository user.UserRepositoryInterface
}

func NewCircleRepository(userRepository user.UserRepositoryInterface) *CircleRepository {
	db, err := sqlx.Connect("sqlite3", dbFileName)
	if err != nil {
		panic(err.Error)
	}

	return &CircleRepository{db, userRepository}
}

func (r CircleRepository) Save(targetCircle *circle.Circle) (err error) {
	found, err := r.Find(targetCircle.Id())
	if err != nil {
		return err
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	if found == nil {
		_, err = tx.NamedExec("INSERT INTO `circle` (`id`, `name`, `owner`, `members`) VALUES (:cid, :cname, :owner, :members);",
			map[string]interface{}{
				"cid":     targetCircle.Id(),
				"cname":   targetCircle.Name(),
				"owner":   targetCircle.OwnerId(),
				"members": targetCircle.MemberIds(),
			})
	} else {
		_, err = tx.NamedExec("UPDATE `user` SET `name`=:uname, `owner`=:owner, `members`=:members WHERE `id`=:uid;",
			map[string]interface{}{
				"cid":     targetCircle.Id(),
				"cname":   targetCircle.Name(),
				"owner":   targetCircle.OwnerId(),
				"members": targetCircle.MemberIds(),
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

func (r CircleRepository) Find(targetCircleId circle.CircleId) (*circle.Circle, error) {
	var id, name, ownerId string
	var members []string

	err := r.db.QueryRow("SELECT * FROM `circle` WHERE `id`=?;", targetCircleId).Scan(&id, &name, &ownerId, &members)
	if err != nil {
		if err != sql.ErrNoRows && !strings.HasPrefix(err.Error(), "could not find name ") {
			return nil, err
		}
		return nil, nil
	}

	cid, _ := circle.NewCircleId(id)
	cname, _ := circle.NewCircleName(name)

	targetOwnerId, err := user.NewUserId(ownerId)
	if err != nil {
		return nil, err
	}

	targetOwner, err := r.userRepository.Find(targetOwnerId)
	if err != nil {
		return nil, err
	}
	if targetOwner == nil {
		return nil, errors.New("invalid error, target circle's owner is already deleted")
	}

	var targetMembers []*user.User
	for _, m := range members {
		uid, err := user.NewUserId(m)
		if err != nil {
			return nil, err
		}

		u, err := r.userRepository.Find(uid)
		if err != nil {
			return nil, err
		}

		targetMembers = append(targetMembers, u)
	}

	targetCircle := circle.NewCircle(cid, cname, targetOwner, targetMembers)
	return targetCircle, nil
}
