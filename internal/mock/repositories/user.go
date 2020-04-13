package mock

import checkin "ori/microservice/internal"

type UserRepository struct {
	GetUserFn func(uid checkin.UserID) (*checkin.User, error)
	StoreFn   func(u *checkin.User) error
}

func (ur *UserRepository) GetUser(uid checkin.UserID) (*checkin.User, error) {
	return ur.GetUserFn(uid)
}
func (ur *UserRepository) Store(u *checkin.User) error {
	return ur.StoreFn(u)
}
