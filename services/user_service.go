package services

import (
	"errors"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nahdukesaba/be-assignment/helpers"
	"github.com/nahdukesaba/be-assignment/repo"
	"github.com/nahdukesaba/be-assignment/services/request"
	"github.com/nahdukesaba/be-assignment/services/response"
)

type UserService struct {
	userRepo repo.UserRepository
}

type Users interface {
	RegisterUser(ctx *gin.Context, form *request.UserRequest) error
	LoginUser(ctx *gin.Context, form *request.UserRequest) (*response.UserLoginUser, error)
}

func NewUserService(db *repo.DB) Users {
	return &UserService{
		userRepo: repo.NewUserRepo(db),
	}
}

func (us *UserService) RegisterUser(ctx *gin.Context, form *request.UserRequest) error {
	if err := form.Validate(); err != nil {
		log.Println("Error Validate form")
		return err
	}

	if exUser, err := us.userRepo.GetUserByUsername(ctx, form.Username); err != nil && !strings.ContainsAny(err.Error(), "no row") {
		log.Printf("Error GetUserByUsername. err: %s\n", err)
		return err
	} else if exUser != nil {
		log.Println("Username already exist")
		return errors.New("username already exist")
	}

	pass, err := helpers.HashPassword(form.Password)
	if err != nil {
		log.Printf("Error HashPassword. err: %s\n", err)
		return err
	}

	user := &repo.User{
		Username: form.Username,
		Password: pass,
	}
	if err := us.userRepo.CreateUser(ctx, user); err != nil {
		log.Printf("Error CreateUser. err: %s\n", err)
		return err
	}

	return nil
}

func (us *UserService) LoginUser(ctx *gin.Context, form *request.UserRequest) (*response.UserLoginUser, error) {
	if err := form.Validate(); err != nil {
		log.Println("Error Validate form")
		return nil, err
	}

	user, err := us.userRepo.GetUserByUsername(ctx, form.Username)
	if err != nil {
		log.Printf("Error GetUserByUsername. err: %s\n", err.Error())
		return nil, err
	}

	if !helpers.CheckPasswordHash(form.Password, user.Password) {
		log.Println("Error Wrong Password")
		return nil, errors.New("wrong password")
	}

	token, err := helpers.CreateToken(form.Username)
	if err != nil {
		log.Printf("Error CreateToken. err: %s\n", err.Error())
		return nil, err
	}

	resp := &response.UserLoginUser{
		Username: form.Username,
		Token:    token,
	}

	return resp, nil
}
