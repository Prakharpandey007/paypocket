package service

import (
	"context"
	"errors"
	"github.com/Prakharpandey007/paypocket/internal/model"
	"github.com/Prakharpandey007/paypocket/internal/repository"
	"github.com/Prakharpandey007/paypocket/internal/utils"
	"github.com/google/uuid"
	"time"
)

type UserService interface {
	SignupUser(ctx context.Context, req model.SignupRequest) (string, error)
	SignupUserReturnUser(ctx context.Context, req model.SignupRequest) (*model.User, error)
	Login(ctx context.Context, req model.LoginRequest) (string, error)
	ListUsers(ctx context.Context) ([]model.User, error)
}

type userService struct {
	repo repository.Repository[model.User]
}

func NewUserService(r repository.Repository[model.User]) UserService {
	return &userService{repo: r}
}
func (s *userService) SignupUser(ctx context.Context, req model.SignupRequest) (string, error) {
	existing, err := s.repo.Query(ctx, map[string]interface{}{"email": req.Email})
	if err != nil {
		return "", err
	}
	if len(existing) > 0 {
		return "", errors.New("email already registered")
	}
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return "", errors.New("failed to hash password")
	}

	// roleUUID, err := uuid.Parse(req.RoleId)
	// if err != nil {
	// 	return "", errors.New("invalid role ID")
	// }
	user := model.User{
		ID:           uuid.New(),
		Email:        req.Email,
		PasswordHash: hashedPassword,
		// RoleID:       roleUUID,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		PhoneNumber:  req.PhoneNumber,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	createdUser, err := s.repo.Create(ctx, &user)
	if err != nil {
		return "", err
	}
	// token, err := utils.GenerateJWT(createdUser.ID, createdUser.Email, createdUser.RoleID)
	token, err := utils.GenerateJWT(createdUser.ID, createdUser.Email,)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil

}

func (s *userService) SignupUserReturnUser(ctx context.Context, req model.SignupRequest) (*model.User, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// roleUUID, err := uuid.Parse(req.RoleId)
	// if err != nil {
	// 	return nil, errors.New("invalid role ID")
	// }

	user := model.User{
		ID:           uuid.New(),
		Email:        req.Email,
		PasswordHash: hashedPassword,
		// RoleID:       roleUUID,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		PhoneNumber:  req.PhoneNumber,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return s.repo.Create(ctx, &user)
}

func (s *userService) Login(ctx context.Context, req model.LoginRequest) (string, error) {
	users, err := s.repo.Query(ctx, map[string]interface{}{"email": req.Email})
	if err != nil || len(users) == 0 {
		return "", errors.New("invalid email or password")
	}
	user := users[0]
	if err := utils.CheckPassword(user.PasswordHash, req.Password); err != nil {
		return "", errors.New("invalid email or password")
	}
	user.UpdatedAt = time.Now()
	if err := s.repo.Update(ctx, &user); err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(user.ID, user.Email,)
	// token, err := utils.GenerateJWT(user.ID, user.Email, user.RoleID)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
func (s *userService) ListUsers(ctx context.Context) ([]model.User, error) {
	return s.repo.List(ctx)
}
