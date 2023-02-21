package repository

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/roviery/catetin-api/domain"
	"github.com/roviery/catetin-api/models"
)

type userRepo struct {
	sql *sql.DB
}

func NewUserRepo(sql *sql.DB) domain.UserRepository {
	return &userRepo{
		sql: sql,
	}
}

func (u *userRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	query := "SELECT id, email FROM user WHERE email = ?"
	row, err := u.sql.Query(query, email)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		err := row.Scan(&user.ID, &user.Email)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}

	fmt.Print(user)
	return nil, nil
}

func (u *userRepo) Store(ctx context.Context, user *models.User) (*models.User, error) {
	query := "INSERT INTO user(fullname, email, password, created_at) VALUES (?,?,?,?)"
	user.ID = randomString(10)
	user.CreatedAt = time.Now().Format(time.RFC1123)
	_, err := u.sql.Query(query, user.Fullname, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		panic(err)
	}

	return user, nil
}

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	rand.Seed(time.Now().UnixNano())

	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
