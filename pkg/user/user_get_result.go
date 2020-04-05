package user

type UserGetResult struct {
	Id          string
	Name        string
	MailAddress string
}

func NewUserGetResult(user UserData) UserGetResult {
	return UserGetResult{user.Id, user.Name, user.MailAddress}
}
