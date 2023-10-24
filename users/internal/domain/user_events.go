package domain

type UserRegistered struct {
	User *User
}

func (UserRegistered) EventName() string { return "users.UserRegistered" }

type UserAuthorized struct {
	User *User
}

func (UserAuthorized) EventName() string { return "users.UserAuthorized" }
