package mock

import (
	"context"
	checkin "ori/microservice/internal"
)

type UserRepository struct {
	GetUserFn func(ctx context.Context, uid checkin.UserID) (*checkin.User, error)
	StoreFn   func(ctx context.Context, u *checkin.User) error
}

func (ur *UserRepository) GetUser(ctx context.Context, uid checkin.UserID) (*checkin.User, error) {
	return ur.GetUserFn(ctx, uid)
}
func (ur *UserRepository) Store(ctx context.Context, u *checkin.User) error {
	return ur.StoreFn(ctx, u)
}
