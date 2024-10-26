package users

import (
	"net/http"

	"github.com/ADAGroupTcc/ms-users-api/exceptions"
	"github.com/ADAGroupTcc/ms-users-api/internal/domain"
	"github.com/ADAGroupTcc/ms-users-api/internal/helpers"
	"github.com/ADAGroupTcc/ms-users-api/internal/services/users"
	"github.com/labstack/echo/v4"
)

type Handler interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
	List(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type userHandler struct {
	userService users.Service
}

func New(userService users.Service) Handler {
	return &userHandler{
		userService,
	}
}

func (h *userHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()

	var userRequest domain.UserRequest
	if err := c.Bind(&userRequest); err != nil {
		return exceptions.New(exceptions.ErrInvalidPayload, err)
	}

	user, err := h.userService.Create(ctx, userRequest)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *userHandler) Get(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	user, err := h.userService.Get(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (h *userHandler) List(c echo.Context) error {
	ctx := c.Request().Context()

	var queryParams helpers.QueryParams
	err := helpers.BindQueryParams(c, &queryParams)
	if err != nil {
		return exceptions.New(exceptions.ErrInvalidPayload, err)
	}

	if queryParams.ShowCategories {
		users, err := h.userService.ListWithCategories(ctx, queryParams.UserIDs)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, users)
	}

	users, err := h.userService.List(ctx, queryParams.UserIDs, queryParams.Limit, queryParams.Offset)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func (h *userHandler) Update(c echo.Context) error {
	ctx := c.Request().Context()

	var userRequest domain.UserPatchRequest
	if err := c.Bind(&userRequest); err != nil {
		return exceptions.New(exceptions.ErrInvalidPayload, err)
	}
	id := c.Param("id")
	if id == "" {
		return exceptions.New(exceptions.ErrInvalidID, nil)
	}

	err := h.userService.Update(ctx, id, userRequest)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *userHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	err := h.userService.Delete(ctx, id)
	if err != nil {
		return c.NoContent(http.StatusNoContent)
	}

	return c.NoContent(http.StatusNoContent)
}
