package controllers

import (
	"github.com/aliftech/locksmith/internal/core/app/dto"
	"github.com/aliftech/locksmith/internal/core/app/services"
	"github.com/aliftech/locksmith/internal/core/helpers"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	Service *services.AuthService
}

func NewAuthController(service *services.AuthService) *AuthController {
	return &AuthController{Service: service}
}

func (ac *AuthController) HandleSignup(c *fiber.Ctx) error {
	input := new(dto.SignupreqBody)
	if err := c.BodyParser(input); err != nil {
		return helpers.Response(c, fiber.StatusInternalServerError, helpers.Message("parse_reqbody_failed", ""), nil)
	}

	user, err := ac.Service.Signup(*input)
	if err != nil {
		// You could map errors to HTTP status codes here
		if err.Error() == "email already registered" {
			return helpers.Response(c, fiber.StatusConflict, helpers.Message("email_exist", ""), nil)
		}
		return helpers.Response(c, fiber.StatusBadRequest, helpers.Message("signup_failed", err.Error()), nil)
	}

	return helpers.Response(c, fiber.StatusCreated, helpers.Message("signup_success", ""), user)
}
