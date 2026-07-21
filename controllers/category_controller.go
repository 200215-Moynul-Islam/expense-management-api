package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"expense-management-api/dto"
	appErrors "expense-management-api/errors"
	"expense-management-api/repositories"
	"expense-management-api/services"
	"expense-management-api/utils"
)

type CategoryController struct {
	BaseController
	categoryService services.CategoryService
}

func (c *CategoryController) Prepare() {
	c.categoryService = services.NewCategoryService(
		repositories.NewCategoryRepository(),
	)
}

func (c *CategoryController) Create() {

	var request dto.CreateCategoryRequest

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

	request.Name = strings.TrimSpace(request.Name)

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

	err = c.categoryService.CreateCategory(
		request,
		userID,
	)

	if err != nil {
		c.handleCategoryError(err)
		return
	}

	utils.SendJSONResponse(
		c.Ctx,
		http.StatusCreated,
		true,
		"Category created successfully.",
		nil,
	)
}

func (c *CategoryController) Update() {

	var request dto.UpdateCategoryRequest

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

	id, err := c.GetInt(":id")
	if err != nil {
		utils.SendJSONResponse(
			c.Ctx,
			http.StatusBadRequest,
			false,
			"Invalid category ID.",
			nil,
		)
		return
	}

	err = c.categoryService.UpdateCategory(
		id,
		userID,
		request,
	)
	if err != nil {
		c.handleCategoryError(err)
		return
	}

	utils.SendJSONResponse(
		c.Ctx,
		http.StatusOK,
		true,
		"Category updated successfully.",
		nil,
	)
}

func (c *CategoryController) Delete() {

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

	id, err := c.GetInt(":id")
	if err != nil {
		utils.SendJSONResponse(
			c.Ctx,
			http.StatusBadRequest,
			false,
			"Invalid category ID.",
			nil,
		)
		return
	}

	err = c.categoryService.DeleteCategory(
		id,
		userID,
	)
	if err != nil {
		c.handleCategoryError(err)
		return
	}

	utils.SendJSONResponse(
		c.Ctx,
		http.StatusOK,
		true,
		"Category deleted successfully.",
		nil,
	)
}

func (c *CategoryController) handleCategoryError(
	err error,
) {

	switch {

	case errors.Is(err, appErrors.ErrCategoryNameRequired),
		errors.Is(err, appErrors.ErrCategoryNameTooLong),
		errors.Is(err, appErrors.ErrInvalidRequest):

		utils.SendJSONResponse(
			c.Ctx,
			http.StatusBadRequest,
			false,
			err.Error(),
			nil,
		)

	case errors.Is(err, appErrors.ErrCategoryExists):

		utils.SendJSONResponse(
			c.Ctx,
			http.StatusConflict,
			false,
			err.Error(),
			nil,
		)

	case errors.Is(err, appErrors.ErrCategoryNotFound):

		utils.SendJSONResponse(
			c.Ctx,
			http.StatusNotFound,
			false,
			err.Error(),
			nil,
		)

	case errors.Is(err, appErrors.ErrForbiddenCategory):

		utils.SendJSONResponse(
			c.Ctx,
			http.StatusForbidden,
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
}

func (c *CategoryController) GetByUserID() {

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

	categories, err := c.categoryService.GetCategoriesByUserID(
		userID,
	)
	if err != nil {
		utils.SendJSONResponse(
			c.Ctx,
			http.StatusInternalServerError,
			false,
			"Internal server error.",
			nil,
		)
		return
	}

	utils.SendJSONResponse(
		c.Ctx,
		http.StatusOK,
		true,
		"Categories retrieved successfully.",
		categories,
	)
}

func (c *CategoryController) GetByID() {

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

	id, err := c.GetInt(":id")
	if err != nil {
		utils.SendJSONResponse(
			c.Ctx,
			http.StatusBadRequest,
			false,
			"Invalid category ID.",
			nil,
		)
		return
	}

	category, err := c.categoryService.GetCategoryByID(
		id,
		userID,
	)
	if err != nil {

		switch {

		case errors.Is(err, appErrors.ErrCategoryNotFound):

			utils.SendJSONResponse(
				c.Ctx,
				http.StatusNotFound,
				false,
				err.Error(),
				nil,
			)

		case errors.Is(err, appErrors.ErrForbiddenCategory):

			utils.SendJSONResponse(
				c.Ctx,
				http.StatusForbidden,
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
		"Category retrieved successfully.",
		category,
	)
}
