package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/roviery/catetin-api/domain"
	"github.com/roviery/catetin-api/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeadlineHandler struct {
	deadlineUsecase domain.DeadlineUsecase
}

func NewDeadlineHandler(deadlineUsecase domain.DeadlineUsecase) *DeadlineHandler {
	return &DeadlineHandler{
		deadlineUsecase: deadlineUsecase,
	}
}

func (d *DeadlineHandler) Create(c echo.Context) error {
	var req entity.DeadlineRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	if len(req.DeadlineDate) == 0 || len(req.Priority) == 0 || len(req.Task) == 0 {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: "empty request data exist",
		})
	}

	ownerId, err := primitive.ObjectIDFromHex(c.Get("user_id").(string))
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	res, err := d.deadlineUsecase.CreateDeadline(c.Request().Context(), &ownerId, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, entity.BaseResponse{
		Message: "success",
		Data:    res,
	})
}

func (d *DeadlineHandler) GetDeadline(c echo.Context) error {
	deadlineId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	res, err := d.deadlineUsecase.GetDeadline(c.Request().Context(), &deadlineId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.BaseResponse{
		Message: "success",
		Data:    res,
	})
}

func (d *DeadlineHandler) GetDeadlines(c echo.Context) error {
	ownerId, err := primitive.ObjectIDFromHex(c.Get("user_id").(string))
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	res, err := d.deadlineUsecase.GetDeadlines(c.Request().Context(), &ownerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.BaseResponse{
		Message: "success",
		Data:    res,
	})
}

func (d *DeadlineHandler) Update(c echo.Context) error {
	var req entity.DeadlineRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	if len(req.DeadlineDate) == 0 || len(req.Priority) == 0 || len(req.Task) == 0 {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: "empty request data exist",
		})
	}

	deadlineId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	res, err := d.deadlineUsecase.UpdateDeadlineByID(c.Request().Context(), &deadlineId, &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.BaseResponse{
		Message: "success",
		Data:    res,
	})
}

func (d *DeadlineHandler) Delete(c echo.Context) error {
	deadlineId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	err = d.deadlineUsecase.DeleteDeadlineByID(c.Request().Context(), &deadlineId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.BaseResponse{
		Message: "success",
	})
}
