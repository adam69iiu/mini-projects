package service

import (
	"context"
	"fmt"

	"github.com/adam69iiu/backend/internal/db"
	"github.com/adam69iiu/backend/internal/dto"
	"github.com/adam69iiu/backend/internal/util"
)

type UserRepo interface{
	CreateUser(ctx context.Context, id, username, email, password string) (db.User, error)
	GetUserByEmail(ctx context.Context, email string) (db.User, error)

}
type AuthService struct {
	userRepo UserRepo
	JWTSecret string
}

func NewAuthService(userRepo UserRepo, secretJwt string) *AuthService {
	return  &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error) {
	_, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err == nil {
		return dto.AuthResponse{}, fmt.Errorf("user with this email exist! : %w",err)
	}
	
	id := util.GenerateID()
	username := util.GenerateUsername()
	hashPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return dto.AuthResponse{}, fmt.Errorf("failed to hash the password : %w",err)
	}
  
	user, err := s.userRepo.CreateUser(ctx, id, username, req.Email,hashPassword)
	if err != nil {
		return dto.AuthResponse{}, fmt.Errorf("failed to create user: %w",err)
	}
  token, err := util.GenerateToken(id, s.JWTSecret)  
	return dto.AuthResponse{
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
		Token: token,
	}, nil
  

}



