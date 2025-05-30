package di

import "go/adv-demo/internal/user"

type IStatRepository interface {
	AddClick(LinkId uint)
}

type IUserRepository interface {
	Create(user *user.User) (*user.User, error)
	FindByEmail(email string) (*user.User, error)
}
