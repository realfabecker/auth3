package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	cordom "github.com/realfabecker/auth3/internal/core/domain"
	corpts "github.com/realfabecker/auth3/internal/core/ports"
)

type AuthController struct {
	authSrv corpts.AuthService
}

func NewAuthController(
	authSrv corpts.AuthService,
) *AuthController {
	return &AuthController{authSrv}
}

// Login user login
//
//	@Summary		User login
//	@Description	User login
//	@Tags			Auth
//	@Param			request	body	cordom.UserLoginDTO	true	"Login payload"
//	@Produce		json
//	@Success		200	{object}	cordom.ResponseDTO[cordom.UserToken]
//	@Failure		400
//	@Failure		500
//	@Router			/auth3/login [post]
func (w *AuthController) Login(c *fiber.Ctx) error {
	q := cordom.UserLoginDTO{}
	if err := c.BodyParser(&q); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	v := validator.New(validator.WithRequiredStructEnabled())
	if err := v.Struct(q); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	token, err := w.authSrv.Login(q.Email, q.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if token.AuthChallenge != nil {
		return c.Status(401).JSON(cordom.ResponseDTO[cordom.UserToken]{
			Status: "error",
			Data:   token,
		})
	}

	return c.JSON(cordom.ResponseDTO[cordom.UserToken]{
		Status: "success",
		Data:   token,
	})
}

// Change Password
//
//	@Summary		Change password
//	@Description	Change password
//	@Tags			Auth
//	@Param			request	body	cordom.UserLoginChangeDTO	true	"Login payload"
//	@Produce		json
//	@Success		200	{object}	cordom.ResponseDTO[cordom.UserToken]
//	@Failure		400
//	@Failure		500
//	@Router			/auth3/change [post]
func (w *AuthController) Change(c *fiber.Ctx) error {
	q := cordom.UserLoginChangeDTO{}
	if err := c.BodyParser(&q); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	v := validator.New(validator.WithRequiredStructEnabled())
	if err := v.Struct(q); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	token, err := w.authSrv.Change(q.Email, q.NewPassword, q.Session)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(cordom.ResponseDTO[cordom.UserToken]{
		Status: "success",
		Data:   token,
	})
}

// Healthcheck api Healthcheck
//
//	@Summary		healthcheck
//	@Description	healthcheck
//	@Tags			Auth
//	@Produce		json
//	@Success		200	{object}	cordom.EmptyResponseDTO
//	@Router			/auth3/status [get]
func (w *AuthController) Healthcheck(c *fiber.Ctx) error {
	return c.JSON(cordom.EmptyResponseDTO{
		Status:  "success",
		Code:    200,
		Message: "Up and Running!",
	})
}
