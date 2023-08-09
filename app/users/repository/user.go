package repository

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/roviery/catetin-api/domain"
	"github.com/roviery/catetin-api/models"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) domain.UserRepository {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	fmt.Print(user)
	return &user, nil
}

func (u *userRepo) Store(ctx context.Context, user *models.User) (*models.User, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	err := u.db.Create(user).Error
	if err != nil {
		return nil, err
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
