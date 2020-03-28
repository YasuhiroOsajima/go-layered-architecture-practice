package user

type User struct {
	id          UserId
	name        UserName
	mailAddress UserMailAddress
	userType    UserType
}

func NewUser(id UserId, name UserName, mailAddress UserMailAddress, userTpe UserType) *User {
	return &User{id, name, mailAddress, userTpe}
}

func NewUserInit(name UserName, mailAddress UserMailAddress, repo UserRepositoryInterface) (*User, error) {
	userId, err := NewUserIdRandom(repo)
	if err != nil {
		return nil, err
	}

	userType := NewUserNormal()

	return &User{userId, name, mailAddress, userType}, nil
}

func (u User) Equals(user *User) bool {
	return u.id == user.id
}

func (u User) Id() UserId {
	return u.id
}

func (u User) Name() UserName {
	return u.name
}

func (u User) MailAddress() UserMailAddress {
	return u.mailAddress
}

func (u User) Type() UserType {
	return u.userType
}

func (u User) IsPremium() bool {
	premiumUser := newPremiumUserType()
	return u.userType == premiumUser
}

func (u *User) ChangeName(name UserName) {
	u.name = name
}

func (u *User) ChangeMailAddress(mailAddress UserMailAddress) {
	u.mailAddress = mailAddress
}

func (u *User) Upgrade() {
	u.userType = newPremiumUserType()
}

func (u *User) Downgrade() {
	u.userType = newNormalUserType()
}
