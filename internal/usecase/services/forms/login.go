package forms

import (
	"github.com/gofiber/fiber/v2"

	"github.com/fiber-go-sis-app/utils/pkg/custom"
	"github.com/fiber-go-sis-app/utils/pkg/jwt"

	constantsEntity "github.com/fiber-go-sis-app/internal/entity/constants"
	formsEntity "github.com/fiber-go-sis-app/internal/entity/forms"

	userRepo "github.com/fiber-go-sis-app/internal/repo/users"
)

func LoginForm(ctx *fiber.Ctx, req formsEntity.LoginRequest) (formsEntity.LoginResponse, error) {
	// Initialization variable
	var res formsEntity.LoginResponse

	data, err := userRepo.GetUserByUserName(ctx, req.UserName)
	if err != nil {
		return res, err
	}

	// check hash password
	if !custom.CheckPasswordHash(req.Password, data.Password) {
		return res, constantsEntity.ErrWrongPassword
	}

	token, err := jwt.CreateJWTToken(formsEntity.JWTRequest{
		UserID: data.UserID,
		Name:   data.UserName,
		Admin:  data.IsAdmin,
	})
	if err != nil {
		return res, err
	}

	return formsEntity.LoginResponse{
		UserID:   data.UserID,
		UserName: data.UserName,
		IsAdmin:  data.IsAdmin,
		JWTToken: token,
	}, nil
}
