package services

import (
	"context"
	"errors"
	"time"
	"yas/cfg"
	"yas/database/models"
	"yas/repositories"
	"yas/types"

	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, userRegistration *types.UserRegistration) (*models.User, error)
	Login(ctx context.Context, login *types.UserLogin) (string, error)
}

type defaultUserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &defaultUserService{
		userRepository: repo,
	}
}

func (d *defaultUserService) Register(ctx context.Context, userRegistration *types.UserRegistration) (*models.User, error) {
	if userRegistration.Password != userRegistration.ConfirmPassword {
		return nil, errors.New("passwords don't match")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(userRegistration.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("error hashing user password")
		return nil, err
	}
	user := models.NewUser()
	user.Email = userRegistration.Email
	user.Name = userRegistration.Name
	user.Password = string(password)

	u, err := d.userRepository.CreateUser(ctx, user)
	if err != nil {
		log.Error().Err(err).Msg("error creating user")
		return nil, err
	}
	return u, nil
}

func (d *defaultUserService) Login(ctx context.Context, login *types.UserLogin) (string, error) {
	u, err := d.userRepository.GetUserByEmail(ctx, login.Email)
	if err != nil {
		log.Error().Err(err).Msgf("error getting user by email %s", login.Email)
		return "", err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(login.Password)); err != nil {
		return "", err
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    u.Id.String(),
		ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
	})
	token, err := claims.SignedString(cfg.SecretKey())
	if err != nil {
		return "", err
	}
	return token, nil
}
