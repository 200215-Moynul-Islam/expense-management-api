package controllers

import (
	"errors"
	"net/http"

	appErrors "expense-management-api/errors"
	"expense-management-api/repositories"
	"expense-management-api/services"
	"expense-management-api/utils"
)

type UserController struct {
	BaseController
	userService services.UserService
}

func (c *UserController) Prepare() {
	c.userService = services.NewUserService(
		repositories.NewUserRepository(),
	)
}

func (c *UserController) GetProfile() {

	userID, ok := c.getUserID()
	if !ok {
		utils.SendJSONResponse(
			c.Ctx,
			http.StatusUnauthorized,
			false,
			"Unauthorized.",
			nil,
		)
		return
	}

	user, err := c.userService.GetByID(userID)
	if err != nil {

		switch {

		case errors.Is(err, appErrors.ErrUserNotFound):

			utils.SendJSONResponse(
				c.Ctx,
				http.StatusNotFound,
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
		http.StatusOK,
		true,
		"User profile retrieved successfully.",
		user,
	)
}
