package store

import "github.com/gopherschool/http-rest-api/internal/app/model"

//UserRepository ...
type UserRepository struct {
	store *Store
}

//Create ...
func (ru *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := ru.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id", u.Email, u.EncryptedPassword).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

//FindByEmail ...
func (ru *UserRepository) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}
