package user

import "go-layered-architecture-practice/internal/domain/models/user"

type UserRepository struct {
	db []*user.User
}

func NewUserRepository() *UserRepository {
	var db []*user.User
	return &UserRepository{db}
}

func (r *UserRepository) Save(targetUser *user.User) error {
	targetIndex := -1
	for i, u := range r.db {
		if targetUser.Id() == u.Id() {
			targetIndex = i
			break
		}
	}

	if targetIndex >= 0 {
		r.db[targetIndex] = targetUser
		return nil
	} else {
		r.db = append(r.db, targetUser)
		return nil
	}
}

func (r UserRepository) Find(targetUserId user.UserId) (*user.User, error) {
	var targetUser *user.User
	for _, u := range r.db {
		if u.Id() == targetUserId {
			targetUser = user.NewUser(u.Id(), u.Name(), u.MailAddress(), u.Type())
		}
	}
	return targetUser, nil
}

func (r UserRepository) FindAll(targetUserName user.UserName) ([]*user.User, error) {
	var users []*user.User
	for _, u := range r.db {
		if u.Name() == targetUserName {
			copyuser := user.NewUser(u.Id(), u.Name(), u.MailAddress(), u.Type())
			users = append(users, copyuser)
		}
	}
	return users, nil
}

func (r *UserRepository) Delete(targetUser *user.User) error {
	for i, u := range r.db {
		if u.Id() == targetUser.Id() {
			r.db = append(r.db[:i], r.db[i+1:]...)
		}
	}
	return nil
}
