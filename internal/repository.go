package checkin

import "sync"

// We need a user repository
type Repository struct {
	m     sync.Mutex
	store map[UserID]*User
	UserRepository
}

func NewUserRepository() UserRepository {
	return &Repository{
		store: make(map[UserID]*User),
	}
}

func (r *Repository) GetUser(uid UserID) (*User, error) {
	user := r.store[uid]
	return user, nil
}

func (r *Repository) Store(user *User) error {
	r.m.Lock()
	defer r.m.Unlock()
	r.store[user.UserID] = user
	return nil
}
