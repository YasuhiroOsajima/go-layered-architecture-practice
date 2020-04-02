package circle

import (
	"database/sql"
	"errors"
	mapset "github.com/deckarep/golang-set"
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
		_, err = tx.NamedExec("INSERT INTO `circle` (`id`, `name`, `owner`) VALUES (:cid, :cname, :owner);",
			map[string]interface{}{
				"cid":   targetCircle.Id(),
				"cname": targetCircle.Name(),
				"owner": targetCircle.OwnerId(),
			})
		if err != nil {
			tx.Rollback()
			return err
		}

	} else {
		_, err = tx.NamedExec("UPDATE `circle` SET `name`=:uname, `owner`=:owner WHERE `id`=:cid;",
			map[string]interface{}{
				"cid":   targetCircle.Id(),
				"cname": targetCircle.Name(),
				"owner": targetCircle.OwnerId(),
			})
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	rows, err := tx.Query("SELECT `user_id` FROM `user_circle` WHERE `circle_id`=?;", targetCircle.Id())
	if err != nil {
		tx.Rollback()
		return err
	}

	if (err != nil && err != sql.ErrNoRows) || (rows == nil) {
		for _, m := range targetCircle.Members() {
			_, err = tx.NamedExec("INSERT INTO `circle_id` (`user_id`, `circle_id`) VALUES (:uid, :uname, :cid);",
				map[string]interface{}{
					"uid": m.Id(),
					"cid": targetCircle.Id(),
				})
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	} else {
		var targetMembers []user.UserId
		for _, m := range targetCircle.Members() {
			targetMembers = append(targetMembers, m.Id())
		}

		var dbMembers []user.UserId
		for rows.Next() {
			var userId string
			err = rows.Scan(&userId)
			if err != nil {
				break
			}

			uid, err := user.NewUserId(userId)
			if err != nil {
				break
			}
			dbMembers = append(dbMembers, uid)
		}

		targetMembersSet := mapset.NewSet()
		for _, m := range targetMembers {
			targetMembersSet.Add(m)
		}
		dbMembersSet := mapset.NewSet()
		for _, m := range dbMembers {
			dbMembersSet.Add(m)
		}
		onlyTargetMembers := targetMembersSet.Difference(dbMembersSet).ToSlice()
		onlyDbMembers := dbMembersSet.Difference(targetMembersSet).ToSlice()

		if len(onlyTargetMembers) > 0 {
			for _, m := range onlyTargetMembers {
				_, err := tx.NamedExec("INSERT INTO `user_circle` (`user_id`, `circle_id`) VALUES (:user_id, :circle_id);",
					map[string]interface{}{
						"user_id":   m,
						"circle_id": targetCircle.Id(),
					})
				if err != nil {
					tx.Rollback()
					return err
				}
			}
		}
		if len(onlyDbMembers) > 0 {
			for _, m := range onlyDbMembers {
				_, err = tx.NamedExec("DELETE FROM `user` WHERE `id`=:id;", m)
				if err != nil {
					tx.Rollback()
					return err
				}
			}
		}
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
