package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"expense-management-api/dto"
	appErrors "expense-management-api/errors"
	"expense-management-api/repositories"
	"expense-management-api/services"
	"expense-management-api/utils"

	beego "github.com/beego/beego/v2/server/web"
)

type AuthController struct {
	beego.Controller
	userService services.UserService
}

func (c *AuthController) Prepare() {
	c.userService = services.NewUserService(
		repositories.NewUserRepository(),
	)
}

func (c *AuthController) Register() {

	var request dto.RegisterRequest

	err := json.Unmarshal(
		c.Ctx.Input.RequestBody,
		&request,
	)
	if err != nil {
		utils.SendJSONResponse(
			c.Ctx,
			http.StatusBadRequest,
			false,
			"Invalid request body.",
			nil,
		)
		return
	}

	err = c.userService.RegisterUser(request)
	if err != nil {

		switch {

		case errors.Is(err, appErrors.ErrInvalidRequest),
			errors.Is(err, appErrors.ErrNameRequired),
			errors.Is(err, appErrors.ErrNameTooLong),
			errors.Is(err, appErrors.ErrEmailRequired),
			errors.Is(err, appErrors.ErrInvalidEmail),
			errors.Is(err, appErrors.ErrPasswordRequired),
			errors.Is(err, appErrors.ErrPasswordTooShort):

			utils.SendJSONResponse(
				c.Ctx,
				http.StatusBadRequest,
				false,
				err.Error(),
				nil,
			)

		case errors.Is(err, appErrors.ErrEmailExists):

			utils.SendJSONResponse(
				c.Ctx,
				http.StatusConflict,
				false,
				err.Error(),
				nil,
			)

		default:

			utils.SendJSONResponse(
				c.Ctx,
				http.StatusInternalServerError,
				false,
				"Internal server error.",
				nil,
			)
		}

		return
	}

	utils.SendJSONResponse(
		c.Ctx,
		http.StatusCreated,
		true,
		"User registered successfully.",
		nil,
	)
}
