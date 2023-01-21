package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/roviery/catetin-api/constant"
	"github.com/roviery/catetin-api/domain"
	"github.com/roviery/catetin-api/entity"
	"github.com/roviery/catetin-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) Login(ctx context.Context, req entity.LoginRequest) (*entity.LoginResponse, error) {
	user, err := u.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if !checkPasswordHash(req.Password, user.Password) {
		return nil, fmt.Errorf("invalid password")
	}

	t, err := createToken(user.Fullname, user.ID)
	if err != nil {
		return nil, err
	}

	return &entity.LoginResponse{
		Token: t,
	}, nil
}

func (u *userUsecase) Register(ctx context.Context, req entity.RegisterRequest) (*entity.RegisterResponse, error) {
	user, _ := u.userRepo.FindByEmail(ctx, req.Email)
	if user != nil {
		return nil, fmt.Errorf("email already registered")
	}

	passHash, err := hashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	req.Password = passHash

	res, err := u.userRepo.Store(ctx, &models.User{
		Fullname: req.Fullname,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &entity.RegisterResponse{
		UserID: res.ID,
	}, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func createToken(userFullname string, userId primitive.ObjectID) (string, error) {
	claims := &entity.JWTCustomClaims{
		Name: userFullname,
		ID:   userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(constant.JWTExpiredTime)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(constant.JWTSecretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}
