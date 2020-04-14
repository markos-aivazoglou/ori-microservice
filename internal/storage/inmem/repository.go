package inmem

import (
	"context"
	checkin "ori/microservice/internal"
	"sync"
)

// We need a user repository
type Repository struct {
	m     sync.Mutex
	store map[checkin.UserID]*checkin.User
}

func NewUserRepository() checkin.UserRepository {
	return &Repository{
		store: make(map[checkin.UserID]*checkin.User),
	}
}

func (r *Repository) GetUser(ctx context.Context, uid checkin.UserID) (*checkin.User, error) {
	user := r.store[uid]
	return user, nil
}

func (r *Repository) Store(ctx context.Context, user *checkin.User) error {
	r.m.Lock()
	defer r.m.Unlock()
	r.store[user.ID] = user
	return nil
}
