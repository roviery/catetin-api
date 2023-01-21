package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/roviery/catetin-api/domain"
	"github.com/roviery/catetin-api/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionHandler struct {
	transactionUsecase domain.TransactionUsecase
}

func NewTransactionHandler(transactionUsecase domain.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{
		transactionUsecase: transactionUsecase,
	}
}

func (t *TransactionHandler) Create(c echo.Context) error {
	var req entity.TransactionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	if len(req.Title) == 0 || req.Expense == 0 {
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

	res, err := t.transactionUsecase.CreateTransaction(c.Request().Context(), &ownerId, &req)
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

func (t *TransactionHandler) GetTransaction(c echo.Context) error {
	transactionId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	res, err := t.transactionUsecase.GetTransaction(c.Request().Context(), &transactionId)
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

func (t *TransactionHandler) GetTransactions(c echo.Context) error {
	ownerId, err := primitive.ObjectIDFromHex(c.Get("user_id").(string))
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	filter := c.QueryParam("filter")

	res, err := t.transactionUsecase.GetTransactions(c.Request().Context(), &ownerId, &filter)
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

func (t *TransactionHandler) Update(c echo.Context) error {
	var req entity.TransactionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	if len(req.Title) == 0 || req.Expense == 0 {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: "empty request data exist",
		})
	}

	transactionId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	res, err := t.transactionUsecase.UpdateTransactionByID(c.Request().Context(), &transactionId, &req)
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

func (f *TransactionHandler) Delete(c echo.Context) error {
	transactoinId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	err = f.transactionUsecase.DeleteTransactionByID(c.Request().Context(), &transactoinId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.BaseResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.BaseResponse{
		Message: "success",
	})
}
