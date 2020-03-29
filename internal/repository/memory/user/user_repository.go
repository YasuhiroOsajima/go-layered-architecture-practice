package user

import "go-layered-architecture-practice/internal/domain/models/user"

type userRepository struct {
	db []*user.User
}

func NewUserRepository() *userRepository {
	var db []*user.User
	return &userRepository{db}
}

func (r *userRepository) Save(targetUser *user.User) error {
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

func (r userRepository) Find(targetUserId user.UserId) (*user.User, error) {
	var targetUser *user.User
	for _, u := range r.db {
		if u.Id() == targetUserId {
			targetUser = user.NewUser(u.Id(), u.Name(), u.MailAddress(), u.Type())
		}
	}
	return targetUser, nil
}

func (r userRepository) FindAll(targetUserName user.UserName) ([]*user.User, error) {
	var users []*user.User
	for _, u := range r.db {
		if u.Name() == targetUserName {
			copyuser := user.NewUser(u.Id(), u.Name(), u.MailAddress(), u.Type())
			users = append(users, copyuser)
		}
	}
	return users, nil
}

func (r *userRepository) Delete(targetUser *user.User) error {
	for i, u := range r.db {
		if u.Id() == targetUser.Id() {
			r.db = append(r.db[:i], r.db[i+1:]...)
		}
	}
	return nil
}
