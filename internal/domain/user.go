package domain

type User struct {
	ID           int64  `db:"id"`
	Email        string `db:"email"`
	PasswordHash string `db:"password_hash"`
	FullName     string `db:"full_name"`
}

type UserRepository interface {
	Create(user *User) error
}
