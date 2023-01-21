package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/roviery/catetin-api/domain"
	"github.com/roviery/catetin-api/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FinanceHandler struct {
	financeUsecase domain.FinanceUsecase
}

func NewFinanceHandler(financeUsecase domain.FinanceUsecase) *FinanceHandler {
	return &FinanceHandler{
		financeUsecase: financeUsecase,
	}
}

func (f *FinanceHandler) Create(c echo.Context) error {
	var req entity.FinanceRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	if len(req.Type) == 0 || req.FundAllocation == 0 {
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

	res, err := f.financeUsecase.CreateFinance(c.Request().Context(), &ownerId, &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, entity.BaseResponse{
		Message: "success",
		Data:    res,
	})
}

func (f *FinanceHandler) GetFinance(c echo.Context) error {
	financeId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	res, err := f.financeUsecase.GetFinance(c.Request().Context(), &financeId)
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

func (f *FinanceHandler) GetFinances(c echo.Context) error {
	ownerId, err := primitive.ObjectIDFromHex(c.Get("user_id").(string))
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	res, err := f.financeUsecase.GetFinances(c.Request().Context(), &ownerId)
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

func (f *FinanceHandler) Update(c echo.Context) error {
	var req entity.FinanceRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	if len(req.Type) == 0 || req.FundAllocation == 0 {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: "empty request data exist",
		})
	}

	financeId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	res, err := f.financeUsecase.UpdateFinanceByID(c.Request().Context(), &financeId, &req)
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

func (f *FinanceHandler) Delete(c echo.Context) error {
	financeId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	err = f.financeUsecase.DeleteFinanceByID(c.Request().Context(), &financeId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.BaseResponse{
		Message: "success",
	})
}
