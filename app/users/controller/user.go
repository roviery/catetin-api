package controller

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/roviery/catetin-api/constant"
	"github.com/roviery/catetin-api/domain"
	"github.com/roviery/catetin-api/entity"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(userUsecase domain.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (u *UserHandler) Login(c echo.Context) error {
	var req entity.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	res, err := u.userUsecase.Login(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = res.Token
	cookie.Expires = time.Now().Add(time.Duration(constant.JWTExpiredTime) * time.Second)
	cookie.Path = "/"
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, entity.BaseResponse{
		Message: "success",
		Data:    res,
	})
}

func (u *UserHandler) Logout(c echo.Context) error {
	cookie, err := c.Cookie("token")
	if err != nil {
		return err
	}
	cookie.Value = ""
	cookie.Expires = time.Now()
	return nil
}

func (u *UserHandler) Register(c echo.Context) error {
	var req entity.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	res, err := u.userUsecase.Register(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.BaseResponse{
		Message: "success",
		Data:    res,
	})
}
